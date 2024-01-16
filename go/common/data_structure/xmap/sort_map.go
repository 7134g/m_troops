package xmap

import "fmt"

type SortMap struct {
	Map
	keys []interface{}
}

func NewSortMap() *SortMap {
	m := &SortMap{
		Map: Map{
			body: make(map[interface{}]interface{}),
			opt:  make(chan optData, 0),
			rst:  make(chan rstData, 1),
		},
		keys: make([]interface{}, 0),
	}

	go m.run()
	return m
}

func (m *SortMap) Put(key interface{}, value interface{}) {
	m.keys = append(m.keys, key)
	m.Map.Put(key, value)
}

func (m *SortMap) Keys() []interface{} {
	return m.keys
}

func (m *SortMap) Clear() {
	m.keys = make([]interface{}, 0)
	m.Map.Clear()
}

func (m *SortMap) String() string {
	str := "SortMap\n"
	str += fmt.Sprintf("%v", m.body)
	return str
}
