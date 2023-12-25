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
