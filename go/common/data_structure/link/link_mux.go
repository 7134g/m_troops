package link

import (
	"sync"
)

// ListNodeMux 链表
type ListNodeMux struct {
	mux sync.Mutex
	*base
}

func NewListNodeMux() *ListNodeMux {
	return &ListNodeMux{
		mux:  sync.Mutex{},
		base: &base{},
	}
}

func (l *ListNodeMux) start() {
	l.mux.Lock()
}
func (l *ListNodeMux) end() {
	l.mux.Unlock()
}

func (l *ListNodeMux) Put(val any) {
	if val == nil {
		return
	}
	l.start()
	defer l.end()

	if l.nodes == nil {
		l.nodes = &node{Val: val}
	} else {
		n := &node{Val: val}
		l.nodes, n.Next = n, l.nodes
	}
}

func (l *ListNodeMux) Pop() any {
	l.start()
	defer l.end()

	if l.nodes == nil {
		return nil
	} else {
		val := l.nodes.Val
		l.nodes = l.nodes.Next
		return val
	}

}

// Reverse 反转链表的实现
func (l *ListNodeMux) Reverse() {
	if l.nodes == nil {
		return
	}
	l.start()
	defer l.end()
	cur := l.nodes
	var pre *node = nil
	var t *node = nil
	for cur != nil {
		t = pre
		pre = cur
		cur = cur.Next
		pre.Next = t
		pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要
	}
	l.nodes = pre
}

func (l *ListNodeMux) List() []any {
	l.start()
	defer l.end()
	result := make([]any, 0)
	cur := l.nodes
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}

	return result
}
