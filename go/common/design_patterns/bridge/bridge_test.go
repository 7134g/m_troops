package bridge

import "testing"

func TestBridge(t *testing.T) {
	var bridge Bridge
	ch := &Chinese{}
	en := &English{}

	bridge = ch
	bridge.Run()
	bridge = en
	bridge.Run()
}
