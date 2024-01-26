package variant

import (
	"fmt"
	"strconv"
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

func TestBinding_Stowage(t *testing.T) {
	type Dst struct {
		Name string
	}

	type Src struct {
		Name string
	}

	dst := &Dst{}
	src := &Src{Name: "test"}

	StructToStruct(dst, src)
	fmt.Println(dst)
}

func TestName(t *testing.T) {
	num := 2.342000000
	str := strconv.FormatFloat(num, 'f', -1, 64)
	fmt.Println(str)
	fmt.Println(fmt.Sprintf("%f", num))
}
