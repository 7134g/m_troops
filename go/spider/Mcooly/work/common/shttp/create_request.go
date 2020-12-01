package shttp

import (
	"github.com/gocolly/colly"
	"io"
	"net/http"
	"net/url"
	"spider/setting"
)

// new get
func NewGETRequest(targetUrl *url.URL, headers map[string]string) *colly.Request {
	r := &colly.Request{
		URL:    targetUrl,
		Method: "GET",
		Ctx:    colly.NewContext(),
	}

	// add head
	AddDefultHead(r, headers)

	return r
}

// new post
func NewPOSTRequest(targetUrl *url.URL, requestBody io.Reader, headers map[string]string) *colly.Request {
	r := &colly.Request{
		URL:    targetUrl,
		Method: "POST",
		Ctx:    colly.NewContext(),
		Body:   requestBody,
	}

	// add head
	AddDefultHead(r, headers)

	return r
}

// add head
func AddDefultHead(r *colly.Request, headers map[string]string) {
	h := make(http.Header)
	r.Headers = &h
	for k, v := range setting.HEADERS {
		r.Headers.Set(k, v)
	}
	if headers != nil {
		for k, v := range headers {
			r.Headers.Set(k, v)
		}
	}

}
