package link

import "context"

type operate uint8

const (
	_read   operate = 1
	_write  operate = 2
	_del    operate = 3
	_keys   operate = 4
	_values operate = 5
	_clean  operate = 6
)

type ListNodeChannel struct {
	w  chan any
	r  chan any
	op chan channelData

	status bool

	*base
}

type channelData struct {
	flag operate
	val  any
}

func NewListNodeChannel() *ListNodeChannel {
	l := &ListNodeChannel{
		w:      make(chan any),
		r:      make(chan any),
		op:     make(chan channelData),
		status: true,
		base:   &base{},
	}
	go l.run()
	return l
}

func (l *ListNodeChannel) run() {
	ctx, cancel := context.WithCancel(context.Background())
	for {
		select {
		case <-ctx.Done():
			cancel()
			l.status = false
			return
		case data := <-l.op:
			flag := data.flag
			var val any
			switch flag {
			case _write:
				if l.nodes == nil {
					l.nodes = &node{Val: data.val}
				} else {
					n := &node{Val: data.val}
					l.nodes, n.Next = n, l.nodes
				}
			case _read:
				if l.nodes == nil {
					val = nil
				} else {
					val = l.nodes.Val
					l.nodes = l.nodes.Next
				}
				l.r <- val
			case _values:
				result := make([]any, 0)
				cur := l.nodes
				for cur != nil {
					result = append(result, cur.Val)
					cur = cur.Next
				}
				l.r <- result
			case _clean:
				cancel()
			}

		}
	}
}

func (l *ListNodeChannel) Put(val any) {
	l.op <- channelData{flag: _write, val: val}
}

func (l *ListNodeChannel) Pop() any {
	l.op <- channelData{flag: _read}
	return <-l.r
}

func (l *ListNodeChannel) List() []any {
	l.op <- channelData{flag: _values}
	val := <-l.r
	return val.([]any)
}

func (l *ListNodeChannel) Close() {
	l.op <- channelData{flag: _clean}
}
