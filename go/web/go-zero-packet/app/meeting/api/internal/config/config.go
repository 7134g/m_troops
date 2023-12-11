package config

import (
	"demo/common/db"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql db.Mysql
	//Sqlite db.SqliteClient
	Mongo db.MongoConfig
	Redis redis.RedisConf

	Auth struct { // JWT 认证需要的密钥和过期时间配置
		AccessSecret string
		AccessExpire int64
	}
}
