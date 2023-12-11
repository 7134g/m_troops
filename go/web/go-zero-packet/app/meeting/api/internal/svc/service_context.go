package svc

import (
	"demo/app/meeting/api/internal/config"
	"demo/common/db"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	MysqlDB *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	mysql, err := db.ConnectMysql(c.Mysql)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:  c,
		MysqlDB: mysql,
	}
}
