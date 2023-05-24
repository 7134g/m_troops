package encoding

import (
	"fmt"
	"testing"
)

func TestFilBySetting(t *testing.T) {
	setting := map[string]interface{}{
		"Name":      "Tom",
		"Age":       17,
		"ClassName": 1,
	}

	stud := struct {
		Name  string
		Age   int
		Scord int
	}{}
	fmt.Println("do task ago, stud = ", stud)
	err := FilBySetting(&stud, setting)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("do task after, stud = ", stud)

	teac := struct {
		Name      string
		ClassName string
	}{}
	fmt.Println("do task ago, stud = ", teac)
	err = FilBySetting(&teac, setting)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("do task after, stud = ", teac)

}
