## go-micro
#### 环境
- go get -v github.com/golang/protobuf/protoc-gen-go
- go get -v github.com/micro/micro/v2/cmd/protoc-gen-micro@master
- go get -v github.com/micro/go-micro/v2@latest
- go get -v github.com/micro/micro/v2@latest
- go get -v github.com/micro/go-plugins


#### 异常
- replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
- replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.14.1

#### 工具
##### 生成go文件
- protoc --micro_out=../ --go_out=../ ./hellow.proto

##### 修改字段名
- go get -u github.com/favadi/protoc-go-inject-tag
    - protoc-go-inject-tag -input=../prods.pb.go

##### 启动服务
- consul agent -dev
- micro --registry consul --registry_address 127.0.0.1:8500 call serviceName ServiceStruct.mothod "{\" id\":\"5\"}"

    ```
    set MICRO_REGISTRY=consul
    set MICRO_REGISTRY_ADDRESS=127.0.0.1:8500
    set MICRO_API_NAMESPACE=prodserver
    set MICRO_API_HANDLER=RPC
    micro api
    
    ```


## GRPC
#### proto依赖包
- import "google/protobuf/timestamp.proto"

#### 环境
- google.golang.org/grpc
- github.com/grpc-ecosystem/grpc-gateway/runtime
- github.com/golang/protobuf
- google.golang.org/protobuf
- google.golang.org/genproto

#### 生成pb文件
- protoc -I . --go_out=plugins=grpc:. ./文件名.proto

#### 生成http的网关
- protoc --grpc-gateway_out=logtostderr=true:../文件名 ./文件名.proto

##### 修改字段名
- go get -u github.com/favadi/protoc-go-inject-tag
    - protoc-go-inject-tag -input=../prods.pb.go

#### 工具
- go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
- go get -v -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
- go get -v -u github.com/golang/protobuf/protoc-gen-go


#### 去除omitempty标签
- 使用protoc-go-inject-tag将pb文件json标签全局替换
- 若grpc-gateway
```
gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
```