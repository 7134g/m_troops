package safe_map

import (
	"sync"
	"testing"
	"time"
)

func TestMap_Put(t *testing.T) {
	g := sync.WaitGroup{}
	m := New()
	m.Put(21, "b")
	go func() {
		for i := 0; i < 100000; i++ {
			g.Add(1)
			m.Put(i, "a")
			g.Done()
		}
	}()

	go func() {
		m.Remove(21)
		for i := 20; i < 2000; i++ {
			g.Add(1)
			m.Get(i)
			g.Done()
		}
	}()

	time.Sleep(time.Second)
	g.Wait()

	t.Log(m.Keys())
	t.Log(m.Values())
	m.Clear()
	t.Log(m.Keys())
	t.Log(m.Values())
	t.Log(m.String())
}

func TestSortMap_Put(t *testing.T) {
	m := NewSortMap()
	for i := 0; i < 10; i++ {
		m.Put(i, "a")
	}

	t.Log(m.Keys())
	t.Log(m.Values())
	m.Clear()
	t.Log(m.Keys())
	t.Log(m.Values())
	t.Log(m.String())
}
