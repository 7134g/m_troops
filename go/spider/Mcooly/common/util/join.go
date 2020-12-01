package util

import (
	"bytes"
	"io"
)

func AppendBytes(o io.Reader, key *[]byte) ([]byte, string) {
	if key == nil {
		key = &[]byte{}
	}

	var buf bytes.Buffer
	body := make([]byte, 1024)
	n, _ := o.Read(body)
	c := body[:n]
	buf.Write(c)

	for n == 1024 {
		n, _ = o.Read(body)
		buf.Write(body[:n])
	}
	*key = append(*key, buf.Bytes()...)
	return *key, buf.String()
}
