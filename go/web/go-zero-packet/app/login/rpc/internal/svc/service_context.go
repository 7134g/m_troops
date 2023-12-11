package svc

import (
	"demo/app/login/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config     config.Config
	RedisClint *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		RedisClint: redis.MustNewRedis(c.Redis.RedisConf),
	}
}
