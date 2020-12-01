package ccgp

import (
	"m_troops/go/spider/Mcooly/common/logs"
	"m_troops/go/spider/Mcooly/work/common"
	"m_troops/go/spider/Mcooly/work/spiders/ccgp/config"
	"m_troops/go/spider/Mcooly/work/spiders/ccgp/detail"
	"m_troops/go/spider/Mcooly/work/spiders/ccgp/index"
	"sync"
)

func Run(wg *sync.WaitGroup) {
	wg.Add(1)
	//config.DataChan = make(chan model.Data, 1000)
	config.TaskDataList = common.NewTaskFile(config.WORKNAME, config.PATH)
	// 创建爬虫 添加 WorkGroup
	spiderIndex := common.NewSpider(&config.WorkGroup)
	spiderDetail := common.NewSpider(&config.WorkGroup)

	// 爬虫执行
	go index.SpiderStart(spiderIndex)
	go detail.SpiderStart(spiderDetail)

	// 阻塞,等待任务完成
	spiderIndex.Collector.Wait()
	spiderDetail.Collector.Wait()
	config.WorkGroup.Wait()
	wg.Done()
	config.TaskDataList.Dump()
	logs.Log.Info(config.WORKNAME, "执行爬虫完成")
}
