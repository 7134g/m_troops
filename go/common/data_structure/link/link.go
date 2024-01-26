package link

type IListNode interface {
	start()
	end()
	Put(val any)
	Pop() any
	Close()
}

type base struct {
	nodes *node
}

func (l *base) start() {}

func (l *base) end() {}

func (l *base) Close() {}

// 链表节点
type node struct {
	Val  any
	Next *node
}
