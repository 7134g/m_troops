package generics

import (
	"sync"
	"testing"
)

func TestSliceUint(t *testing.T) {
	x := sliceMap[string, []byte]{
		lock: sync.RWMutex{},
		body: map[string][]byte{},
	}
	x.Set("a", []byte("bbbbbbb"))
	t.Log(x.Get("a"))
}
