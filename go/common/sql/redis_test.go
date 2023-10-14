package model

import "testing"

func TestRedis(t *testing.T) {
	Redis()
	_ = RedisClient.Close()
}
