package main

import (
	"fmt"
	"strconv"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// 订阅回调
func subCallBackFunc(client MQTT.Client, msg MQTT.Message) {
	time.Sleep(time.Second)
	fmt.Printf("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
}

// 连接MQTT服务
func connMQTT(broker, user, passwd string) (bool, MQTT.Client) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetUsername(user)
	opts.SetPassword(passwd)

	mc := MQTT.NewClient(opts)
	if token := mc.Connect(); token.Wait() && token.Error() != nil {
		return false, mc
	}

	return true, mc
}

// 订阅消息
func subscribe() {
	// sub的用户名和密码
	b, mc := connMQTT("tcp://127.0.0.1:1883", "sub", "aaabbbc")
	if !b {
		fmt.Println("sub connMQTT failed")
		return
	}
	mc.Subscribe("topic_tp", 0x00, subCallBackFunc)
}

// 发布消息
func publish() {
	// pub的用户名和密码
	b, mc := connMQTT("tcp://127.0.0.1:1883", "pub", "aaabbb")
	if !b {
		fmt.Println("pub connMQTT failed")
		return
	}

	n := 0
	for {
		msg := "Hello, this is publisher " + strconv.Itoa(n)
		mc.Publish("topic_tp", 0x00, true, msg)
		fmt.Println("Publish: msg is ", msg)
		time.Sleep(time.Millisecond * 500)
		n++
	}
}

func main() {
	subscribe()
	publish()
}
