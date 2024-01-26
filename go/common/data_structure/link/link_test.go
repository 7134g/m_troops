package link

import (
	"encoding/json"
	"testing"
)

func TestListNode(t *testing.T) {
	head := NewListNodeMux()
	head.Put(1)
	head.Put(2)
	head.Put(3)
	b, err := json.Marshal(head.List())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))

	t.Log(head.Pop())
	b, err = json.Marshal(head.List())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))

	head.Reverse()
	b, err = json.Marshal(head.List())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}

func TestListNodeChannel(t *testing.T) {
	head := NewListNodeChannel()
	head.Put(1)
	head.Put(2)
	head.Put(3)
	b, err := json.Marshal(head.List())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))

	t.Log(head.Pop())
	b, err = json.Marshal(head.List())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
	head.Close()

}
