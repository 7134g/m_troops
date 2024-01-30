package proxy

import (
	"bytes"
	"net/http"
)

// 用于读出 response 再重新写入
type writer struct {
	write        http.ResponseWriter
	code         int
	responseBody *bytes.Buffer
}

func newWrite(write http.ResponseWriter) *writer {
	return &writer{
		write:        write,
		responseBody: bytes.NewBuffer(nil),
	}
}

func (w *writer) Header() http.Header {
	return w.write.Header()
}

func (w *writer) Write(bytes []byte) (int, error) {
	w.responseBody.Write(bytes)
	return w.write.Write(bytes)
}

func (w *writer) WriteHeader(statusCode int) {
	w.code = statusCode
	w.write.WriteHeader(statusCode)
}
