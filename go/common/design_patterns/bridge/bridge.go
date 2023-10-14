// 桥接模式
package bridge

import "fmt"

type Bridge interface {
	Run()
}

type Chinese struct {
}

func (self *Chinese) Run() {
	fmt.Println("chinese")
}

type English struct {
}

func (self *English) Run() {
	fmt.Println("english")
}
