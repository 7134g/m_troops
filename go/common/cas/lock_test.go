package cas

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewSpinLock(t *testing.T) {

	t.Run("lock", func(t *testing.T) {
		runSpeed(&sync.Mutex{})
	})

	t.Run("cas", func(t *testing.T) {
		runSpeed(NewSpinLock())
	})

}

func runSpeed(lock sync.Locker) {
	data := map[int]struct{}{}
	wg := sync.WaitGroup{}
	st := time.Now()
	for i := 0; i < 1000000; i++ {
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
			_, exist := data[i]
			if exist {

			}
			lock.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(st, "=====>", time.Since(st))
}
