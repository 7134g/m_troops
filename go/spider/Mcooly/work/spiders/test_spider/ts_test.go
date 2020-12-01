package test_spider

import (
	"m_troops/go/spider/Mcooly/common/logs"
	"m_troops/go/spider/Mcooly/pipelines"

	"testing"
)

func init() {
	pipelines.InitMysqlDB()
	logs.InitTheWorldLog()
}

func TestTSpider(t *testing.T) {
	TSpider()
}
