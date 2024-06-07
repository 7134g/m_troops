package encoding

import "testing"

func TestByteCountIEC(t *testing.T) {
	t.Log(ByteCountIEC(100))
	t.Log(ByteCountIEC(1024 * 3))
	t.Log(ByteCountIEC(1024 * 1024 * 3))
	t.Log(ByteCountIEC(1024 * 1024 * 1024 * 3))
}
