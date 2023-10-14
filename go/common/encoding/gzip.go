package encoding

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
)

func Compression(oriData []byte) []byte {
	var buf bytes.Buffer
	write := gzip.NewWriter(&buf)
	_, err := write.Write(oriData)
	if err != nil {
		log.Fatalln(err)
	}
	_ = write.Flush()
	return buf.Bytes()
}

func Decompress(oriData []byte) []byte {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln(r)
		}
	}()

	var buf bytes.Buffer
	buf.Write(oriData)
	read, err := gzip.NewReader(&buf)
	if err != nil {
		log.Fatalln(err)
	}
	pData, _ := ioutil.ReadAll(read)
	return pData
}
