package db

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

func AcquireLock(ctx context.Context, client *redis.Client, key string, value string, expiration time.Duration) (bool, error) {
	result, err := client.SetNX(ctx, key, value, expiration).Result()
	if err != nil {
		return false, err
	}

	if result == true {
		return true, nil
	}

	return false, nil
}

func GetLock(ctx context.Context, client *redis.Client, key string, value string, expiration time.Duration) (bool, error) {

	timeoutCtx, cancel := context.WithTimeout(context.Background(), expiration)
	defer cancel()

	for {
		select {
		case <-timeoutCtx.Done():
			return false, errors.New("lock timeout")
		default:
			break
		}

		acquire, err := AcquireLock(ctx, client, key, value, expiration)
		if err != nil {
			return false, err
		}

		if !acquire {
			time.Sleep(time.Second)
			continue
		}

		return true, nil
	}
}

func ReleaseLock(ctx context.Context, client *redis.Client, key string) (bool, error) {
	result, err := client.Del(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if result == 1 {
		return true, nil
	}

	return false, nil
}
