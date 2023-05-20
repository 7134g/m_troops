## 注解说明
    // @Summary 接口概要说明
    // @Description 接口详细描述信息
    // @Tags 用户信息   //swagger API分类标签, 同一个tag为一组
    // @accept json  //浏览器可处理数据类型，浏览器默认发 Accept: */*
    // @Produce  json  //设置返回数据的类型和编码
    // @Param id path int true "ID"    //url参数：（name；参数类型[query(?id=),path(/123)]；数据类型；required；参数描述）
    // @Param name query string false "name"
    // @Success 200 {object} Res {"code":200,"data":null,"msg":""}      //成功返回的数据结构， 最后是示例
    // @Failure 400 {object} Res {"code":200,"data":null,"msg":""}
    // @Router /test/{id} [get]    //路由信息，一定要写上


## 例子
```
package main

import (
    "apiwendang/controller"
    _ "apiwendang/docs"
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)


// @title Docker监控服务
// @version 1.0
// @description docker监控服务后端API接口文档

// @contact.name API Support
// @contact.url http://www.swagger.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:9009
// @BasePath
func main() {
    r := gin.New()

    r.Use(Cors())
    //url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
    r.POST("/test/:id", controller.Test)
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.Run(":9009")
}


func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Next()
    }
}


// @Summary 测试接口
// @Description 描述信息
// @Success 200 {string} string    "ok"
// @Router / [get]
func Test(ctx *gin.Context)  {
    ctx.JSON(200, "ok")
}
```


## source
- https://www.cnblogs.com/zhzhlong/p/11800787.html
