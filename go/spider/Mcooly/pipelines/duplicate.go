package pipelines

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gocolly/colly"
	"spider/common/logs"
	"spider/common/util"
	"strings"
)

type Duplicate struct {
	ID         string `gorm:"column:id;index:duplicate"`
	SpiderName string `gorm:"column:spider_name;index:duplicate"`
}

func FilterCheckExist(r *colly.Request, spiderName string) bool {
	key := make([]byte, 0)
	key = append(key, []byte(r.URL.String())...)
	if r.Method != "GET" {
		key = append(key, []byte("\r\nPOST")...)
		_, body := util.AppendBytes(r.Body, &key)
		r.Body = strings.NewReader(body)
	}
	shaChunk := sha1.New()
	shaChunk.Write(key)
	id := hex.EncodeToString(shaChunk.Sum([]byte("")))

	// 查是否存在
	rs := Duplicate{ID: id, SpiderName: spiderName}
	if err := MysqlDB.Table("duplicate").First(&rs).Error; err == nil {
		return true
	}

	// 不存在写入
	err := MysqlDB.Table("duplicate").Create(rs).Error
	if err != nil {
		logs.Log.Error(err)
	}
	return false
}
