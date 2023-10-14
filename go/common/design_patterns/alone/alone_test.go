package alone

import (
	"fmt"
	"testing"
)

func TestAlone(t *testing.T) {
	obj1 := GetObjectAlone("first")
	fmt.Println(obj1.name)
	obj2 := GetObjectAlone("second")
	fmt.Println(obj2.name)
}
