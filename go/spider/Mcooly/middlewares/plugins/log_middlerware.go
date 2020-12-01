package plugins

import (
	"fmt"
	"github.com/gocolly/colly"
	"m_troops/go/spider/Mcooly/common/logs"
	"m_troops/go/spider/Mcooly/work/model"
)

func MidderwareLogger(s *model.SpiderParams) {
	s.Collector.OnRequest(func(r *colly.Request) {
		var msg string
		if s.ProxyValue != "" {
			msg = fmt.Sprintf("%s[%s] %s %s", s.SpiderName, r.Method, s.ProxyValue, r.URL.String())
		} else {
			msg = fmt.Sprintf("%s[%s] localhost %s", s.SpiderName, r.Method, r.URL.String())
		}
		logs.Log.Debug(msg)
	})
}

func TestMidderware(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		logs.Log.Debug("sssssss", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {
		logs.Log.Debug("pppppppppp", r.Request.URL.String())
	})
}
