package detail

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"m_troops/go/spider/Mcooly/common/processing"
	"m_troops/go/spider/Mcooly/pipelines"
	"m_troops/go/spider/Mcooly/setting"
	"m_troops/go/spider/Mcooly/work/common/ctx"
	"m_troops/go/spider/Mcooly/work/common/shttp"
	"m_troops/go/spider/Mcooly/work/model"
	"m_troops/go/spider/Mcooly/work/spiders/ccgp/config"
	smodel "m_troops/go/spider/Mcooly/work/spiders/ccgp/model"
	"net/url"
	"strings"
	"time"
)

// http://bidding.hunan.gov.cn/history/notice/ff8080816e24ff09016ed92c3f2b77f8?isdetail=1
var spider *model.SpiderParams

func SpiderStart(s *model.SpiderParams) {
	defer spider.ExcetePanic()
	spiderConfig(s)
	spider.Parse(nil, response)
	spider.Collector.OnResponse(save)
	go NewRequest() // 添加任务

	spider.Start() // 执行爬虫
}

func spiderConfig(s *model.SpiderParams) {
	spider = s
	spider.InitCollector(config.WORKNAME + "-detail ")
}

func NewRequest() {
	for {
		//Data, ok := <-config.DataChan

		data := config.TaskDataList.Pull()
		if data == nil {
			time.Sleep(setting.SPIDERSLEEPTIME)
			continue
		}
		Data := smodel.Data{}
		config.TaskDataList.Bind(data, &Data)

		if !spider.SpiderStatus {
			return
		}

		u := fmt.Sprintf(config.UrlDetail, Data.ID)
		targetUrl, err := url.Parse(u)
		if err != nil {
			spider.SpiderErrorCount("NewRequest", err, u)
		}

		// add head
		r := shttp.NewGETRequest(targetUrl, config.HeaderPatch)

		// need to save data
		ctx.PutInt(r.Ctx, "type", Data.Type)
		ctx.PutString(r.Ctx, "title", Data.Title)
		ctx.PutString(r.Ctx, "ori_url", Data.OriUrl)

		// new request
		spider.AddNewRequest(r, 0)
	}
}

func response(r *colly.Response) {
	html := string(r.Body)
	category := ctx.GetInt(r.Ctx, "type")
	if category == setting.CTXNILERROR || category == setting.CTXASSERTERROR {
		spider.HttpRetry(r.Request, "type is nil ", r.Request.URL.String())
		return
	}
	title := ctx.GetString(r.Ctx, "title")
	if title == "" {
		spider.HttpRetry(r.Request, "title is nil ", r.Request.URL.String())
		return
	}

	// xpath
	//c := processing.ReCleanHtml(html)
	//fmt.Println(c)
	doc, err := htmlquery.Parse(strings.NewReader(html))
	if err != nil {
		spider.SetSpiderCode(r.Ctx, setting.EXCUTEERROR)
		spider.SpiderErrorCount("doc xpath error: ", err, r.Request.URL.String())
		return
	}

	// 内容
	var content, part string
	contentNodes := htmlquery.Find(doc, `//body/*`)
	for _, node := range contentNodes {
		part = htmlquery.OutputHTML(node, true)
		part = processing.ReCleanHtml(part)
		part = processing.ReCleanEmpty(part)
		content += part + "\n"
	}

	if content == "" {
		// 不是代码问题
		//msg := fmt.Sprintf(" [OriUrl: %s] [ErrUrl: %s]", ctx.GetString(r.Ctx, "ori_url"), r.Request.URL.String())
		//spider.SpiderErrorCount("content is nothing: ", err, msg)
		return
	}

	// 构造数据包
	result := smodel.Ccgp{
		Type:    category,
		Title:   title,
		Url:     r.Request.URL.String(),
		Html:    html,
		Content: content,
	}

	r.Ctx.Put("result", result)

}

func save(r *colly.Response) {
	result := ctx.GetObject(r.Ctx, "result")
	if result != nil {
		pipelines.PushPipe(result.(smodel.Ccgp), config.TABLENAME)
	}
}
