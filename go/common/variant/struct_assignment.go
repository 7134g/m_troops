package variant

import (
	"errors"
	"reflect"
)

func MapToStruct(st interface{}, setting map[string]interface{}) error {
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

func StructToMap(obj interface{}) map[string]interface{} {
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
