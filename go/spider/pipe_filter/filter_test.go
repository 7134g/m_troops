package pipe_filter

import "testing"

func TestNewPipeLine(t *testing.T) {
	sf := NewSumFilter()
	slf := NewSplitFilter("x")
	tif := NewToIntFilter()
	_ = NewPipeLine("shm", sf, slf, tif)
}
