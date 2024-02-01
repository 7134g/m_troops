package generics

import (
	"cmp"
	"sync"
)

type dataType interface {
	cmp.Ordered
}

type CmpMap[K, D dataType] struct {
	lock sync.RWMutex

	body map[K]D
}

func (m *CmpMap[K, D]) Set(key K, value D) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.body[key] = value
}

func (m *CmpMap[K, D]) Get(key K) (D, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	value, exist := m.body[key]
	return value, exist
}

func (m *CmpMap[K, D]) Inc(key K, count D) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	value, exist := m.body[key]
	if exist {
		m.body[key] = count + value
	} else {
		m.body[key] = count
	}
}

func (m *CmpMap[K, D]) Del(key K) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	delete(m.body, key)
}

func (m *CmpMap[K, D]) Each(f func(key K, value D)) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for key, value := range m.body {
		f(key, value)
	}

}
