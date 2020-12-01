package common

import (
	"github.com/gocolly/colly"
	"m_troops/go/spider/Mcooly/common/logs"
)

func SpiderRequestError(r *colly.Request, err error) {
	if err != nil {
		logs.Errors(err, r.URL.String())
	}
}

func SpiderResponseError(r *colly.Response, err error) {
	if err != nil {
		logs.Errors(err, r.Request.URL.String())
	}
}
