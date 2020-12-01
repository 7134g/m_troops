package middlewares

import (
	"m_troops/go/spider/Mcooly/middlewares/plugins"
	"m_troops/go/spider/Mcooly/work/model"
)

func InitializeMiddlewares(s *model.SpiderParams) {
	// 基础插件
	//plugins.MidderwarePanic(s)
	plugins.MidderwareLogger(s)    // 日志
	plugins.MidderwareHttpError(s) // 错误计数及重试
}
