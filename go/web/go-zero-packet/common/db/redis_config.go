package db

import (
	"context"
	"fmt"
	rdsV8 "github.com/go-redis/redis/v8"
	"time"
)

func InitRdsClient(addr, pwd string) *rdsV8.Client {
	client := rdsV8.NewClient(&rdsV8.Options{
		Addr:         addr,
		Password:     pwd,
		DB:           0,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		PoolTimeout:  2 * time.Minute,
		IdleTimeout:  10 * time.Minute,
		PoolSize:     1000,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("redis init failed. err:%v", err.Error()))
	}
	return client
}
