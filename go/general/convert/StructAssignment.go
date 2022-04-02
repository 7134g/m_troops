package convert

import (
	"errors"
	"reflect"
)

func FilBySetting(st interface{}, setting map[string]interface{}) error {
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
