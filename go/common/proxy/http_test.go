package proxy

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http/httptest"
	"testing"
)

func TestNewWrite(t *testing.T) {
	response := &httptest.ResponseRecorder{}
	lw := newWrite(response)
	_, _ = lw.Write([]byte("abc"))
	t.Log(lw.responseBody.String())
}

func TestParseHtml(t *testing.T) {
	data := bytes.NewBuffer([]byte(`
<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Go 语言环境安装 | 菜鸟教程</title>
		<title>不要的</title>
		<title>多余的</title>
	</head>
</html>
`))
	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		log.Fatal(err)
	}
	node := doc.Find("title").First().Text()
	t.Log(node)
}
