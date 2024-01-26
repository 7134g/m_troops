package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type idType interface {
	int64 | primitive.ObjectID | string
}

func checkID[D idType](id D) D {
	var value any = &id
	switch v := value.(type) {
	case *int64:
		if 0 == *v {
			//*v = snow_flake.GetSnowflakeID()
			// todo 雪花id
		}

		return id
	case *primitive.ObjectID:
		if primitive.NilObjectID == *v {
			*v = primitive.NewObjectID()
		}
		return id
	default:
		return id
	}
}
