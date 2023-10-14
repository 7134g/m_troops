package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var count int64 = 0

func monitor() {
	var lastCount int64
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			nowCount := atomic.LoadInt64(&count)
			speed := nowCount - lastCount
			lastCount = nowCount
			fmt.Printf("每秒 %d 次\n", speed)
		}
	}
}
