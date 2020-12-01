package middlerware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//设置可跨域的url
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//设置超时时间
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		//可访问的方法
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST...") // * 允许所有
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}

	}
}
