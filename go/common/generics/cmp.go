package generics

import (
	"cmp"
	"sync"
)

type dataType interface {
	cmp.Ordered
}

type CmpMap[D dataType] struct {
	lock sync.RWMutex

	body map[string]D
}

func (m *CmpMap[D]) Set(key string, value D) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.body[key] = value
}

func (m *CmpMap[D]) Get(key string) (D, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	value, exist := m.body[key]
	return value, exist
}
