package proxy

import (
	"fmt"
	"io/ioutil"
	"m_troops/go/spider/Mcooly/common/logs"
	"m_troops/go/spider/Mcooly/setting"
	"net/http"
	"net/url"
)

func ProxyFiddle() func(*http.Request) (*url.URL, error) {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:8888")
	}
	return proxy
}

// 获取新值
func ProxyPoolGet() (func(*http.Request) (*url.URL, error), string) {
	defer proxyPanic()
	resp, err := http.Get(fmt.Sprintf("http://%s/max", setting.PROXY_HOST))
	if resp == nil {
		return nil, ""
	}
	defer resp.Body.Close()
	if err != nil {
		return nil, ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ""
	}

	proxyUrl := "http://" + string(body)
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyUrl)
	}
	return proxy, string(body)
}

// 删除无效
func PoolDelete(proxy string) {
	defer proxyPanic()
	u := fmt.Sprintf("http://%s/useless?proxy=%s", setting.PROXY_HOST, proxy)
	resp, err := http.Get(u)
	if resp == nil {
		return
	}
	defer resp.Body.Close()
	if err != nil {
		return
	}
}

func proxyPanic() {
	if err := recover(); err != nil {
		logs.Errors(fmt.Sprintf("proxy panic: "), err)
	}
}
