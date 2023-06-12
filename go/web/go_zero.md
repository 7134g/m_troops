### 生成api服务
- goctl api go -api *.api -dir ../ --style=go_zero



### 生成rpc调用
- goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=go_zero





## 命令

+ 创建API服务

```shell
goctl api new 服务名称
# 1. 创建 user 服务
goctl api new user
# 2. 创建 admin 服务
goctl api new admin
# 3. 创建 open 服务
goctl api new open
```

+ 生成服务代码

```shell
goctl api go -api 服务名称.api -dir . -style go_zero
# 1. 生成 user api 服务代码
goctl api go -api user.api -dir . -style go_zero
# 2. 生成 admin api 服务代码
goctl api go -api admin.api -dir . -style go_zero
# 3. 生成 user rpc 服务代码
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. --style go_zero
# 4. 生成 device rpc 服务代码
goctl rpc protoc device.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. --style go_zero
# 5. 生成 open api 服务代码
goctl api go -api open.api -dir . -style go_zero
```

+ 启动服务

```shell
go run 服务名称.go -f 配置文件地址
# 1. 启动 user api 服务
go run user.go -f etc/user-api.yaml
# 2. 启动 admin api 服务
go run admin.go -f etc/admin-api.yaml
# 3. 启动 user rpc 服务
go run user.go -f etc/user.yaml
# 4. 启动 device rpc 服务
go run device.go -f etc/device.yaml
# 5. 启动 open api 服务
go run open.go -f etc/open-api.yaml
```
