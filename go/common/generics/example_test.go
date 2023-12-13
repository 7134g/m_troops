package generics

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestGenID(t *testing.T) {
	t.Run("int64 0", func(t *testing.T) {
		id := GenID(int64(0))
		t.Log(id)
	})

	t.Run("primitive.NilObjectID", func(t *testing.T) {
		id := GenID(primitive.NilObjectID)
		t.Log(id)
	})

	t.Run("string", func(t *testing.T) {
		id := GenID("")
		t.Log(id)
	})
}
