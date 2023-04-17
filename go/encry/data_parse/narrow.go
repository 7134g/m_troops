package data_parse

import (
	"bytes"
	"compress/gzip"
	"encry/common/excetion"
	"io/ioutil"
	"log"
)

func Compression(oridata []byte) (pdata []byte) {
	var buf bytes.Buffer
	write := gzip.NewWriter(&buf)
	_, err := write.Write(oridata)
	if err != nil {
		log.Fatalln("[x] Compression data failure, system msg: ", err)
	}
	write.Flush()
	pdata = buf.Bytes()
	return
}

func Decompress(oridata []byte) (pdata []byte) {
	defer excetion.ErrorRecover("[X] Gzip decompress error")

	var buf bytes.Buffer
	buf.Write(oridata)
	read, err := gzip.NewReader(&buf)
	if err != nil {
		log.Fatalln("[x] Failed to decompress data, system msg: ", err)
	}
	pdata, _ = ioutil.ReadAll(read)
	return
}
