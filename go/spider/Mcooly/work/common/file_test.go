package common

import (
	"fmt"
	"m_troops/go/spider/Mcooly/setting"
	"testing"
)

func TestTaskFile_Load(t *testing.T) {
	path := setting.GetSpiderTaskFilePath("ccgp")
	b := NewTaskFile("test", path)
	b.Load()
	m := model.Data{
		ID:     "sss",
		Type:   1,
		Title:  "url",
		OriUrl: "title",
	}
	b.Append(m)
	m = model.Data{
		ID:     "sss",
		Type:   2,
		Title:  "url",
		OriUrl: "title",
	}
	b.Append(m)
	n := model.Data{}
	b.Bind(b.Pull(), &n)
	fmt.Println(b.Pop())
	b.Dump()
	fmt.Println(b.Data)
	fmt.Println(n)
}
