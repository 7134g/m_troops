### 生成api服务
- goctl api go -api *.api -dir ../ --style=go_zero



### 生成rpc调用
- goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=go_zero