package handler

import (
	"demo/app/meeting/api/internal/svc"
	"demo/common/ws"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func RegisterHandlersWS(server *rest.Server, serverCtx *svc.ServiceContext) {
	hub := ws.NewHub()
	go hub.Run()

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/ws",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			ws.ServeWs(hub, w, r)
		},
	})
}
