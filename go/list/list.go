package list

import (
	"container/list"
	"sync"
	"time"
)

var TickerMsg = time.Millisecond * 200

// Message 消息结构体
type Message struct {
	Data
	Stamp int64 `json:"-"`
}

type Data struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// MyList 消息队列
type MyList struct {
	list  *list.List
	cache map[string]struct{}
	mux   sync.Mutex
}

func NewMyList() *MyList {
	return &MyList{list: list.New(), mux: sync.Mutex{}, cache: make(map[string]struct{})}
}

// PushBack 头插
func (l *MyList) PushBack(code int, msg string) {
	l.mux.Lock()
	if _, ok := l.cache[msg]; !ok {
		l.cache[msg] = struct{}{}
		l.list.PushBack(Message{Data: Data{Code: code, Msg: msg}, Stamp: time.Now().Unix()})
	}
	l.mux.Unlock()
}

// PushFront 尾插
func (l *MyList) PushFront(code int, msg string) {
	l.mux.Lock()
	l.list.PushFront(Message{Data: Data{Code: code, Msg: msg}, Stamp: time.Now().Unix()})
	l.mux.Unlock()
}

// PushFrontData 尾插
func (l *MyList) PushFrontData(data Message) {
	l.mux.Lock()
	l.list.PushFront(data)
	l.mux.Unlock()
}

// GetFront 从尾取
func (l *MyList) GetFront() interface{} {
	var inter interface{}
	l.mux.Lock()
	if l.Len() > 0 {
		inter = l.list.Remove(l.list.Front())
	}
	l.mux.Unlock()
	return inter
}

// GetBack 从头取
func (l *MyList) GetBack() interface{} {
	var inter interface{}
	l.mux.Lock()
	if l.Len() > 0 {
		inter = l.list.Remove(l.list.Back())
	}
	l.mux.Unlock()
	return inter
}

func (l *MyList) Remove(e *list.Element) {
	l.mux.Lock()
	l.list.Remove(e)
	l.mux.Unlock()
}

func (l *MyList) Len() int {
	length := l.list.Len()
	return length
}
