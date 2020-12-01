package common

import (
	"encoding/json"
	"fmt"
	"io"
	"m_troops/go/spider/Mcooly/common/logs"
	"m_troops/go/spider/Mcooly/setting"
	"os"
	"sync"
)

type TaskFile struct {
	Name string
	Data []interface{}
	path string
	mux  sync.RWMutex
}

func NewTaskFile(name, path string) *TaskFile {
	err := setting.MakeDir(path)
	if err != nil {
		logs.Exit(err)
	}
	return &TaskFile{
		Name: name,
		Data: make([]interface{}, 0),
		path: path,
	}
}

func (t *TaskFile) Load() {
	t.mux.Lock()
	defer t.mux.Unlock()
	filePtr, err := os.OpenFile(t.path, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("Open file failed [Err:%s]\n", err.Error())
		return
	}
	defer filePtr.Close()
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&t.Data)
	if err != nil && err != io.EOF {
		logs.Errors(err)
	}
}

func (t *TaskFile) Append(d interface{}) {
	t.mux.Lock()
	defer t.mux.Unlock()
	t.Data = append(t.Data, d)
}

func (t *TaskFile) Dump() {
	t.mux.Lock()
	defer t.mux.Unlock()
	filePtr, err := os.OpenFile(t.path, os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("dump file failed", err.Error())
		return
	}
	defer filePtr.Close()

	if len(t.Data) == 0 {
		return
	}
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(t.Data)
	if err != nil {
		logs.Errors(err)
	}
}

func (t *TaskFile) Pop() interface{} {
	t.mux.Lock()
	defer t.mux.Unlock()
	if len(t.Data) == 0 {
		return nil
	}
	p := t.Data[len(t.Data)-1]
	t.Data = t.Data[:len(t.Data)-1]
	return p
}

func (t *TaskFile) Pull() interface{} {
	t.mux.Lock()
	defer t.mux.Unlock()
	if len(t.Data) == 0 {
		return nil
	}
	p := t.Data[0]
	t.Data = t.Data[1:]
	return p
}

func (t *TaskFile) Bind(o, n interface{}) {
	t.mux.Lock()
	defer t.mux.Unlock()
	b, _ := json.Marshal(o)
	_ = json.Unmarshal(b, n)
}
