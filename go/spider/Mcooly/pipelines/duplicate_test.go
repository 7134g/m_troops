package pipelines

import (
	"fmt"
	"spider/common/logs"
	"spider/common/util"
	"spider/work/common/http"
	"strings"
	"testing"
)

var HeaderPatch map[string]string

func TestFilterCheckExistS(t *testing.T) {
	InitMysqlDB()
	logs.InitTheWorldLog()
	url := "https://www.baidu.com/s"
	r := http.NewGETRequest(util.UrlParse(url), HeaderPatch)
	fmt.Println(FilterCheckExist(r, "test"))

	p := http.NewPOSTRequest(util.UrlParse(url), strings.NewReader("name=123"), HeaderPatch)
	fmt.Println(FilterCheckExist(p, "test"))
	fmt.Println(FilterCheckExist(p, "test1"))

	p1 := http.NewPOSTRequest(util.UrlParse(url), strings.NewReader("name=456"), HeaderPatch)
	fmt.Println(FilterCheckExist(p1, "test1"))
}
