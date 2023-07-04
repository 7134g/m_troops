## 参考

1. gin

<https://www.cnblogs.com/quchunhui/p/16673000.html>
<https://github.com/swaggo/gin-swagger>
<https://github.com/swaggo/swag/blob/master/README_zh-CN.md>

2. go-zero

<https://blog.csdn.net/qq_26372385/article/details/88886047>
<https://github.com/zeromicro/goctl-swagger>

## 安装

1. gin
- `go install github.com/swaggo/swag/cmd/swag@latest`
- `go get -u github.com/swaggo/gin-swagger`
- `go get -u github.com/swaggo/files`

2. go-zero

- `go install github.com/zeromicro/goctl-swagger@latest`



## 生成文档
1. gin
   - `swag init -g ./main.go -o ./docs`

1. go-zero
   1. go-swagger 文档
      - `goctl api plugin -plugin goctl-swagger="swagger -filename user.json -host 127.0.0.2 -basepath /api" -api user.api -dir .`

   1. docker 方式查看文档
      - ` docker run --rm -p 8083:8080 -e SWAGGER_JSON=/foo/user.json -v $PWD:/foo swaggerapi/swagger-ui`

   1. web方式查看文档
      - `git clone git@github.com:go-swagger/go-swagger.git`
      - `cd go-swagger`
      - `go install ./cmd/swagger`
      - `swagger serve -F=swagger ./user.json`

## 描述

```

@Tags 接口分组，相当于归类
@Summary 接口的简要描述
@Description 接口的详细描述
@Id 全局标识符
@Version 接口版本号
@Accept 发起请求的数据类型
@Produce 接口返回的数据类型
@Param 请求参数描述
@Success 请求成功后返回信息
@Failure 请求失败后返回信息
@Router 请求的路由路径和请求方式。
@contact.name 接口联系人
@contact.url 联系人网址
@contact.email 联系人邮箱

```

### GET

    // @Summary 查看迁移任务详细信息
    // @Description 查看迁移任务详细信息
    // @Accept json
    // @Produce  json
    // @Param task_id query string true "task_id"
    // @Param timestamp query int false "task_id"
    // @Success 200 {object} models.Response "请求成功"
    // @Failure 400 {object} models.ResponseErr "请求错误"
    // @Failure 500 {object} models.ResponseErr "内部错误"
    // @Router /task [get]

### POST

    // @Summary 创建镜像迁移任务
    // @Description 创建镜像迁移任务
    // @Accept  json
    // @Produce  json
    // @Param data body CreateTaskReq true "请示参数data"
    // @Success 200 {object} models.Response "请求成功"
    // @Failure 400 {object} models.ResponseErr "请求错误"
    // @Failure 500 {object} models.ResponseErr "内部错误"
    // @Router /task [post]

### DELETE

    // @Summary 删除镜像迁移任务
    // @Description 删除镜像迁移任务
    // @Accept  json
    // @Produce  json
    // @Param data body TaskReq true "请示参数data"
    // @Success 200 {object} models.Response "请求成功"
    // @Failure 400 {object} models.ResponseErr "请求错误"
    // @Failure 500 {object} models.ResponseErr "内部错误"
    // @Router /task [delete]

### PUT

    // @Summary 更新镜像迁移任务
    // @Description 更新镜像迁移任务
    // @Accept  json
    // @Produce  json
    // @Param data body CreateTaskReq true "请示参数data"
    // @Success 200 {object} models.Response "请求成功"
    // @Failure 400 {object} models.ResponseErr "请求错误"
    // @Failure 500 {object} models.ResponseErr "内部错误"
    // @Router /task [put]



## 例子
```golang
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
