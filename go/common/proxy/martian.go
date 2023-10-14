// 使用官方包 github.com/google/martian 实现拦截
package proxy

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/google/martian"
	"github.com/google/martian/auth"
	"github.com/google/martian/log"
	"github.com/google/martian/priority"
	"github.com/google/martian/proxyauth"
	"net"
	"net/http"
	"net/url"
)

var (
	Proxy       *martian.Proxy
	ProxyServer string
	UserName    string
	PassWord    string
)

func Martian() error {
	Proxy = martian.NewProxy()
	group := priority.NewGroup()

	if UserName != "" {
		a := proxyauth.NewModifier()
		group.AddRequestModifier(a, 2)
		group.AddResponseModifier(a, 2)
	}

	s := &Skip{}
	group.AddRequestModifier(s, 1)
	group.AddResponseModifier(s, 1)

	Proxy.SetRequestModifier(group)
	Proxy.SetResponseModifier(group)
	// 使用代理发请求时候装载证书
	if ProxyServer != "" {
		fmt.Println("开启代理：", ProxyServer)
		CertReload()
		mc, err := GetMITMConfig()
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

func (r *Skip) ModifyRequest(req *http.Request) error {
	if ProxyServer != "" {
		u, err := url.Parse(ProxyServer)
		if err != nil {
			return err
		}

		Proxy.SetDownstreamProxy(u)
	}

	ctx := martian.NewContext(req)
	authCTX := auth.FromContext(ctx)
	if authCTX.ID() != fmt.Sprintf("%s:%s", UserName, PassWord) {
		authCTX.SetError(errors.New("auth error"))
		ctx.SkipRoundTrip()
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
