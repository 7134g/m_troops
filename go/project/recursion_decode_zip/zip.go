package main

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"context"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"
)

const tempDir = "./temp"

func CleanTemp() {
	_ = filepath.Walk(tempDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		//log.Println(path)
		_ = os.Remove(path)

		return nil
	})
}

type ZipCrypto struct {
	password []byte
	Keys     [3]uint32
}

func newZipCrypto(passphrase []byte) *ZipCrypto {
	z := &ZipCrypto{}
	z.password = passphrase
	z.init()
	return z
}

func (z *ZipCrypto) init() {
	z.Keys[0] = 0x12345678
	z.Keys[1] = 0x23456789
	z.Keys[2] = 0x34567890

	for i := 0; i < len(z.password); i++ {
		z.updateKeys(z.password[i])
	}
}

func (z *ZipCrypto) updateKeys(byteValue byte) {
	z.Keys[0] = crc32update(z.Keys[0], byteValue)
	z.Keys[1] += z.Keys[0] & 0xff
	z.Keys[1] = z.Keys[1]*134775813 + 1
	z.Keys[2] = crc32update(z.Keys[2], (byte)(z.Keys[1]>>24))
}

func (z *ZipCrypto) magicByte() byte {
	var t = z.Keys[2] | 2
	return byte((t * (t ^ 1)) >> 8)
}

func (z *ZipCrypto) Encrypt(data []byte) []byte {
	length := len(data)
	chiper := make([]byte, length)
	for i := 0; i < length; i++ {
		v := data[i]
		chiper[i] = v ^ z.magicByte()
		z.updateKeys(v)
	}
	return chiper
}

func (z *ZipCrypto) Decrypt(ciphertext []byte) []byte {
	length := len(ciphertext)
	plain := make([]byte, length)
	for i, c := range ciphertext {
		v := c ^ z.magicByte()
		z.updateKeys(v)
		plain[i] = v
	}
	return plain
}

func crc32update(pCrc32 uint32, b byte) uint32 {
	return crc32.IEEETable[(pCrc32^uint32(b))&0xff] ^ (pCrc32 >> 8)
}

func ZipCryptoDecrypt(r *io.SectionReader, password []byte) (*io.SectionReader, error) {
	z := newZipCrypto(password)
	b := make([]byte, r.Size())
	_, _ = r.Read(b)

	m := z.Decrypt(b)
	return io.NewSectionReader(bytes.NewReader(m), 12, int64(len(m))), nil
}

//type Unzip struct {
//	offset       int64
//	fp           *os.File
//	name         string
//	targetDstDir string
//
//	zr *zip.Reader
//}
//
//func (uz *Unzip) init() (err error) {
//	uz.fp, err = os.Open(uz.name)
//	return err
//}
//
//func (uz *Unzip) close() {
//	if uz.fp != nil {
//		_ = uz.fp.Close()
//	}
//}
//
//func (uz *Unzip) Size() int64 {
//	if uz.fp == nil {
//		if err := uz.init(); err != nil {
//			return -1
//		}
//	}
//
//	fi, err := uz.fp.Stat()
//	if err != nil {
//		return -1
//	}
//
//	return fi.Size() - uz.offset
//}
//
//func (uz *Unzip) ReadAt(p []byte, off int64) (int, error) {
//	if uz.fp == nil {
//		if err := uz.init(); err != nil {
//			return 0, err
//		}
//	}
//
//	return uz.fp.ReadAt(p, off+uz.offset)
//}
//
//// DeCompressZip 解压zip包
//func DeCompressZip(zipFile, dest, passwd string, _ []string, offset int64) error {
//	uz := &Unzip{offset: offset, name: zipFile}
//	defer uz.close()
//
//	zr, err := zip.NewReader(uz, uz.Size())
//	if err != nil {
//		return err
//	}
//
//	if passwd != "" {
//		// Register a custom Deflate compressor.
//		zr.RegisterDecompressor(zip.Deflate, func(r io.Reader) io.ReadCloser {
//			rs := r.(*io.SectionReader)
//			r, _ = ZipCryptoDecrypt(rs, []byte(passwd))
//			return flate.NewReader(r)
//		})
//
//		zr.RegisterDecompressor(zip.Store, func(r io.Reader) io.ReadCloser {
//			rs := r.(*io.SectionReader)
//			r, _ = ZipCryptoDecrypt(rs, []byte(passwd))
//			return io.NopCloser(r)
//		})
//	}
//
//	for _, f := range zr.File {
//		fpath := filepath.Join(dest, f.Name)
//		if f.FileInfo().IsDir() {
//			_ = os.MkdirAll(fpath, os.ModePerm)
//			continue
//		}
//
//		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
//			return err
//		}
//
//		inFile, err := f.Open()
//		if err != nil {
//			return err
//		}
//
//		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
//		if err != nil {
//			_ = inFile.Close()
//			return err
//		}
//
//		_, err = io.Copy(outFile, inFile)
//		_ = inFile.Close()
//		_ = outFile.Close()
//		switch {
//		case err == nil:
//			return nil
//		case err.Error() == "zip: checksum error":
//			return decodeError
//		default:
//			return err
//		}
//
//	}
//
//	return nil
//}

var decodeError = errors.New("decode error")

type DeCompressZip struct {
	lock *sync.Mutex

	concurrence uint   // 并发数
	src         string // 压缩包路径
	dst         string // 解压路径

	passChan chan string // 密码通道
	passwd   string      // 密码
	srcFile  *os.File    // 压缩包文件对象

	zr *zip.Reader // zip对象
}

func (uz *DeCompressZip) init() (err error) {
	uz.srcFile, err = os.Open(uz.src)
	return err
}

func (uz *DeCompressZip) close() {
	if uz.srcFile != nil {
		_ = uz.srcFile.Close()
	}
}

func (uz *DeCompressZip) Size() int64 {
	if uz.srcFile == nil {
		if err := uz.init(); err != nil {
			return -1
		}
	}

	fi, err := uz.srcFile.Stat()
	if err != nil {
		return -1
	}

	return fi.Size()
}

func (uz *DeCompressZip) ReadAt(p []byte, off int64) (int, error) {
	if uz.srcFile == nil {
		if err := uz.init(); err != nil {
			return 0, err
		}
	}

	return uz.srcFile.ReadAt(p, off)
}

func NewDeCompressZip(zipFile, dest string) (*DeCompressZip, error) {
	err := os.MkdirAll(tempDir, os.ModeDir)
	if err != nil {
		return nil, err
	}

	uz := &DeCompressZip{lock: &sync.Mutex{}, src: zipFile, dst: dest, concurrence: 1}
	if err := uz.init(); err != nil {
		return nil, err
	}

	return uz, nil
}

func (uz *DeCompressZip) SetPasswdTask(passChan chan string, concurrence uint) {
	uz.passChan = passChan
	uz.concurrence = concurrence
}

func (uz *DeCompressZip) tryDecrypt(number int, zf *zip.File) (bool, error) {
	inFile, err := zf.Open()
	if err != nil {
		return false, err
	}
	defer inFile.Close()

	fPath := filepath.Join(tempDir, fmt.Sprintf("%d", number))
	outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zf.Mode())
	if err != nil {
		return false, err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, inFile)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (uz *DeCompressZip) verifyPasswd() bool {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	var success bool
	execute := func(number int) {
		workerName := fmt.Sprintf("line_%d", number)
		wg.Add(1)
		defer wg.Done()

		var targetFile *zip.File
		zr, err := zip.NewReader(uz, uz.Size())
		if err != nil {
			log.Fatal(err)
		}
		if len(zr.File) == 0 {
			log.Fatal("zip file length is 0")
		}
		for _, file := range zr.File {
			if file.FileInfo().IsDir() {
				continue
			}

			targetFile = file
			break
		}
		for {
			select {
			case passwd := <-uz.passChan:
				if passwd == "" {
					log.Println(workerName, "passwd done, but crypt fail.")
					return
				}

				atomic.AddInt64(&cryptCount, 1)

				zr.RegisterDecompressor(zip.Deflate, func(r io.Reader) io.ReadCloser {
					rs := r.(*io.SectionReader)
					r, _ = ZipCryptoDecrypt(rs, []byte(passwd))
					return flate.NewReader(r)
				})
				ok, err := uz.tryDecrypt(number, targetFile)
				if err != nil {
					//log.Printf("%s passwd: %s err: %v\n", workerName, passwd, err)
					continue
				}
				if !ok {
					continue
				}

				// 锁
				uz.lock.Lock()
				log.Println(workerName, "crypt success: ", passwd)
				cancel() // 关闭其他解析 goroutin
				success = ok
				if uz.passwd == "" {
					uz.passwd = passwd
				}
				uz.lock.Unlock()
				return
			case <-ctx.Done():
				return
			}
		}
	}

	for i := 0; i < int(uz.concurrence); i++ {
		go execute(i)
	}

	time.Sleep(time.Second)
	wg.Wait()
	return success
}

func (uz *DeCompressZip) run() error {
	defer CleanTemp()

	success := true
	if uz.passChan != nil {
		success = uz.verifyPasswd()
	}

	if !success {
		return errors.New("crypt zip is error")
	}

	zr, err := zip.NewReader(uz, uz.Size())
	if err != nil {
		log.Fatal(err)
	}
	uz.zr = zr
	uz.zr.RegisterDecompressor(zip.Deflate, func(r io.Reader) io.ReadCloser {
		rs := r.(*io.SectionReader)
		r, _ = ZipCryptoDecrypt(rs, []byte(uz.passwd))
		return flate.NewReader(r)
	})

	for _, f := range uz.zr.File {
		fpath := filepath.Join(uz.dst, f.Name)
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		inFile, err := f.Open()
		if err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			_ = inFile.Close()
			return err
		}

		_, err = io.Copy(outFile, inFile)
		_ = inFile.Close()
		_ = outFile.Close()
	}

	return nil
}
