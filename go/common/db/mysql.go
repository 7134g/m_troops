package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"sync"
	"time"
)

var (
	DbUserName string
	DbPassWord string
	DbAddr     string
	DbName     string
	mysqlConn  string

	mysqlDB *gorm.DB
	dbConf  = &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	once sync.Once
)

type TestTable struct {
}

func InitMysqlConnect() {
	//mysqlConn = "root:123456@tcp(127.0.0.1:3306)/dataBaseName?parseTime=true"
	once.Do(func() {
		mysqlConn = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", DbUserName, DbPassWord, DbAddr, DbName)
		db, err := gorm.Open(mysql.Open(mysqlConn), dbConf)
		if err != nil || db == nil {
			panic(err)
		}
		mysqlDB = db

		_ = db.AutoMigrate(&TestTable{})

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("连接数据库失败: %v", err)
		}

		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConns 设置打开数据库连接的最大数量
		sqlDB.SetMaxOpenConns(100)
		// SetConnMaxLifetime 设置了连接可复用的最大时间
		sqlDB.SetConnMaxLifetime(time.Hour)
	})

}
