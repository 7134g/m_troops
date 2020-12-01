package common

import (
	"github.com/gocolly/colly"
	"m_troops/go/spider/Mcooly/middlewares"
	"m_troops/go/spider/Mcooly/setting"
	"m_troops/go/spider/Mcooly/work/model"
	"sync"
)

func NewSpider(wg *sync.WaitGroup) *model.SpiderParams {
	spider := model.SpiderParams{WG: wg}
	spider.WG.Add(1)

	// 创建控制器
	spider.Collector = colly.NewCollector(
		colly.IgnoreRobotsTxt(), // robot
		colly.Async(false),      // 异步
		colly.AllowURLRevisit(), // 关闭colly的去重,但程序关闭，则关闭，废弃
	)
	spider.ConcurrentCount = setting.CONCURRENT
	spider.Collector.SetRequestTimeout(setting.TIMEOUT)
	// 基础插件
	middlewares.InitializeMiddlewares(&spider)

	return &spider
}
