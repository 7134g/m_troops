package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisConfig struct {
	Host string
	Type string
	Pass string
}

func InitRdsClient(addr, pwd string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     pwd,
		DB:           0,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		PoolTimeout:  2 * time.Minute,
		PoolSize:     1000,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("redis init failed. err:%v", err.Error()))
	}
	return client
}
