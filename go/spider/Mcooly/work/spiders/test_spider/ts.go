package test_spider

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"m_troops/go/spider/Mcooly/common/proxy"
	"m_troops/go/spider/Mcooly/common/util"
	"m_troops/go/spider/Mcooly/work/common"
	"m_troops/go/spider/Mcooly/work/common/ctx"
	"m_troops/go/spider/Mcooly/work/common/shttp"
	"m_troops/go/spider/Mcooly/work/model"
	"sync"
	"time"
)

var (
	WorkGroup sync.WaitGroup
	spider    *model.SpiderParams
	testURL   = "http://whatip.org/"
	//testURL = "https://www.hnsggzy.com/queryContent-jygk.jspx?title=&origin=&inDates=&channelId=845&ext=%E6%8B%9B%E6%A0%87%2F%E8%B5%84%E5%AE%A1%E5%85%AC%E5%91%8A&beginTime=&endTime="
)

func TSpider() {
	n := common.NewSpider(&WorkGroup)
	SpiderStart(n)
}

func SpiderStart(s *model.SpiderParams) {
	defer spider.ExcetePanic()
	spiderConfig(s)
	spider.Parse(nil, response)
	extensions.Referer(spider.Collector)

	parseUrl := util.UrlEncode(testURL)
	newUrl := util.UrlParse(parseUrl)
	r := shttp.NewGETRequest(newUrl, nil)
	// 穿透
	ctx.PutBool(r.Ctx, "dont_filter", true)

	spider.AddNewRequest(r, 0)

	// 阻塞到超过等待时间
	spider.Start()
}

func spiderConfig(s *model.SpiderParams) {
	spider = s
	s.Proxy, _ = proxy.ProxyPoolGet()
	spider.Collector.SetRequestTimeout(60 * time.Second)
	spider.ConcurrentCount = 1
	_ = spider.Collector.Limit(&colly.LimitRule{
		Delay:       0,
		DomainGlob:  "*",
		Parallelism: spider.ConcurrentCount,
	})
	spider.InitCollector("测试")
}

func response(r *colly.Response) {
	body := string(r.Body)
	fmt.Println(body)
}
