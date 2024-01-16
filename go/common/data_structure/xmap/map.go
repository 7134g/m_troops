package xmap

import "fmt"

type operate uint8

const (
	_read   operate = 1
	_write  operate = 2
	_del    operate = 3
	_keys   operate = 4
	_values operate = 5
	_clean  operate = 6
)

type optData struct {
	key   interface{}
	value interface{}
	flag  operate
}

type rstData struct {
	value interface{}
	found bool
}

type Map struct {
	body map[interface{}]interface{}
	opt  chan optData
	rst  chan rstData
}

// New instantiates a hash map.
func New() *Map {
	m := &Map{
		body: make(map[interface{}]interface{}),
		opt:  make(chan optData),
		rst:  make(chan rstData),
	}
	go m.run()
	return m
}

func (m *Map) run() {
	for {
		select {
		case d := <-m.opt:
			switch d.flag {
			case _read:
				value, found := m.body[d.key]
				m.rst <- rstData{
					value: value,
					found: found,
				}
			case _write:
				m.body[d.key] = d.value
			case _del:
				delete(m.body, d.key)
			case _keys:
				_ks := make([]interface{}, m.Size())
				count := 0
				for key := range m.body {
					_ks[count] = key
					count++
				}
				m.rst <- rstData{
					value: _ks,
				}
			case _values:
				_vs := make([]interface{}, m.Size())
				count := 0
				for _, value := range m.body {
					_vs[count] = value
					count++
				}
				m.rst <- rstData{
					value: _vs,
				}
			case _clean:
				m.body = make(map[interface{}]interface{})
			}
		}
	}
}

// Put inserts element into the map.
func (m *Map) Put(key interface{}, value interface{}) {
	m.opt <- optData{
		key:   key,
		value: value,
		flag:  _write,
	}
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	m.opt <- optData{
		key:  key,
		flag: _read,
	}
	data := <-m.rst
	return data.value, data.found
}

// Remove removes the element from the map by key.
func (m *Map) Remove(key interface{}) {
	m.opt <- optData{
		key:  key,
		flag: _del,
	}
}

// Empty returns true if map does not contain any elements
func (m *Map) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map) Size() int {
	return len(m.body)
}

// Keys returns all keys (random order).
func (m *Map) Keys() []interface{} {
	m.opt <- optData{
		flag: _keys,
	}
	data := <-m.rst
	return data.value.([]interface{})
}

// Values returns all values (random order).
func (m *Map) Values() []interface{} {
	m.opt <- optData{
		flag: _values,
	}
	data := <-m.rst
	return data.value.([]interface{})
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.opt <- optData{
		flag: _clean,
	}
}

// String returns a string representation of container
func (m *Map) String() string {
	str := "SafeMap\n"
	str += fmt.Sprintf("%v", m.body)
	return str
}
