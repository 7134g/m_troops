// 使用官方包 github.com/google/martian 实现拦截
package proxy

import (
	"bytes"
	"fmt"
	"github.com/google/martian"
	"github.com/google/martian/log"
	"net"
	"net/http"
	"net/url"
)

var (
	Proxy       *martian.Proxy
	ProxyServer string
)

func Martian() error {
	Proxy = martian.NewProxy()
	Proxy.SetRequestModifier(&Skip{})
	Proxy.SetResponseModifier(&Skip{})
	// 使用代理发请求时候装载证书
	if ProxyServer != "" {
		fmt.Println("开启代理：", ProxyServer)
		mc, err := GlobalCert.GetMITMConfig()
		if err != nil {
			return err
		}
		Proxy.SetMITM(mc)
	}

	log.SetLevel(log.Silent)
	listener, err := net.Listen("tcp", ":1080")
	if err != nil {
		return err
	}

	err = Proxy.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}

type Skip struct {
}

func (r *Skip) ModifyRequest(res *http.Request) error {
	if ProxyServer != "" {
		u, err := url.Parse(ProxyServer)
		if err != nil {
			return err
		}

		Proxy.SetDownstreamProxy(u)
	}
	return nil
}

func (r *Skip) ModifyResponse(res *http.Response) error {
	return nil
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
