package pipelines

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"m_troops/go/spider/Mcooly/common/logs"
	"m_troops/go/spider/Mcooly/setting"
	"strings"
)

var (
	MysqlDB *gorm.DB
)

func InitMysqlDB() {
	db, err := gorm.Open("mysql", setting.MYSQLCONNYAML)
	if err != nil || db == nil {
		panic(err)
	}
	MysqlDB = db
	db.SingularTable(true)
	db.AutoMigrate(&Duplicate{})

}

func PushPipe(data interface{}, tableName string) {
	// add spider save
	switch tableName {
	case "regulatory":
		pending := data.(interface{})
		// 存在
		err := MysqlDB.Table(tableName).Where("url = ?", pending.Url).First(&pending).Error
		if err == nil {
			logs.Log.Info("xxxxx mysql exist")
			return
		}

		// 不存在
		save(pending, tableName)

	default:
		logs.Log.Error(errors.New("error mysql"))
		return
	}

}

func save(data interface{}, tableName string) {
	err := MysqlDB.Table(tableName).Create(data).Error
	if err != nil && !strings.Contains(err.Error(), "PRIMARY") && !strings.Contains(err.Error(), "Error 1062: ") {
		logs.Log.Error(err)
	}
}
