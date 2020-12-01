package processing

import (
	"m_troops/go/spider/Mcooly/common/logs"
	"regexp"
	"strings"
)

func ReCleanEmpty(target string) string {
	r, _ := regexp.Compile(`&nbsp;|\s+`)
	result := r.ReplaceAllString(target, "")
	return result
}

func ReCleanHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

func ReGetOneString(expr, target string) (string, bool) {
	re, err := regexp.Compile(expr)
	logs.Errors(err)
	result := re.FindStringSubmatch(target)
	if result == nil {
		return "", false
	}
	return result[1], true
}
