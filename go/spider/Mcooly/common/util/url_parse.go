package util

import (
	"m_troops/go/spider/Mcooly/common/logs"
	"net/url"
)

// 解析
func UrlParse(u string) *url.URL {
	urlParse, err := url.Parse(u)
	logs.Errors(err)
	return urlParse
}

// 编码
func UrlEncode(s string) string {
	urlParse := UrlParse(s)
	params := urlParse.Query()
	urlParse.RawQuery = params.Encode()
	return urlParse.String()
}

// 修改 query
func UrlChangeQuery(s string, q map[string]string) string {
	urlParse := UrlParse(s)
	params := urlParse.Query()
	for k, v := range q {
		params.Set(k, v)
	}
	urlParse.RawQuery = params.Encode()
	return urlParse.String()
}
