# 下载
    wget https://s3.amazonaws.com/bitly-downloads/nsq/nsq-1.1.0.linux-amd64.go1.10.3.tar.gz
# 解压
    tar -zxvf nsq-1.1.0.linux-amd64.go1.10.3.tar.gz
# 启动服务
    cd nsq-1.1.0.linux-amd64.go1.10.3/bin/
    nohup ./nsqlookupd > /dev/null 2>&1 &
    nohup ./nsqd --lookupd-tcp-address=127.0.0.1:4160 > /dev/null 2>&1 &
    nohup ./nsqadmin --lookupd-http-address=127.0.0.1:4161 > /dev/null 2>&1 &

# 验证
    curl -d 'hello world' 'http://127.0.0.1:4151/pub?topic=test'

# 实例
main.go
```go
package main

import (
    "github.com/gin-gonic/gin"
    "wages_service/servers"
    "wages_service/tasks"
)

var GinEngine *gin.Engine

func main() {

    // 运行 task
    tasks.SyncDataRun()

    // 运行 nsq
    servers.NsqRun()

    // 运行server
    servers.HttpRun(GinEngine)
}
```

server.go
```go
package servers

import (
    "fmt"
    "github.com/nsqio/go-nsq"
)

// 默认配置
const HOST  = "127.0.0.1:4150"
const TOPIC_NAME  = "test"
const CHANNEL_NAME  = "test-channel"

// 启动Nsq
func NsqRun()  {
    Consumer()
}

// nsq发布消息
func Producer(msgBody string) {
    // 新建生产者
    p, err := nsq.NewProducer(HOST, nsq.NewConfig())
    if err != nil {
        panic(err)
    }
    // 发布消息
    if err := p.Publish(TOPIC_NAME, []byte(msgBody)); err != nil {
        panic(err)
    }
}


// nsq订阅消息
type ConsumerT struct{}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
    fmt.Println(string(msg.Body))
    return nil
}

func Consumer() {
    c, err := nsq.NewConsumer(TOPIC_NAME, CHANNEL_NAME, nsq.NewConfig())   // 新建一个消费者
    if err != nil {
        panic(err)
    }
    c.AddHandler(&ConsumerT{})                                           // 添加消息处理
    if err := c.ConnectToNSQD(HOST); err != nil {            // 建立连接
        panic(err)
    }
}
```
client.go
```go
package main

import "wages_service/servers"

func main()  {
    // 发送消息到nsq
    servers.Producer("hello world!!!")
}

```

