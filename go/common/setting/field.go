package setting

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
)

type emptyType struct{}

const emptyTypeValue = "_emptyType"

// 根据目标字段类型进行值转换
func convertToFieldType(value string, dstValueType reflect.StructField) (reflect.Value, error) {
	if value == emptyTypeValue {
		return reflect.Zero(reflect.TypeOf(&emptyType{})), nil
	}

	dstValueTypeKind := dstValueType.Type.Kind()
	switch dstValueTypeKind {
	case reflect.Int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return reflect.ValueOf(nil), fmt.Errorf("failed to convert value to int")
		}

		return reflect.ValueOf(intValue), nil
	case reflect.Int32:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return reflect.ValueOf(nil), fmt.Errorf("failed to convert value to int")
		}

		return reflect.ValueOf(int32(intValue)), nil
	case reflect.Int64:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return reflect.ValueOf(nil), fmt.Errorf("failed to convert value to int")
		}

		return reflect.ValueOf(int64(intValue)), nil
	case reflect.Uint64:
		uintValue, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("failed to convert value to uint")
		}

		return reflect.ValueOf(uintValue), nil
	case reflect.Uint32:
		uintValue, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("failed to convert value to uint")
		}

		return reflect.ValueOf(uint32(uintValue)), nil
	case reflect.Uint:
		uintValue, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("failed to convert value to uint")
		}

		return reflect.ValueOf(uint(uintValue)), nil
	case reflect.Float64:
		float64Value, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("failed to convert value to uint")
		}

		return reflect.ValueOf(float64Value), nil
	case reflect.String:
		return reflect.ValueOf(value), nil
	default:
		// 对于其他数据类型，可以根据需要进行相应的处理和转换
		return reflect.Zero(dstValueType.Type), nil
	}
}

type fieldInterface interface {
	GetField() string
	GetValue() interface{}
}

type record struct {
	Field        string      `json:"field,omitempty"`       // redis 中的field
	Value        interface{} `json:"value,omitempty"`       // 实时值
	DefaultValue interface{} `json:"default_value"`         // 默认值
	Description  string      `json:"description,omitempty"` // 描述
	Category     string      `json:"category,omitempty"`
}

func (r *record) GetField() string {
	return r.Field
}

// GetValue 获取值
func (r *record) GetValue() interface{} {
	if r.Value == nil {
		return r.DefaultValue
	}
	return r.Value
}

// SetValue 更新实时值
func (r *record) SetValue(ctx context.Context, value interface{}) {
	r.Value = value
}

type key interface {
	~int | ~int32 | ~int64 |
		~uint | ~uint32 | ~uint64 |
		~float64 |
		~string
}

type field[K key] struct {
	record
	DefaultValue K `json:"default_value"` // 默认值
}

func (f *field[K]) Get() K {
	value, ok := f.Value.(K)
	if !ok {
		return f.DefaultValue
	}

	return value
}
