package proxy

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io"
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

func ParseHtmlTitle(r io.Reader) (string, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", err
	}
	node := doc.Find("title")
	if node == nil {
		return "", errors.New("cannot find title")
	}
	fn := node.First()
	if fn == nil {
		return "", errors.New("cannot find title")
	}
	title := fn.Text()
	return title, nil
}

// ExtractRequestToString 提取请求包
func ExtractRequestToString(res *http.Request) string {
	buf := bytes.NewBuffer([]byte{})
	defer buf.Reset()
	err := res.Write(buf)
	if err != nil {
		return ""
	}

	return buf.String()
}
