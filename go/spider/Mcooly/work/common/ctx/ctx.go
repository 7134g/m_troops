package ctx

import (
	"github.com/gocolly/colly"
	"m_troops/go/spider/Mcooly/setting"
)

func IsError(n int) bool {
	return n == setting.CTXNILERROR || n == setting.CTXASSERTERROR
}

func PutInt(ctx *colly.Context, key string, value int) {
	ctx.Put(key, float64(value))
}

func GetInt(ctx *colly.Context, key string) int {
	var value float64
	var ok bool

	valueCtx := ctx.GetAny(key)
	// 没有赋值
	if valueCtx == nil {
		return setting.CTXNILERROR
	}

	// 断言异常
	if value, ok = valueCtx.(float64); !ok {
		PutInt(ctx, "spider_code", setting.CTXASSERTERROR)
		return setting.CTXASSERTERROR
	}

	return int(value)
}

func PutString(ctx *colly.Context, key string, value string) {
	ctx.Put(key, value)
}

func GetString(ctx *colly.Context, key string) string {
	return ctx.Get(key)
}

func PutObject(ctx *colly.Context, key string, object interface{}) {
	ctx.Put(key, object)
}

func GetObject(ctx *colly.Context, key string) interface{} {
	return ctx.GetAny(key)
}

func PutBool(ctx *colly.Context, key string, value bool) {
	ctx.Put(key, value)
}

func GetBool(ctx *colly.Context, key string) bool {
	var value bool
	var ok bool

	valueCtx := ctx.GetAny(key)
	// 没有赋值
	if valueCtx == nil {
		PutInt(ctx, "spider_code", setting.CTXNILERROR)
		return false
	}

	// 断言异常
	if value, ok = valueCtx.(bool); !ok {
		PutInt(ctx, "spider_code", setting.CTXASSERTERROR)
		return false
	}

	return value
}
