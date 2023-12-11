#### etcd服务
```shell
docker pull bitnami/etcd:latest
docker run -d --name etcd \
    --network bridge \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://192.168.1.26:2379 \
    bitnami/etcd:latest
```
- etcd: 虚拟机中的docker `192.168.1.8:2379`
  - 查看所有键：`etcdctl --endpoints http://192.168.1.26:2379 get "" --prefix=true`


### 生成代码语句
- `goctl api go -api *.api -dir ./api --style=go_zero`
- `goctl rpc protoc article.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc --style go_zero`
- `goctl model mysql datasource -url="root:mysql@tcp(127.0.0.1:3306)/blog" -table="article" -dir="./model/mysql/article_gorm" --style=go_zero -cache=true --home=./deploy/template`
- `goctl model mysql ddl -src="./deploy/sql/op.sql" --dir="./model/gorm_zero" --style=go_zero --home=./deploy/template --strict`



#### swagger
- `go install github.com/zeromicro/goctl-swagger@latest`
- `goctl api plugin -plugin goctl-swagger="swagger -filename article.json" -api article.api -dir .`
- `swagger serve -F=swagger ./article.json`


#### 依赖
- 包
```shell
go get github.com/zeromicro/go-zero@latest
go get google.golang.org/grpc
go get github.com/Masterminds/squirrel
go get github.com/jinzhu/copier
go get github.com/go-playground/validator/v10

go get github.com/pion/webrtc/v3
```
- 工具
```shell
go install github.com/zeromicro/goctl-swagger@latest
go install github.com/zeromicro/go-zero/tools/goctl@latest
```
- 使用 goctl 安装
```shell
goctl env check --install --verbose --force
```
