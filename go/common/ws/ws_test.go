package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"testing"
)

func TestGinWs(t *testing.T) {
	h := NewHub()
	go h.Run()

	route := gin.Default()
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})
	route.GET("/ws", func(ctx *gin.Context) {
		//conn, err := upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
		//if err != nil {
		//	ctx.JSON(http.StatusOK, err)
		//	return
		//}
		//
		//client := &Client{
		//	hub:  h,
		//	conn: conn,
		//	send: make(chan []byte, bufSize),
		//	readHandle: func(message []byte) []byte {
		//		return message
		//	},
		//	writeHandle: func(message []byte) []byte {
		//		return append([]byte("send write "), message...)
		//	},
		//}
		//
		//client.hub.register <- client
		//go client.writePump()
		//go client.readPump()
		if err := DefaultServeWs(h, ctx.Writer, ctx.Request); err != nil {
			ctx.JSON(http.StatusOK, err)
			return
		}

	})

	if err := route.Run(":10999"); err != nil {
		panic(err)
	}

}

func TestZeroWs(t *testing.T) {
	h := NewHub()
	go h.Run()

	engine := rest.MustNewServer(rest.RestConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Host:         "localhost",
		Port:         10999,
		Timeout:      10000,
		CpuThreshold: 500,
	})
	defer engine.Stop()

	engine.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("ok"))
		},
	})

	engine.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/ws",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			client, err := InitClient(h, w, r)
			if err != nil {
				logx.Error(err)
				return
			}
			client.SetReadHandle(func(message []byte) []byte {
				return append([]byte("receive a message "), message...)
			})

			client.SetWriteHandle(func(message []byte) []byte {
				return append(message, []byte(", write back")...)
			})
			Run(client)
		},
	})

	engine.PrintRoutes()
	engine.Start()
}
