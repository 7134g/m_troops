goctl api go -api login.api -dir ../api  --style=go_zero --home=../../../deploy/template

#goctl rpc protoc login.proto --go_out=../rpc/pb --go-grpc_out=../rpc/pb --zrpc_out=../rpc --style=go_zero --home=../../../deploy/template
#protoc-go-inject-tag -input="../rpc/pb/login/login.pb.go"

#mkdir ../gateway
#mkdir ../gateway/pb
#protoc --descriptor_set_out=../gateway/pb/login.pb login.proto
