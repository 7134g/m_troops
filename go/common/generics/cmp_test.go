package generics

import (
	"sync"
	"testing"
)

func TestUint(t *testing.T) {
	x := CmpMap[string, uint]{
		lock: sync.RWMutex{},
		body: map[string]uint{},
	}
	x.Set("a", uint(1))
	t.Log(x.Get("a"))
}
