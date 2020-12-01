package plugins

import (
	"github.com/gocolly/colly/proxy"
	"m_troops/go/spider/Mcooly/common/logs"
	"m_troops/go/spider/Mcooly/work/model"
)

func MiddlewareFiddlerProxy(s *model.SpiderParams) {
	// Rotate two socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher("http://127.0.0.1:8888")
	logs.Errors(err)
	s.Collector.SetProxyFunc(rp)
}
