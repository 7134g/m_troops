// 单例模式
package main

import (
    "fmt"
    "sync"
)

type Alone struct {
    name string
}

var First *Alone
var once sync.Once

func GetObjectAlone(s string) *Alone {
    once.Do(func() {
        First = &Alone{name: s}
    })
    return First
}

func main() {
    obj1 := GetObjectAlone("first")
    fmt.Println(obj1.name)
    obj2 := GetObjectAlone("second")
    fmt.Println(obj2.name)
}
