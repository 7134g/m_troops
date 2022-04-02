// 使用官方包 github.com/google/martian 实现拦截
package proxy

import (
	"github.com/google/martian"
	"net"
	"net/http"
)

var (
	Proxy *martian.Proxy
)

func Martian() error {
	Proxy = martian.NewProxy()
	Proxy.SetRequestModifier(&Skip{})
	Proxy.SetResponseModifier(&Skip{})

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
	return nil
}

func (r *Skip) ModifyResponse(res *http.Response) error {
	return nil
}
