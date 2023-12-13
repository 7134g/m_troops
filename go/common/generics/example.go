package generics

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type idType interface {
	int64 | primitive.ObjectID | string
}

// GenID 判断类型并赋值
func GenID[D idType](id D) D {
	var value any = &id
	switch v := value.(type) {
	case *int64:
		if 0 == *v {
			*v = time.Now().UnixNano()
		}

		return id
	case *primitive.ObjectID:
		if primitive.NilObjectID == *v {
			*v = primitive.NewObjectID()
		}
		return id
	case *string:
		if "" == *v {
			*v = primitive.NewObjectID().Hex()
		}
		return id
	default:
		return id
	}
}

// 基础类型封装
type key interface {
	~int | ~int32 | ~int64 |
		~uint | ~uint32 | ~uint64 |
		~float64 |
		~string
}

type record struct {
	Value interface{} `json:"value,omitempty"` // 实时值
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
