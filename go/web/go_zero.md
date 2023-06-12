### 生成api服务
- goctl api go -api *.api -dir ../user --style=go_zero



### 生成rpc调用
- goctl rpc protoc *.proto --go_out=../user --go-grpc_out=../user --zrpc_out=../user --style=go_zero


### 启动
- go run user.go
