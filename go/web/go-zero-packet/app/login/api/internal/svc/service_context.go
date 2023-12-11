package svc

import (
	"demo/app/login/api/internal/config"
	"demo/app/login/rpc/login_client"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	RpcLogin   login_client.Login
	RedisClint *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		RpcLogin:   login_client.NewLogin(zrpc.MustNewClient(c.RpcAdminCfg)),
		RedisClint: redis.MustNewRedis(c.Redis),
	}
}
