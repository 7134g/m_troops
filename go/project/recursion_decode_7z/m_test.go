package main

import (
	"sync/atomic"
	"testing"
)

func TestName(t *testing.T) {
	var count atomic.Int64

	count.Add(1)

}
