package util

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	u := ""
	p := UrlParse(u)
	fmt.Println(u)
	fmt.Println(p.String())

}

func TestEscapeQuery(t *testing.T) {
	u := "https://www.hnsggzy.com/queryContent_1-jygk.jspx?title=&origin=&inDates=&channelId=845&ext=招标%2F资审公告&beginTime=&endTime="
	fmt.Println(u)
	fmt.Println(UrlEncode(u))

}
