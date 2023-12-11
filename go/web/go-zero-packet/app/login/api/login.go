package main

import (
	"demo/app/login/api/internal/config"
	"demo/app/login/api/internal/handler"
	"demo/app/login/api/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/router"
	"net/http"
)

var configFile = flag.String("f", "etc/login-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithRouter(&myRouter{router.NewRouter()}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	logx.DisableStat()
	server.PrintRoutes()
	server.Start()
}

type myRouter struct {
	httpx.Router
}

func (r *myRouter) Handle(method, path string, handler http.Handler) error {
	if method == http.MethodGet && path == "/" {
		return r.Router.Handle(method, path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("pre-handle")
			handler.ServeHTTP(w, r)
		}))
	}

	return r.Router.Handle(method, path, handler)
}
