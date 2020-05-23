package main

import "bytes"

type Hset struct {
    m map[interface{}]bool
}

// 给set结构初始化
func NewHset() *Hset {
    return &Hset{m: make(map[interface{}]bool)}
}

func (h *Hset) Add(v interface{}) (b bool) {
    if h.m[v]{
        return false
    }
    h.m[v] = true
    return true
}

func (h *Hset) Delete(v interface{}) {
    delete(h.m, v)
}

func (h *Hset) Clear() {
    h.m = make(map[interface{}]bool)
}

func (h *Hset) Contains(v interface{}) bool {
    return h.m[v]
}

func (h *Hset) Len() int {
    return len(h.m)
}

func (h *Hset) Same(other *Hset) bool {
    if other == nil{
        return false
    }

    if other.Len() != h.Len(){
        return false
    }

    for k := range h.m{
        if h.m[k] != other.m[k]{
            return false
        }
    }
    return true
}

func (h *Hset) Elements() []interface{} {
    length := len(h.m)

    result := make([]interface{}, length)

    index := 0

    for k := range h.m{
        result[index] = k
        index ++
    }

    return result
}

func (h *Hset) String() string {
    var buf bytes.Buffer

    buf.WriteString("set{")

    length := h.Len()


    for i, v := range h.Elements() {
        buf.WriteString(fmt.Sprintf("%v", v))
        if i < length - 1{
            buf.WriteString(" ")
        }

    }

    buf.WriteString("}")

    return buf.String()
}
