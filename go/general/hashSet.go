package general

var itemExists = struct{}{}

type Set struct {
	items map[interface{}]struct{}
}

func New() *Set {
	return &Set{items: make(map[interface{}]struct{})}
}

func (set *Set) Add(item interface{}) {
	set.items[item] = itemExists
}

func (set *Set) Remove(item interface{}) {
	delete(set.items, item)
}

func (set *Set) Contains(item interface{}) bool {
	if _, contains := set.items[item]; !contains {
		return false
	}
	return true
}
