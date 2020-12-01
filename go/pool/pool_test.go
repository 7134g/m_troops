package pool

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestNewPool(t *testing.T) {
	pool, cancel := NewPool(500, true, time.Second)
	defer pool.Close()
	defer func() {
		time.Sleep(5 * time.Second)
		for i := 0; i < 500; i++ {
			var task = &Task{
				TaskFunc: func(i []interface{}) {
					time.Sleep(10 * time.Millisecond)
					for _, m := range i {
						if x, ok := m.(int); ok {
							fmt.Printf("second task ->%d goroutine: %d\n", x, runtime.NumGoroutine())
						}
					}
				},
				Param: []interface{}{i},
			}
			err := pool.Submit(task)
			if err != nil {
				t.Fatal(err)
			}
			if i > 100 && i < 250 {
				pool.ReSize(10)
			}
			if i > 250 && i < 400 {
				pool.ReSize(30)
			}
			if i > 400 {
				pool.ReSize(10)
			}
			//t.Logf("---running goroutine pool: %d, system: %d", pool.Running(), runtime.NumGoroutine())
		}
		pool.Wait()
		time.Sleep(10 * time.Second)
		// running goroutine pool: 1, system: 4
		t.Logf("running goroutine pool: %d, system: %d", pool.Running(), runtime.NumGoroutine())
	}()
	for i := 1; i < 1000; i++ {
		var task = &Task{
			TaskFunc: func(i []interface{}) {
				time.Sleep(10 * time.Millisecond)
				for _, m := range i {
					if x, ok := m.(int); ok {
						fmt.Printf("fist task ->%d goroutine: %d\n", x, runtime.NumGoroutine())
					}
				}
			},
			Param: []interface{}{i},
		}
		err := pool.Submit(task)
		if err != nil {
			t.Fatal(err)
		}
		if i > 200 && i < 500 {
			pool.ReSize(3)
		}
		if i > 500 && i < 800 {
			pool.ReSize(5)
		}
		cancel()
		if i > 800 && i < 1000 {
			pool.ReSize(2)
		}
		//t.Logf("running goroutine pool: %d, system: %d", pool.Running(), runtime.NumGoroutine())
	}
}
