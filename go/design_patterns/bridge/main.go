// 桥接模式
package main

import "fmt"

type Bridge interface {
    Run()
}

type Chinese struct {

}

func (self *Chinese) Run()  {
    fmt.Println("chinese")
}

type English struct {

}

func (self *English) Run()  {
    fmt.Println("english")
}

func main() {
    var bridge Bridge
    ch := &Chinese{}
    en := &English{}

    bridge = ch
    bridge.Run()
    bridge = en
    bridge.Run()
}
