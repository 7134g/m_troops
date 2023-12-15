package main

import (
	"log"
	"sync/atomic"
	"time"
)

var cryptCount int64 = 0

func monitor() {
	var lastCount int64
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			nowCount := atomic.LoadInt64(&cryptCount)
			speed := nowCount - lastCount
			lastCount = nowCount
			log.Printf("每秒尝试密码 %d 次\n", speed)
		}
	}
}

func generate(prefix string, characters string, remainingLength int) {
	if remainingLength == 0 {
		//*result = append(*result, prefix)
		if prefix == "" {
			return
		}
		passwordChan <- prefix
		return
	}

	for _, char := range characters {
		newPrefix := prefix + string(char)
		generate(newPrefix, characters, remainingLength-1)
	}
}

func generateAllPossibleStrings() {
	for i := minLen; i <= maxLen; i++ {
		generate("", characters, i)
	}

	close(passwordChan)
}
