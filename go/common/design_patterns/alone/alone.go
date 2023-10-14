// 单例模式
package alone

import (
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
