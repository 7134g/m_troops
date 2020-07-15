package design_patterns

import (
	"fmt"
	"sync"
)

type Example struct {
	name string
}

var Singleton *Example
var once sync.Once

func GetObject(s string) *Example {
	once.Do(func() {
		Singleton = &Example{name: s}
	})
	return Singleton
}

func test() {
	obj := GetObject("once")
	fmt.Printf("%p %v \n", obj, obj.name)
	obj2 := GetObject("two")
	fmt.Printf("%p %v \n", obj2, obj2.name)
}
