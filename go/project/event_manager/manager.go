package event_manager

/*
事件管理器是一个用于管理事件和监听器的结构。它具有两个方法：Add和Run。Add方法用于将新的监听器附加到事件上，而Run方法用于执行事件管理器。
*/

type Listener[T any] func(T)

type Manager[T any] interface {
	Add(n string, l Listener[T])
	Run()
}

/*

BaseManager提供了Add方法用于添加监听器和Invoke方法用于触发指定事件的监听器。

*/

type BaseManager[T any] struct {
	lst map[string][]Listener[T]
}

func (m *BaseManager[T]) Invoke(n string, args T) {
	for _, ls := range m.lst[n] {
		go ls(args)
	}
}

func (m *BaseManager[T]) Add(n string, l Listener[T]) {
	m.lst[n] = append(m.lst[n], l)
}
