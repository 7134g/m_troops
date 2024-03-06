package lock

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewSpinLock(t *testing.T) {

	t.Run("lock", func(t *testing.T) {
		data := map[int]struct{}{}
		wg := sync.WaitGroup{}
		st := time.Now()
		lock := sync.Mutex{}
		for i := 0; i < 1000; i++ {
			wg.Add(2)
			go func(i int) {
				defer wg.Done()
				lock.Lock()
				data[i] = struct{}{}
				lock.Unlock()
			}(i)

			go func(i int) {
				defer wg.Done()
				lock.Lock()
				x := data[i]
				fmt.Printf("%d%v", i, x)
				lock.Unlock()
			}(i)
		}
		wg.Wait()
		fmt.Println(st, "=====>", time.Since(st))
	})

	t.Run("cas", func(t *testing.T) {
		data := map[int]struct{}{}
		wg := sync.WaitGroup{}
		lock := NewSpinLock()
		st := time.Now()
		fmt.Println(st)
		for i := 0; i < 1000; i++ {
			wg.Add(2)
			go func(i int) {
				defer wg.Done()
				lock.Lock()
				data[i] = struct{}{}
				lock.Unlock()
			}(i)

			go func(i int) {
				defer wg.Done()
				lock.Lock()
				x := data[i]
				fmt.Printf("%d%v", i, x)
				lock.Unlock()
			}(i)
		}
		wg.Wait()
		fmt.Println(st, "=====>", time.Since(st))
	})

}
