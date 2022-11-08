package files

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const PathSeparator = string(filepath.Separator)
const Home = "~"
const HomePath = Home + PathSeparator

// AbsolutePath 获取程序目录的绝对路径
func AbsolutePath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Fatalln(err)
	}
	pwd, err := filepath.Abs(filepath.Dir(file))
	if err != nil {
		log.Fatalln(err)
	}
	pwd = strings.Replace(pwd, "\\", "/", -1)
	if strings.HasSuffix(pwd, "/") == false {
		pwd = pwd + "/"
	}
	return pwd
}

// MakeDir 创建目录
func MakeDir(_dir string) bool {
	if IsExist(_dir) == true {
		return true
	}
	err := os.MkdirAll(_dir, os.ModePerm)
	if err != nil {
		return false
	} else {
		return true
	}
}

// IsExist 判断文件或目录是否存在
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

// IsDir 判断是否是目录
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// ReadFile 读取文件
func ReadFile(_file string) string {
	b, err := ioutil.ReadFile(_file)
	if err != nil {
		//Log(err)
		return ""
	}
	str := string(b)
	return str
}

// WriteFile 写文件
func WriteFile(path, data string) bool {
	if ioutil.WriteFile(path, []byte(data), 0644) == nil {
		return true
	} else {
		return false
	}
}

// WriteFileAppend 追加的方式写文件
func WriteFileAppend(path, data string) bool {
	var err error
	fl, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return false
	}
	defer fl.Close()
	_, err = fl.Write([]byte(data))
	if err == nil {
		return true
	}
	return false
}

// ReadDir 读取目录下的文件
func ReadDir(path string) ([]string, error) {
	fs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	files := make([]string, 0)
	for _, item := range fs {
		if item.IsDir() {
			continue
		}
		files = append(files, item.Name())
	}
	return files, nil
}

// Copy 文件拷贝,将src拷贝到dst
func Copy(src, dst string) error {
	f1, err := os.OpenFile(src, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f1.Close()
	//reader := bufio.NewReaderSize(f1, 1024*32)
	f2, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	//writer := bufio.NewWriterSize(f2, 1024*32)
	defer f2.Close()
	_, err = io.Copy(f2, f1)
	return err
}

// Zip 压缩srcFile为压缩文件
func Zip(srcFile string, destZip string) error {
	zipFile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()
	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	return filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = strings.TrimPrefix(path, filepath.Dir(srcFile)+"/")
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})
}

// UnZip 解压文件到指定文件夹
func UnZip(zipFile, dstDir string) (string, error) {
	file, err := zip.OpenReader(zipFile)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var decodeName, zipDir string
	for _, f := range file.File {
		if err = func(f *zip.File) error {
			if f.Flags == 0 {
				//如果标致位是0  则是默认的本地编码   默认为gbk
				i := bytes.NewReader([]byte(f.Name))
				decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
				content, _ := ioutil.ReadAll(decoder)
				decodeName = string(content)
			} else {
				//如果标志为是 1 << 11也就是 2048  则是utf-8编码
				decodeName = f.Name
			}
			if zipDir == "" {
				zipDir = filepath.Base(decodeName)
			}
			decodeName = strings.ReplaceAll(decodeName, `\`, string(filepath.Separator))
			fpath := filepath.Join(dstDir, decodeName)
			if f.FileInfo().IsDir() {
				if err = os.MkdirAll(fpath, os.ModePerm); err != nil {
					return err
				}
			} else {
				if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
					return err
				}
				inFile, err := f.Open()
				if err != nil {
					return err
				}
				defer inFile.Close()

				outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
				if err != nil {
					return err
				}
				defer outFile.Close()

				_, err = io.Copy(outFile, inFile)
				if err != nil {
					return err
				}
			}
			return nil
		}(f); err != nil {
			return "", err
		}
	}
	return zipDir, nil
}

// Abs returns the full path of a file or directory, "~" is replaced with home.
func Abs(name string) string {
	if name == "" {
		return ""
	}

	if len(name) > 2 && name[:2] == HomePath {
		if usr, err := user.Current(); err == nil {
			name = filepath.Join(usr.HomeDir, name[2:])
		}
	}

	result, err := filepath.Abs(name)

	if err != nil {
		panic(err)
	}

	return result
}
func Rewrite(oldPath, newPath string) error {
	w, err := os.Create(oldPath)
	defer func() { _ = w.Close() }()
	if err != nil {
		err = fmt.Errorf("Rewrite.oldPath: %w", err)
		return err
	}
	r, err := os.Open(newPath)
	if err != nil {
		err = fmt.Errorf("Rewrite.newPath: %w", err)
		return err
	}
	defer func() { _ = r.Close() }()
	_, err = io.Copy(w, r)
	return err
}

func DirIsNull(dirPath string) bool {
	fList, err := ioutil.ReadDir(dirPath)
	if err != nil || len(fList) == 0 {
		log.Println(err)
		return false
	}
	return true
}

func CopyFile(src, dst string) error {
	r, err := os.Open(src)
	if err != nil {
		err = fmt.Errorf("CopyFile.read: %w", err)
		return err
	}
	defer func() { _ = r.Close() }()

	w, err := os.Create(dst)
	if err != nil {
		err = fmt.Errorf("CopyFile.write: %w", err)
		return err
	}
	defer func() { _ = w.Close() }()
	_, err = io.Copy(w, r)
	return nil
}
