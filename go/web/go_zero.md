### 依赖
- go install github.com/zeromicro/go-zero/tools/goctl@latest
- go get -u github.com/zeromicro/go-zero@latest
- go install github.com/favadi/protoc-go-inject-tag@latest
- go get github.com/go-playground/validator/v10


### 启动
- `go run user.go`
- `grpcui -plaintext localhost:12345` grpc调试

### 生成api服务
- mkdir user
- cd user
- vim user.api
- goctl api go -api *.api -dir ./api --style=go_zero



### 生成rpc调用
- mkdir user
- cd user
- vim user.proto
- goctl rpc protoc user.proto --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc --style go_zero
- protoc-go-inject-tag -input=./rpc/user/user.pb.go


### 生成 dockerfile 和 k8s 服务发现部署

#### 生成 dockerfile
1. 
   - `cd user/api`
   - `goctl docker -go user.go`
   - `docker build user-api:v1 .`
   - `docker push user-api:v1`


2. 
   - `cd user/rpc`
   - `goctl docker -go user.go`
   - `docker build user-rpc:v1 .`
   - `docker push user-rpc:v1`


   
#### k8s

api的yaml文件配置k8s
```yaml
Name: user-api
Host: 0.0.0.0
Port: 14000

UserRpcConfig:
  Target: k8s://user-model/user-rpc:14001
```
rpc的yaml文件配置k8s
```yaml
Name: user.rpc
ListenOn: 127.0.0.1:14001
```




1. 为角色增加调用其他pod的k8s权限
`kubectl apply -f auth.yaml`
````yaml
#创建账号
apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: user-model
  name: find-endpoints

---
#创建角色对应操作
apiVersion: www.test.com:8001/v1
kind: ClusterRole
metadata:
  name: discov-endpoints
rules:
- apiGroups: [""]
  resources: ["endpoints"]
  verbs: ["get","list","watch"]

---
#给账号绑定角色
apiVersion: www.test.com:8001/v1
kind: ClusterRoleBinding
metadata:
  name: find-endpoints-discov-endpoints
roleRef:
  apiGroup: www.test.com:8001
  kind: ClusterRole
  name: discov-endpoints
subjects:
- kind: ServiceAccount
  name: find-endpoints
  namespace: user-model
````


2. 生成及测试是否成功创建权限
```shell
kubectl apply -f auth.yaml
kubectl get sa -n user-model
```
sa 指的是 Service Account 该条指令是查看是否存在该服务账号

3. 生成 k8s deployment yaml文件
   1. 使用刚刚构建的docker images构建 rpc和api 的k8s yaml
   ```shell
   goctl kube deploy -name user-api -namespace user-model -image user-api:v1 -o user-api-k8s-deployment.yaml -port 14001 --serviceAccount find-endpoints -nodePort 24000 -replicas 2 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100
   goctl kube deploy -name user-rpc -namespace user-model -image user-rpc:v1 -o user-rpc-k8s-deployment.yaml -port 14000 --serviceAccount find-endpoints -replicas 2 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100
   ```
   其中 nodePort 为暴露的服务端口，详细参数含义（https://go-zero.dev/docs/tutorials/cli/kube）

   2. 构建 svc 和 pod
   ```shell
   kubectl apply -f user-rpc-k8s-deployment.yaml
   kubectl apply -f user-api-k8s-deployment.yaml
   ```






