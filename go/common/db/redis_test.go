package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"testing"
	"time"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,
	})
}

func TestReleaseLock(t *testing.T) {
	key := "mylock"                // 锁的键名
	value := "myvalue"             // 锁的值（可以是唯一标识符）
	expiration := 10 * time.Second // 锁的超时时间

	ctx := context.Background()

	defer client.Close()
	go func() {
		acquiredis, err := GetLock(ctx, client, key, value, expiration)
		if err != nil {
			log.Fatal(err)
		}

		if acquiredis {
			fmt.Println("ok lock line 1", time.Now())
			time.Sleep(time.Second * 5)
			ok, err := ReleaseLock(ctx, client, key)
			fmt.Println("release lock line 1", ok, err)
		}
	}()

	go func() {
		//time.Sleep(time.Second)
		acquiredis, err := GetLock(ctx, client, key, value, expiration)
		if err != nil {
			log.Fatal(err)
		}

		if acquiredis {
			fmt.Println("ok lock line 2", time.Now())
			time.Sleep(time.Second * 3)
			ok, err := ReleaseLock(ctx, client, key)
			fmt.Println("release lock line 2", ok, err)
		}
	}()

	time.Sleep(time.Minute)
}
