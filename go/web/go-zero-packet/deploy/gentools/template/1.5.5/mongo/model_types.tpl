package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type {{.Type}} struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`

}
