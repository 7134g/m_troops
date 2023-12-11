package bridging

import (
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func RegisterHandlers(server *rest.Server) {
	route := oldRoute()

	zeroRouteByGin := make([]rest.Route, 0)
	for _, g := range route.Routes() {
		r := rest.Route{
			Method:  g.Method,
			Path:    g.Path,
			Handler: route.Handler().ServeHTTP,
		}

		zeroRouteByGin = append(zeroRouteByGin, r)
	}
	server.AddRoutes(zeroRouteByGin)
}

func oldRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
		return
	})

	return r
}
