### 生成api服务
- mkdir user
- cd user
- vim user.api
- goctl api go -api user.api -dir ./api --style=go_zero



### 生成rpc调用
- mkdir user
- cd user
- vim user.proto
- goctl rpc protoc user.proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc --style go_zero


### 启动
- go run user.go
