package system

import "testing"

func TestSetProxy(t *testing.T) {
	SetProxy("127.0.0.1:7890")
}
