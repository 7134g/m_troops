package general

import (
	"errors"
	"reflect"
)

//package main
//
//import (
//	"errors"
//	"reflect"
//)
//
//type Student struct {
//	Name string
//	Age int
//	Scord int
//}
//
//type Teacher struct {
//	Name string
//	ClassName string
//}
//
//func main() {
//	setting := map[string]interface{}{
//		"Name": "Tom",
//		"Age": 17,
//		"ClassName": 1,
//	}
//
//	stud := Student{}
//	fmt.Println("do task ago, stud = ", stud)
//	err := filBysetting(&stud, setting)
//	if err != nil{
//		fmt.Println(err)
//	}
//	fmt.Println("do task after, stud = ", stud)
//
//	teac := Teacher{}
//	fmt.Println("do task ago, stud = ", teac)
//	err = filBysetting(&teac, setting)
//	if err != nil{
//		fmt.Println(err)
//	}
//	fmt.Println("do task after, stud = ", teac)
//
//}

func filBysetting(st interface{}, setting map[string]interface{}) error {
	isPtr := reflect.TypeOf(st).Kind() == reflect.Ptr
	isStruct := reflect.TypeOf(st).Elem().Kind() == reflect.Struct
	if !isPtr || !isStruct {
		return errors.New("st 必须是指针变量")
	}

	if setting == nil {
		return errors.New("setting is not nil")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range setting {
		if field, ok = reflect.TypeOf(st).Elem().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil
}
