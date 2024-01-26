package db

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceString(t *testing.T) {
	type data struct {
		Name   string       `json:"name" redis:"name"`
		Age    int          `json:"age" redis:"age"`
		Photos *SliceString `json:"photos" redis:"photos"`
	}

	ctx := context.Background()
	key := "test_hash"
	client.FlushDB(ctx)
	err := client.HSet(ctx, key,
		"name", "liming",
		"age", 21,
		"photos", NewSliceString([]string{"1.jpg", "2.png"}),
	).Err()
	if err != nil {
		t.Fatal(err)
	}

	d := data{}
	if err = client.HGetAll(ctx, key).Scan(&d); err != nil {
		t.Fatal(err)
	}

	b, _ := json.Marshal(d)
	t.Log(string(b))
	t.Log(d.Photos.Get())
	assert.Equal(t, data{
		Name:   "liming",
		Age:    21,
		Photos: NewSliceString([]string{"1.jpg", "2.png"}),
	}, d)
}
