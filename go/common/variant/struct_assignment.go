package variant

import (
	"errors"
	"fmt"
	"reflect"
)

func MapToStruct(st interface{}, data map[string]interface{}) error {
	isPtr := reflect.TypeOf(st).Kind() == reflect.Ptr
	isStruct := reflect.TypeOf(st).Elem().Kind() == reflect.Struct
	if !isPtr || !isStruct {
		return errors.New("st 必须是指针变量")
	}

	if data == nil {
		return errors.New("data is not nil")
	}

	vstr := reflect.ValueOf(st)
	if vstr.IsNil() {
		return errors.New("st 不能是空指针")
	}
	vstr = vstr.Elem()

	for k, v := range data {
		field, ok := vstr.Type().FieldByName(k)
		if !ok {
			continue
		}

		tv := reflect.ValueOf(v)
		if !tv.Type().ConvertibleTo(field.Type) {
			return fmt.Errorf("无法将 %s 类型的值转换为 %s 类型", tv.Type(), field.Type)
		}

		tv = tv.Convert(field.Type)
		if !vstr.FieldByName(k).CanSet() {
			return fmt.Errorf("字段 %s 无法设置", k)
		}

		vstr.FieldByName(k).Set(tv)
	}

	return nil
}

func StructToMapWithJson(obj interface{}) map[string]interface{} {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	objType := objValue.Type()

	data := make(map[string]interface{})
	for i := 0; i < objValue.NumField(); i++ {
		field := objType.Field(i)
		value := objValue.Field(i).Interface()

		if len(field.Tag.Get("json")) == 0 {
			data[field.Name] = value
		} else {
			data[field.Tag.Get("json")] = value
		}

	}

	return data
}

func StructToStruct(dst, src interface{}) {
	dstValue := reflect.ValueOf(dst).Elem()
	srcValue := reflect.ValueOf(src).Elem()

	for i := 0; i < srcValue.NumField(); i++ {
		srcFieldValue := srcValue.Field(i)
		srcFieldType := srcValue.Type().Field(i)

		dstFieldValue := dstValue.FieldByName(srcFieldType.Name)

		if dstFieldValue.IsValid() && dstFieldValue.CanSet() && dstFieldValue.Type() == srcFieldType.Type {
			dstFieldValue.Set(srcFieldValue)
		}
	}
}
