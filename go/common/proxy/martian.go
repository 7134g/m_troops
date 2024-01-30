package proxy

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/google/martian"
	"github.com/google/martian/auth"
	"github.com/google/martian/log"
	"github.com/google/martian/mitm"
	"github.com/google/martian/priority"
	"github.com/google/martian/proxyauth"
	"net"
	"net/http"
	"net/url"
)

var (
	MonitorAddress = "127.0.0.1:10888" // 监听地址
)

var (
	httpMartian *martian.Proxy // 拦截器全局对象
	certFlag    bool           // 开启自签证书验证
)

var (
	serverProxyUrlParse *url.URL // 解析代理

	serverProxyFlag     bool   // 启用代理
	serverProxy         string // 服务代理地址
	serverProxyUsername string // 用户名
	serverProxyPassword string // 密码
)

func init() {
	log.SetLevel(log.Silent)
}

func OpenCert() {
	certFlag = true
	_ = LoadCert()
}

func SetServeProxyAddress(address, username, password string) {
	serverProxyFlag = true
	serverProxy = address
	serverProxyUsername = username
	serverProxyPassword = password
}

func Martian() error {
	httpMartian = martian.NewProxy()
	if certFlag {
		mc, err := mitm.NewConfig(ca, private)
		if err != nil {
			return err
		}
		httpMartian.SetMITM(mc)
	}

	if serverProxyFlag {
		u, err := url.Parse(serverProxy)
		if err != nil {
			return err
		}
		serverProxyUrlParse = u
	}

	group := priority.NewGroup()
	xs := newSkip()
	group.AddRequestModifier(xs, 10)
	group.AddResponseModifier(xs, 10)
	xa := newAuth(proxyauth.NewModifier())
	group.AddRequestModifier(xa, 12)
	group.AddResponseModifier(xa, 12)
	httpMartian.SetRequestModifier(group)
	httpMartian.SetResponseModifier(group)

	listener, err := net.Listen("tcp", MonitorAddress)
	if err != nil {
		return err
	}

	err = httpMartian.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}

type skip struct {
}

func newSkip() *skip {
	return &skip{}
}

func (r *skip) ModifyRequest(req *http.Request) error {
	// todo 编写想要处理的请求
	fmt.Println(req.Method, req.URL.String())

	return nil
}

func (r *skip) ModifyResponse(res *http.Response) error {
	// todo 编写想要处理的请求
	return nil
}

type xauth struct {
	pAuth *proxyauth.Modifier
}

func newAuth(pAuth *proxyauth.Modifier) *xauth {
	return &xauth{pAuth: pAuth}
}

func (r *xauth) ModifyRequest(req *http.Request) error {
	if serverProxy == "" {
		return nil
	}

	httpMartian.SetDownstreamProxy(serverProxyUrlParse)

	if serverProxyUsername != "" {
		un := base64.StdEncoding.EncodeToString([]byte(serverProxyUsername))
		pw := base64.StdEncoding.EncodeToString([]byte(serverProxyPassword))
		//req.Header.Set("Proxy-Authorization", fmt.Sprintf("Basic %s:%s", un, pw))
		ctx := martian.NewContext(req)
		authCTX := auth.FromContext(ctx)
		if authCTX.ID() != fmt.Sprintf("%s:%s", un, pw) {
			authCTX.SetError(errors.New("auth error"))
			ctx.SkipRoundTrip()
		}
	}

	return nil
}

func (r *xauth) ModifyResponse(res *http.Response) error {
	if serverProxy == "" {
		return nil
	}
	return r.pAuth.ModifyResponse(res)
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
