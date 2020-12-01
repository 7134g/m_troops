package config

import (
	"m_troops/go/spider/Mcooly/setting"
	"m_troops/go/spider/Mcooly/work/common"
	"sync"
)

// 湖南省政府采购网
var (
	WorkGroup    sync.WaitGroup
	TABLENAME    = "ccgp"
	WORKNAME     = "ccgp"
	PATH         = setting.GetSpiderTaskFilePath(WORKNAME)
	TaskDataList *common.TaskFile
)

//var DataChan chan model.Data

var (
	UrlIndex = []string{
		"http://www.ccgp-hunan.gov.cn/mvc/getNoticeList4Web.do",
	}
	UrlDetail   = "http://www.ccgp-hunan.gov.cn/mvc/viewNoticeContent.do?noticeId=%s&area_id="
	HeaderPatch = map[string]string{
		"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
	}
)

// index Post params
var (
	// post params page
	PostBody = map[string]string{
		"nType":     "", // 类别, 空代表，全部
		"startDate": "2020-01-01",
		"endDate":   "2020-10-27",
		"page":      "1",  // 页数
		"pageSize":  "18", // 返回条数
	}

	NType = map[int][]string{
		1: {"prcmNotices"},                     //【招标/资审公告】
		2: {"invalidNotices", "modfiyNotices"}, //【招标公告澄清/更正】
		3: {"dealNotices", "endNotices"},       //【中标候选人公示】
		4: {"contractNotices"},                 //【中标结果公示】
	}
	MinYear = 2015
)
