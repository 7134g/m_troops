package util

import (
	"fmt"
	"testing"
)

func TestDeepCopyMap(t *testing.T) {
	a := map[string]interface{}{
		"aaa":  "3",
		"aaaa": "4",
	}
	a1 := DeepCopy(a).(map[string]interface{})
	b := map[string]string{
		"bb":   "3",
		"bbbb": "4",
	}
	b1 := DeepCopy(b).(map[string]string)
	a1["aaa"] = "333333333333"
	fmt.Printf("%p, %p,\n %p, %p\n", a1, a, b1, b)

}
