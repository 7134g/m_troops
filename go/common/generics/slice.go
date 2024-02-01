package generics

import "sync"

type sliceType interface {
	[]byte | []string | []int
}

type sliceMap[K dataType, D sliceType] struct {
	lock sync.RWMutex

	body map[K]D
}

func (m *sliceMap[K, D]) Set(key K, value D) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.body[key] = value
}

func (m *sliceMap[K, D]) Get(key K) (D, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	value, exist := m.body[key]
	return value, exist
}

func (m *sliceMap[K, D]) Each(f func(key K, value D)) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for key, value := range m.body {
		f(key, value)
	}

}
