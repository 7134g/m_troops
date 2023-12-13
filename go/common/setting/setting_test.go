package setting

import (
	"context"
	"encoding/json"
	"fmt"
	rdsV8 "github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var rds *rdsV8.Client

func init() {
	fmt.Println("=====================> test")
	rds = rdsV8.NewClient(&rdsV8.Options{
		Addr:         "127.0.0.1:6379",
		DB:           0,
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		PoolTimeout:  2 * time.Minute,
		IdleTimeout:  10 * time.Minute,
		PoolSize:     1000,
	})
	rds.HSet(context.Background(), "setting:test", "test_int_64", int64(1))
	rds.Publish(context.Background(), "setting:pub_sub:test", "update")
}

func TestInitialize(t *testing.T) {
	if err := Initialize(rds, "test"); err != nil {
		t.Fatal(err)
	}
	t.Log(Global)

	t.Log(Global.LoadRedisData(context.Background()))

	fmt.Printf("\n\n\n\n============\n\n\n\n")

	t.Log(Global)

	t.Log(Global.GenericsInt64.Get())
}

func TestServiceConfig_GetAllSetting(t *testing.T) {
	ori := ServiceSetting{}
	ori.GenericsInt64 = field[int64]{}
	ori.GenericsInt64.Value = int64(1000)
	ori.GenericsInt64.Description = "5 level"
	ori.GenericsInt64.DefaultValue = 2000
	m, err := ori.GetAllSetting()
	if err != nil {
		t.Fatal(err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	str := string(b)
	t.Log(str)
	t.Log(ori.GenericsInt64.Get())
}

func TestServiceSetting_LoadRedisData(t *testing.T) {
	ctx := context.Background()

	env := "test"
	ori := ServiceSetting{
		rds:    rds,
		rdsKey: fmt.Sprintf("setting:%s", env),
		subKey: fmt.Sprintf("setting:pub_sub:%s", env),
	}
	if err := ori.LoadRedisData(ctx); err != nil {
		t.Fatal(err)
	}

	t.Log(ori)
}

func TestGenericsField(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		f := field[int]{}
		f.DefaultValue = 10
		t.Log(f.Get())

		f.Value = 12
		t.Log(f.Get())
	})

	t.Run("int32", func(t *testing.T) {
		f := field[int32]{}
		f.DefaultValue = 10
		t.Log(f.Get())

		f.Value = int32(12)
		t.Log(f.Get())
	})

	t.Run("int64", func(t *testing.T) {
		f := field[int64]{}
		f.DefaultValue = 10
		t.Log(f.Get())

		f.Value = int64(12)
		t.Log(f.Get())
	})

	t.Run("uint", func(t *testing.T) {
		f := field[uint]{}
		f.DefaultValue = 10
		t.Log(f.Get())

		f.Value = uint(12)
		t.Log(f.Get())
	})

	t.Run("uint32", func(t *testing.T) {
		f := field[uint32]{}
		f.DefaultValue = 10
		t.Log(f.Get())

		f.Value = uint32(12)
		t.Log(f.Get())
	})

	t.Run("uint64", func(t *testing.T) {
		f := field[uint64]{}
		f.DefaultValue = 10
		t.Log(f.Get())

		f.Value = uint64(12)
		t.Log(f.Get())
	})

	t.Run("float64", func(t *testing.T) {
		f := field[float64]{}
		f.DefaultValue = 10
		t.Log(f.Get())

		f.Value = float64(12)
		t.Log(f.Get())
	})

	t.Run("string", func(t *testing.T) {
		f := field[string]{}
		f.DefaultValue = "zzzz"
		t.Log(f.Get())

		f.Value = "xxxx"
		t.Log(f.Get())
	})
}
