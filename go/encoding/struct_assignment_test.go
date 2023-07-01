package encoding

import (
	"fmt"
	"testing"
)

func TestMapToStruct(t *testing.T) {
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
	err := MapToStruct(&stud, setting)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("do task after, stud = ", stud)

	teac := struct {
		Name      string
		ClassName string
	}{}
	fmt.Println("do task ago, stud = ", teac)
	err = MapToStruct(&teac, setting)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("do task after, stud = ", teac)

}

func TestStructToMap(t *testing.T) {
	s := struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}{"John Doe", 30, "john@example.com"}

	t.Log(StructToMap(s))
}
