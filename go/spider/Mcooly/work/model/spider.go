/*
dont_filter      穿透
spider_code      爬虫响应状态
*/

package model

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"m_troops/go/spider/Mcooly/common/logs"
	"m_troops/go/spider/Mcooly/common/proxy"
	"m_troops/go/spider/Mcooly/common/util"
	"m_troops/go/spider/Mcooly/pipelines"
	"m_troops/go/spider/Mcooly/setting"
	"m_troops/go/spider/Mcooly/work/common/ctx"
	"m_troops/go/spider/Mcooly/work/common/shttp"
	"m_troops/go/spider/Mcooly/work/pool"

	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type SpiderParams struct {
	SpiderName           string                                // 爬虫名字
	Collector            *colly.Collector                      // 爬虫控制器
	Pool                 *pool.Queue                           // 任务队列
	WG                   *sync.WaitGroup                       // 任务锁
	ConcurrentCount      int                                   // 并发数
	Duplicate            bool                                  // 去重
	continuousRepetition int32                                 // 连续重复任务数
	Proxy                func(*http.Request) (*url.URL, error) // 开启代理
	ProxyValue           string                                // 代理ip值
	HttpTimeout          time.Duration                         // 连接超时时间
	SpiderStatus         bool                                  // 爬虫状态
	WaitTime             time.Duration                         // 等待时间

	// 不能配置字段
	synchronous         sync.Mutex // 爬虫锁
	statusLock          sync.Mutex // 爬虫开关锁
	errorCount          int32      // 连续错误次数
	retryCount          int32      // 重试上限
	spiderSleepMaxCount int        // 爬虫最长等待时间 默认为 50秒，检查100次, 间隔 500ms
}

// 初始化爬虫
func (s *SpiderParams) InitCollector(name string) {
	var err error
	// 爬虫基本信息
	s.SpiderStatus = true
	s.SpiderName = name             // 设置爬虫名字
	s.Duplicate = setting.DUPLICATE // 默认开启去重
	if s.WaitTime == 0 {
		s.WaitTime = setting.WAIT_TIME_DEFULT
	}

	// 设置 config
	s.SpiderReSetting()
	s.HttpReSetting()

	// 创建队列
	s.Pool, err = pool.New(s.ConcurrentCount, &queue.InMemoryQueueStorage{MaxSize: setting.REQUESTSMAX}, setting.REQUESTSMAX)
	logs.Errors(err)

	// 设置爬虫名字
	s.Pool.SpiderName = name
}

// 任务开启，阻塞到超过等待时间
func (s *SpiderParams) Start() {
	defer s.Stop() // 停止爬虫

	logs.Log.Info(s.SpiderName + "爬虫启动")

	// 执行爬虫
	for {
		// 爬虫是否已经停止
		if !s.SpiderStatus {
			logs.Log.Info("[STOP] ", s.SpiderName, " 爬虫已停止")
			return
		}

		// 获取任务池数量
		n := s.GetQueueCount()

		// 没有任务,退出
		if s.SpiderRunErrorGet() > setting.SPIDERSLEEPCOUNT && n == 0 {
			logs.Log.Warning(s.SpiderName + " 没有任务在执行, 超过爬虫最长等待时间")
			return
		}

		// 没有任务,等待,计数
		if n == 0 {
			time.Sleep(setting.SPIDERSLEEPTIME + s.WaitTime)
			msg := fmt.Sprintf("执行等待次数[%d/%d] %s 爬虫没有任务在执行,等待新任务",
				s.SpiderRunErrorGet(), setting.SPIDERSLEEPCOUNT, s.SpiderName)
			//logs.Log.Info("[WAIT] 执行等待次数[", s.SpiderRunErrorGet(), "/", setting.SPIDERSLEEPCOUNT, "] ", s.SpiderName+" 爬虫没有任务在执行,等待新任务")
			logs.Log.Warning(msg)
			s.SpiderRunErrorAdd()
			continue
		}

		// 任务量太低，等待一会
		if n < 10 {
			time.Sleep(setting.SPIDERSLEEPTIME)
		}

		//s.LockQueueRun()

		// 开始执行任务
		s.SpiderRunErrorZero()
		logs.Log.Info(fmt.Sprintf("开始执行 %s 队列中 %d 个任务", s.SpiderName, n))
		err := s.Pool.Run(s.Collector)
		logs.Errors(err)
	}
}

func (s *SpiderParams) SpiderRunErrorZero() {
	s.statusLock.Lock()
	defer s.statusLock.Unlock()
	s.spiderSleepMaxCount = 0
}

func (s *SpiderParams) SpiderRunErrorAdd() {
	s.statusLock.Lock()
	defer s.statusLock.Unlock()
	s.spiderSleepMaxCount += 1
}

func (s *SpiderParams) SpiderRunErrorGet() int {
	s.statusLock.Lock()
	defer s.statusLock.Unlock()
	return s.spiderSleepMaxCount
}

// 任务停止
func (s *SpiderParams) Stop() {
	s.statusLock.Lock()
	defer s.statusLock.Unlock()

	// 已经停止了
	if !s.SpiderStatus {
		return
	}

	// 清理最后可能存在的任务
	if !s.Pool.IsEmpty() {
		_ = s.Pool.Run(s.Collector)
	}

	logs.Log.Info(s.SpiderName + "爬虫结束")
	s.WG.Done() // 结束爬虫
	s.SpiderStatus = false
}

// 加载执行逻辑
func (s *SpiderParams) Parse(request colly.RequestCallback, response colly.ResponseCallback) {
	if request != nil {
		s.Collector.OnRequest(request)
	}

	if response != nil {
		s.Collector.OnResponse(response)
	}

}

// spider 重设
func (s *SpiderParams) SpiderReSetting() {
	// 重设并发数
	newRule := &colly.LimitRule{
		Delay:       s.WaitTime,
		DomainGlob:  "*",
		Parallelism: setting.CONCURRENT,
	}
	if newRule.Parallelism != s.ConcurrentCount {
		newRule.Parallelism = s.ConcurrentCount
	}

	err := s.Collector.Limit(newRule)
	logs.Errors(err, "colly.LimitRule error")
}

// Http 重设
func (s *SpiderParams) HttpReSetting() {
	var p func(*http.Request) (*url.URL, error)

	// 设置代理
	if s.Proxy != nil {
		p = s.Proxy
	} else {
		p = nil
		//p = proxy.ProxyFiddle()
	}

	hr := http.Transport{
		Proxy: p, // 代理
		DialContext: (&net.Dialer{
			Timeout:   s.HttpTimeout, // 超时时间
			KeepAlive: s.HttpTimeout, // keepAlive 超时时间
		}).DialContext, // 超时控制
		MaxIdleConns:        100,              // 最大空闲连接数
		IdleConnTimeout:     90 * time.Second, // 空闲连接超时
		TLSHandshakeTimeout: 30 * time.Second, // TLS 握手超时
	}

	s.Collector.WithTransport(&hr)
}

// 对重复url 计数
func (s *SpiderParams) SpiderContinuousRepetition() {
	s.synchronous.Lock()
	defer s.synchronous.Unlock()
	s.continuousRepetition += 1
	if s.continuousRepetition > setting.CONTINUOUSREPETITION {
		s.Stop()
	}
}

// http请求重试，并计数
func (s *SpiderParams) HttpRetry(r *colly.Request, errs ...interface{}) {
	// 爬虫是否已经停止
	if !s.SpiderStatus {
		return
	}
	retryCount := ctx.GetInt(r.Ctx, "retryCount")
	if retryCount == setting.CTXNILERROR {
		retryCount = 1
		ctx.PutInt(r.Ctx, "retryCount", retryCount)
	}

	httpCode := ctx.GetInt(r.Ctx, "httpCode")
	if httpCode == setting.CTXNILERROR {
		httpCode = 500
		ctx.PutInt(r.Ctx, "httpCode", httpCode)
	}

	if retryCount <= setting.RETRYCOUNT {
		// 判断该连接是否重试次数已满
		ctx.PutInt(r.Ctx, "retryCount", retryCount+1)
		ctx.PutInt(r.Ctx, "spider_code", setting.SPIDERRETRY)
		s.SpiderRunErrorZero() // 重置沉睡最大值
		msg := fmt.Sprintf("重试 %d 次 ,httpCode: %d, URL: %s Msg: ", retryCount, httpCode, r.URL.String())
		newError := make([]interface{}, 0)
		newError = append(newError, msg)
		newError = append(newError, errs...)
		logs.Log.Error(newError...)
		err := r.Retry()
		pe := ctx.GetBool(r.Ctx, "proxy_error")
		if !pe {
			ctx.PutBool(r.Ctx, "proxy_error", false)
			s.SpiderErrorCount(err)
		}

	} else {
		logs.Log.Error("达到重试最大值 ", r.URL.String())
	}
}

// spider错误计数
func (s *SpiderParams) SpiderErrorCount(errs ...interface{}) {
	s.synchronous.Lock()
	defer s.synchronous.Unlock()
	s.errorCount += 1

	var msgError []interface{}
	msg := fmt.Sprintf("%s 当前错误次数[%d/50]", s.SpiderName, s.errorCount)
	msgError = append(msgError, msg)
	msgError = append(msgError, errs...)
	logs.Log.Error(msgError...)

	if s.errorCount >= setting.SPIDERERRORCOUNT {
		logs.Errors(s.SpiderName + "爬虫连续错误次数已达上限， 退出爬虫")
		s.Stop()
	}
}

// spider错误置零
func (s *SpiderParams) SpiderErrorZero(r *colly.Response) {
	s.synchronous.Lock()
	defer s.synchronous.Unlock()
	// spider 计数 成功获取到新的响应，错误置零
	if s.errorCount != 0 {
		s.errorCount = 0
	}
	// 连续重复置零
	s.continuousRepetition = 0
	// 状态改为真
	ctx.PutInt(r.Ctx, "spider_code", setting.DONE)

}

// 添加新的url任务
func (s *SpiderParams) AddNewUrl(u string, page int) int {
	//n := s.GetQueueCount()
	//if n >= setting.REQUESTSMAX {
	//	time.Sleep(setting.SPIDERSLEEPTIME)
	//	return page
	//}

	// 放弃当前请求
	r := shttp.NewGETRequest(util.UrlParse(u), nil)
	if s.CheckDuplicate(r, s.SpiderName) {
		logs.Log.Debug("过滤当前请求: ", r.URL.String())
		s.SpiderContinuousRepetition()
		return page + 1
	}

	// 强行等待里面的goroutine 少于 配置的最大值
	//s.LockQueueMax()

	err := s.Pool.AddURL(u)
	if err != nil {
		s.SpiderErrorCount(s.SpiderName, "Pool.AddURL", err)
	}
	return page + 1
}

// 添加新的请求任务
func (s *SpiderParams) AddNewRequest(r *colly.Request, page int) int {
	//n := s.GetQueueCount()
	//if n >= setting.REQUESTSMAX {
	//	time.Sleep(setting.SPIDERSLEEPTIME)
	//	return page
	//}

	// 放弃当前请求,过滤第一个请求
	if s.CheckDuplicate(r, s.SpiderName) {
		logs.Log.Debug(s.SpiderName, " 过滤当前请求: ", r.URL.String())
		s.SpiderContinuousRepetition()
		return page + 1
	}

	// 强行等待里面的goroutine 少于 配置的最大值
	//s.LockQueueMax()

	err := s.Pool.AddRequest(r)
	if err != nil {
		s.SpiderErrorCount(s.SpiderName, "Pool.AddNewRequest", err)
		return page
	}
	return page + 1
}

// 查看队列数据量
func (s *SpiderParams) GetQueueCount() int {
	s.synchronous.Lock()
	defer s.synchronous.Unlock()

	n := s.Pool.GetTaskCount()
	return n
}

func (s *SpiderParams) GetQueueActiveCount() int {
	s.synchronous.Lock()
	defer s.synchronous.Unlock()

	n := s.Pool.GetActiveCount()
	return n
}

//func (s *SpiderParams) LockQueueMax() {
//	var count int
//	for {
//		count = s.GetQueueCount()
//		if count > setting.GOROUTINECOUNT {
//			logs.Log.Info("[WAIT] ", s.SpiderName+" 当前任务超过最大值，开始沉睡等待")
//			time.Sleep(setting.SPIDERSLEEPTIME)
//		} else {
//			//logs.Log.Debug(s.SpiderName + " 队列值小于最大值，继续执行任务...")
//			break
//		}
//	}
//
//}
//
//func (s *SpiderParams) LockQueueRun() {
//	var count int
//	for {
//		count = s.GetQueueActiveCount()
//		if count > setting.GOROUTINECOUNT {
//			logs.Log.Info("[WAIT] ", s.SpiderName+" 当活动中线程超过最大值， 开始沉睡等待现有执行完毕")
//			time.Sleep(setting.SPIDERSLEEPTIME)
//		} else {
//			//logs.Log.Debug(s.SpiderName + " 队列值小于最大值，继续执行任务...")
//			break
//		}
//	}
//}

// 去重
func (s *SpiderParams) CheckDuplicate(r *colly.Request, name string) bool {
	// 是否开启去重
	if !s.Duplicate {
		return false
	}

	// 重试
	spiderCode := ctx.GetInt(r.Ctx, "spider_code")
	if spiderCode == setting.SPIDERRETRY {
		return false
	}

	// 穿透
	dontFilter := ctx.GetBool(r.Ctx, "dont_filter")
	if dontFilter {
		ctx.PutBool(r.Ctx, "dont_filter", false)
		return false
	}

	// 去重
	return pipelines.FilterCheckExist(r, name)
}

// panic
func (s *SpiderParams) ExcetePanic() {
	if err := recover(); err != nil {
		logs.Errors(fmt.Sprintf("%s 爬虫程序 panic: ", s.SpiderName), err)
	}
}

func (s *SpiderParams) SetSpiderCode(c *colly.Context, sgin int) {
	ctx.PutInt(c, "spider_code", sgin)
}

func (s *SpiderParams) AddProxy() {
	s.Proxy, s.ProxyValue = proxy.ProxyPoolGet()
	s.HttpReSetting()
}

func (s *SpiderParams) ChangeProxy() {
	proxy.PoolDelete(s.ProxyValue)
	s.Proxy, s.ProxyValue = proxy.ProxyPoolGet()
}
