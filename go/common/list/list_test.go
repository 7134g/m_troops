package list

import (
	"fmt"
	"testing"
)

func TestNewMyList(t *testing.T) {
	m := NewMyList()
	m.PushFront(100, "100")
	m.PushFront(101, "100")
	m.PushFront(102, "100")
	x := m.GetBack()
	fmt.Println(x)
}
