package db

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"time"
)

func overwriteGormLogMode(mode string) logger.LogLevel {
	switch mode {
	case "dev":
		return logger.Info
	case "test":
		return logger.Warn
	case "prod":
		return logger.Error
	case "silent":
		return logger.Silent
	default:
		return logger.Info
	}
}

type Mysql struct {
	Path          string // 服务器地址
	Port          int    `json:",default=3306"`                                               // 端口
	Config        string `json:",default=charset%3Dutf8mb4%26parseTime%3Dtrue%26loc%3DLocal"` // 高级配置
	Dbname        string // 数据库名
	Username      string // 数据库用户名
	Password      string // 数据库密码
	MaxIdleConns  int    `json:",default=10"`                               // 空闲中的最大连接数
	MaxOpenConns  int    `json:",default=10"`                               // 打开到数据库的最大连接数
	LogMode       string `json:",default=dev,options=dev|test|prod|silent"` // 是否开启Gorm全局日志
	LogZap        bool   // 是否通过zap写入日志文件
	SlowThreshold int64  `json:",default=1000"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + fmt.Sprintf("%d", m.Port) + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetGormLogMode() logger.LogLevel {
	return overwriteGormLogMode(m.LogMode)
}

func ConnectMysql(m Mysql) (*gorm.DB, error) {
	masterURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&timeout=10s&readTimeout=5s&writeTimeout=5s",
		m.Username, m.Password, m.Path, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(masterURL), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(m.GetGormLogMode()),
	})
	if err != nil {
		logx.Error(context.Background(), "open mysql fail")
		panic(err)
	}
	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{mysql.Open(masterURL)},
		//Replicas: []gorm.Dialector{mysql.Open(slaveURL)},
		Policy: dbresolver.RandomPolicy{}}).
		// should use go1.15
		// SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(4 * time.Hour).
		SetMaxIdleConns(8).
		SetMaxOpenConns(58))
	if err != nil {
		logx.Error(context.Background(), "open mysql fail", err)
		return nil, err
	}
	return db, nil
}
