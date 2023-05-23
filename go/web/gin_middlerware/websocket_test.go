package middlerware

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestWebsocket(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", ping)
	t.Log(r.Run("localhost:2303"))
	// 双击打开 ws.html 测试
}
