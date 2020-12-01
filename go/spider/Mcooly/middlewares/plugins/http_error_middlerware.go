package plugins

import (
	"github.com/gocolly/colly"
	"m_troops/go/spider/Mcooly/common/logs"
	"m_troops/go/spider/Mcooly/work/common/ctx"
	"m_troops/go/spider/Mcooly/work/model"
	"strings"
)

// 响应错误及重试
func MidderwareHttpError(s *model.SpiderParams) {
	s.Collector.OnResponse(func(r *colly.Response) {
		ctx.PutInt(r.Ctx, "httpCode", r.StatusCode)
		// spider 计数 成功获取到新的响应，错误置零
		s.SpiderErrorZero(r)
	})

	s.Collector.OnError(func(r *colly.Response, err error) {
		// http 计数
		ctx.PutInt(r.Ctx, "httpCode", r.StatusCode)

		proxyFlag := false
		errString := ""
		if err != nil {
			errString = err.Error()
			proxyFlag = strings.Contains(errString, "proxyconnect")
			if proxyFlag {
				errString = "proxy connect error"
				s.ChangeProxy()
			}
			if strings.Contains(errString, "Client.Timeout") {
				errString = "Client Timeout error"
			}
			if strings.Contains(errString, "wsarecv:") {
				errString = "An existing connection was forcibly closed by the remote host"
			}

		}

		if r.StatusCode == 403 || proxyFlag {
			ctx.PutBool(r.Ctx, "proxy_error", true)
			ctx.PutInt(r.Ctx, "retryCount", 0)
			s.AddProxy()
			logs.Log.Errorf("%s ip不可用，更换ip, StatusCode：%d  新ip：%s ERROR：%s", s.SpiderName, r.StatusCode, s.ProxyValue, errString)
			//s.Stop()
			//return
		}
		s.HttpRetry(r.Request, errString)
	})
}
