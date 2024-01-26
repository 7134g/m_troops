package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

// 自定义命令监视器
var monitor = &event.CommandMonitor{
	//Started: func(ctx context.Context, event *event.CommandStartedEvent) {
	//	fmt.Printf("Command started: %s\n", event.Command)
	//},
	//Failed: func(ctx context.Context, event *event.CommandFailedEvent) {
	//	fmt.Printf("Command failed: %s\n", event.Failure)
	//},
	Succeeded: func(ctx context.Context, event *event.CommandSucceededEvent) {
		if event.Duration > time.Millisecond*100 {
			//logger.Slow(ctx, fmt.Sprintf("%d Executed command: %s \n", event.Duration.Milliseconds(), event.Reply))
			//fmt.Printf("%d Executed command: %s \n", event.Duration.Milliseconds(), event.Reply)
			log.Println(fmt.Sprintf("%d Executed command: %s \n", event.Duration.Milliseconds(), event.Reply))
		}
	},
}

type MongodbConfig struct {
	URI    string `json:"uri"`
	DBName string `json:"db_name"`
}

func NewMongoDBInstance(config *MongodbConfig) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 建立连接
	client, err := mongo.Connect(ctx,
		options.Client().
			// 连接地址
			ApplyURI(config.URI).
			SetMonitor(monitor).SetTimeout(time.Second*10).
			// 设置连接数
			SetMaxPoolSize(20))
	if err != nil {
		//log.Println(err)
		return nil, err
	}
	// 测试连接
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		//log.Println(err)
		return nil, err
	}

	return client, nil
}
