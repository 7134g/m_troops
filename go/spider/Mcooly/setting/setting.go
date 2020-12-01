package setting

import (
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
)

var (
	PROJECT_NAME        = "hn_crawler"
	PROJECT_ROOT        = GetPath()
	PROJECT_SPIDER_TASK = `\work\spiders\*\task\task.json`
)

// mysql 连接
//var MYSQLCONNDEFULT string = "root:123456@tcp(127.0.0.1:3306)/hnspider?charset=utf8"
var MYSQLCONNYAML = "root:123456@tcp(127.0.0.1:3306)/hnspider?charset=utf8"

var TIMERHOUR = 0
var TIMERMINUTE = 0

// ua
var USERAGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36"

// headers
var HEADERS = map[string]string{
	"Connection":                "keep-alive",
	"Pragma":                    "no-cache",
	"Cache-Control":             "no-cache",
	"Upgrade-Insecure-Requests": "1",
	"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
	"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"Accept-Encoding":           "gzip, deflate, br",
	"Accept-Language":           "zh-CN,zh;q=0.9",
}

// http
var (
	PROXY_HOST  = "127.0.0.1:5555"
	RETRYCOUNT  = 5                // http 重试最大次数
	TIMEOUT     = 30 * time.Second // 超时时间
	HTTPSETTING = http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   TIMEOUT, // 超时时间
			KeepAlive: TIMEOUT, // keepAlive 超时时间
		}).DialContext,
		MaxIdleConns:        100,              // 最大空闲连接数
		IdleConnTimeout:     90 * time.Second, // 空闲连接超时
		TLSHandshakeTimeout: 10 * time.Second, // TLS 握手超时
	}
)

// spider
var (
	DUPLICATE      = true
	REQUESTSMAX    = 100000 // 队列请求最大值
	GOROUTINECOUNT = 100
	LOGLEVEL       = logrus.DebugLevel
	//LOGLEVEL = logrus.InfoLevel
	//LOGLEVEL                           = logrus.ErrorLevel
	WAIT_TIME_DEFULT time.Duration = 0
	//WAIT_TIME_CCGP       time.Duration = 0
	//WAIT_TIME_HNBIDDING  time.Duration = 0
	//WAIT_TIME_HNZTB      time.Duration = 0
	//WAIT_TIME_REGULATORY time.Duration = 0
	WAIT_TIME_SERVEPLAT  time.Duration = 0
	CONCURRENT                         = 1               // 并发数
	CONTINUOUSREPETITION int32         = 100             // spider 连续错误次数最大值
	SPIDERERRORCOUNT     int32         = 50              // spider 连续睡眠次数最大值
	SPIDERSLEEPCOUNT                   = 180             // spider 无任务睡眠最大次数
	SPIDERSLEEPTIME                    = time.Second * 5 // spider 无任务睡眠一次时间
	//RULE                               = &colly.LimitRule{
	//	Delay:       WAIT_TIME_DEFULT,
	//	DomainGlob:  "*",
	//	Parallelism: CONCURRENT,
	//}
)

const (
	DONE        = 20001 // 已执行
	EXCUTEERROR = 50001 // 执行过程中异常, 重发url

	CTXASSERTERROR = -50002 // 断言异常
	CTXNILERROR    = -50003 // 空
	SPIDERRETRY    = -50004 // 重发
)
