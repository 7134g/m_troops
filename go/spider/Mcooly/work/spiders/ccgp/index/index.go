package index

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"m_troops/go/spider/Mcooly/common/util"
	"m_troops/go/spider/Mcooly/setting"
	"m_troops/go/spider/Mcooly/work/common/ctx"
	"m_troops/go/spider/Mcooly/work/common/shttp"
	"m_troops/go/spider/Mcooly/work/model"
	"m_troops/go/spider/Mcooly/work/spiders/ccgp/config"
	smodle "m_troops/go/spider/Mcooly/work/spiders/ccgp/model"
	"strconv"
)

var spider *model.SpiderParams

func SpiderStart(s *model.SpiderParams) {
	defer spider.ExcetePanic()
	spiderConfig(s)
	spider.Parse(nil, response)
	spider.Collector.OnResponse(DeepPage)

	postForm := config.PostBody
	postForm["startDate"] = fmt.Sprintf("%d-01-01", config.MinYear)
	postForm["endDate"] = fmt.Sprintf("%d-12-31", config.MinYear)
	for category, v := range config.NType {
		for _, v1 := range v {
			postForm["nType"] = v1
			//fmt.Println(category, postForm)
			newRequest(config.UrlIndex[0], postForm, category)
		}
	}

	// 阻塞到超过等待时间
	spider.Start()
}

func spiderConfig(s *model.SpiderParams) {
	spider = s
	//spider.Duplicate = false
	spider.InitCollector(config.WORKNAME + "-index ")
	extensions.Referer(spider.Collector)
}

// 初始函数
func newRequest(u string, postForm map[string]string, category int) {
	defer spider.ExcetePanic()
	body := shttp.CreatePostFormReader(postForm)
	r := shttp.NewPOSTRequest(util.UrlParse(u), body, config.HeaderPatch)

	// 初次添加必要信息
	ctx.PutInt(r.Ctx, "category", category)
	ctx.PutInt(r.Ctx, "page", 1)
	ctx.PutObject(r.Ctx, "postForm", postForm)
	ctx.PutObject(r.Ctx, "postForm_json", postForm)

	// 穿透
	ctx.PutBool(r.Ctx, "dont_filter", true)

	spider.AddNewRequest(r, 0)
}

func response(r *colly.Response) {
	defer spider.ExcetePanic()

	var ok bool
	category := ctx.GetInt(r.Ctx, "category")
	if category == setting.CTXASSERTERROR {
		spider.HttpRetry(r.Request)
		return
	}
	// parse
	var packJson interface{}
	err := json.Unmarshal(r.Body, &packJson)
	if err != nil {
		spider.SetSpiderCode(r.Ctx, setting.EXCUTEERROR)
		spider.HttpRetry(r.Request)
		return
	}

	firstJson := packJson.(map[string]interface{})

	// 最大页数
	max := ctx.GetInt(r.Ctx, "max")
	if ctx.IsError(max) {
		var maxFloat float64
		if maxFloat, ok = firstJson["total"].(float64); !ok {
			spider.HttpRetry(r.Request)
			return
		}
		max = (int(maxFloat) / 18) + 1
		ctx.PutInt(r.Ctx, "max", max)
	}

	// 题目和url
	var secendJson []interface{}
	var thirdJson map[string]interface{}
	if secendJson, ok = firstJson["rows"].([]interface{}); !ok {
		spider.SetSpiderCode(r.Ctx, setting.EXCUTEERROR)
		spider.HttpRetry(r.Request)
		return
	}

	// 空
	if len(secendJson) == 0 {
		return
	}

	for _, v := range secendJson {
		if thirdJson, ok = v.(map[string]interface{}); !ok {
			spider.SetSpiderCode(r.Ctx, setting.EXCUTEERROR)
			spider.HttpRetry(r.Request)
			return
		}
		title := thirdJson["NOTICE_TITLE"].(string)
		id := thirdJson["NOTICE_ID"].(float64)
		// 传输 detail 需要的数据
		oriUrl := r.Request.URL.String()
		oriData, _ := json.Marshal(ctx.GetObject(r.Ctx, "postForm_json"))
		ori := oriUrl + "\tPOST\t" + string(oriData)
		pending := smodle.Data{ID: strconv.Itoa(int(id)), Type: category, Title: title, OriUrl: ori}
		//config.DataChan <- pending
		config.TaskDataList.Append(pending)
	}

}

func DeepPage(r *colly.Response) {
	defer spider.ExcetePanic()

	max := ctx.GetInt(r.Ctx, "max")
	if ctx.IsError(max) {
		spider.HttpRetry(r.Request, "Response: DeepPage max")
		return
	}

	page := ctx.GetInt(r.Ctx, "page")
	if ctx.IsError(page) {
		spider.HttpRetry(r.Request, "Response: DeepPage page")
		return
	}

	if page != 1 {
		return
	}
	page += 1

	category := ctx.GetInt(r.Ctx, "category")
	if ctx.IsError(category) {
		spider.HttpRetry(r.Request, "Response: DeepPage category")
		return
	}

	// 一波
	for page <= max {
		// 翻页需要修改的内容
		resBody := ctx.GetObject(r.Ctx, "postForm").(map[string]interface{})
		form := make(map[string]string)
		for k, v := range resBody {
			form[k] = v.(string)
		}
		form["page"] = strconv.Itoa(page)

		// 发新请求
		body := shttp.CreatePostFormReader(form)

		newRequest := shttp.NewPOSTRequest(r.Request.URL, body, config.HeaderPatch)
		newRequest.Ctx = colly.NewContext()
		ctx.PutObject(newRequest.Ctx, "postForm_json", form)
		ctx.PutObject(newRequest.Ctx, "postForm", form)
		ctx.PutInt(newRequest.Ctx, "page", page)
		ctx.PutInt(newRequest.Ctx, "max", max)
		ctx.PutInt(newRequest.Ctx, "category", category)
		page = spider.AddNewRequest(newRequest, page)

		//if page == 10 {
		//	return
		//}
	}

}
