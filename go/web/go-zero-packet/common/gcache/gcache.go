package gcache

import (
	"context"
	"demo/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
	"time"
)

type GCache struct {
	skipCache bool
	Expiry    time.Duration

	Cache cache.Cache
	Data  any
}

/* cache 实现 */

func (u *GCache) Del(keys ...string) error {
	if u.skipCache {
		return nil
	}
	return u.Cache.Del(keys...)
}

func (u *GCache) DelCtx(ctx context.Context, keys ...string) error {
	if u.skipCache {
		return nil
	}
	return u.Cache.DelCtx(ctx, keys...)
}

func (u *GCache) Get(key string, val any) error {
	u.begin(val, key)
	if u.skipCache {
		return nil
	}
	return u.Cache.Get(key, val)
}

func (u *GCache) GetCtx(ctx context.Context, key string, val any) error {
	u.begin(val, key)
	if u.skipCache {
		return nil
	}
	return u.Cache.GetCtx(ctx, key, val)
}

func (u *GCache) IsNotFound(err error) bool {
	return u.Cache.IsNotFound(err)
}

func (u *GCache) Set(key string, val any) error {
	u.begin(val, key)
	if u.skipCache {
		return nil
	}
	defer func() { u.end(val, key) }()
	return u.Cache.Set(key, val)
}

func (u *GCache) SetCtx(ctx context.Context, key string, val any) error {
	u.begin(val, key)
	if u.skipCache {
		return nil
	}
	defer func() { u.end(val, key) }()
	return u.Cache.SetCtx(ctx, key, val)
}

func (u *GCache) SetWithExpire(key string, val any, expire time.Duration) error {
	u.begin(val, key)
	if u.skipCache {
		return nil
	}
	defer func() { u.end(val, key) }()
	return u.Cache.SetWithExpire(key, &u.Data, u.Expiry)
}

func (u *GCache) SetWithExpireCtx(ctx context.Context, key string, val any, expire time.Duration) error {
	u.begin(val, key)
	if u.skipCache {
		return nil
	}
	defer func() { u.end(val, key) }()
	return u.Cache.SetWithExpireCtx(ctx, key, &u.Data, u.Expiry)
}

func (u *GCache) Take(val any, key string, query func(val any) error) error {
	u.begin(val, key)
	if u.skipCache {
		return query(val)
	}
	defer func() { u.end(val, key) }()
	return u.Cache.Take(&u.Data, key, query)
}

func (u *GCache) TakeCtx(ctx context.Context, val any, key string, query func(val any) error) error {
	u.begin(val, key)
	if u.skipCache {
		return query(val)
	}
	defer func() { u.end(val, key) }()
	return u.Cache.TakeCtx(ctx, &u.Data, key, query)
}

func (u *GCache) TakeWithExpire(val any, key string, query func(val any, expire time.Duration) error) error {
	u.begin(val, key)
	if u.skipCache {
		return query(val, u.Expiry)
	}
	defer func() { u.end(val, key) }()
	return u.Cache.TakeWithExpire(&u.Data, key, query)
}

func (u *GCache) TakeWithExpireCtx(ctx context.Context, val any, key string, query func(val any, expire time.Duration) error) error {
	u.begin(val, key)
	if u.skipCache {
		return query(val, u.Expiry)
	}
	defer func() { u.end(val, key) }()
	return u.Cache.TakeWithExpireCtx(ctx, &u.Data, key, query)
}

/* 增加自定义逻辑 */
const (
	defaultExpiry = time.Hour * 24 * 7
)

func (u *GCache) begin(val any, key string) {

}

func (u *GCache) end(val any, key string) {
	//u.copy(val)
}

func (u *GCache) copy(val any) {
	if u.Data == nil {
		return
	}
	_ = copier.Copy(val, u.Data)
}

func (u *GCache) OpenSkipCache() {
	u.skipCache = true
}

func (u *GCache) CloseSkipCache() {
	u.skipCache = false
}

// NewGCacheNode redis node
func NewGCacheNode(data interface{}, rds *redis.Redis, opts ...cache.Option) *GCache {
	node := cache.NewNode(rds, syncx.NewSingleFlight(), cache.NewStat("node"), xerr.ErrNotFound, opts...)
	return &GCache{
		Data:   data,
		Cache:  node,
		Expiry: defaultExpiry,
	}
}

// NewGCacheCluster redis Cluster
func NewGCacheCluster(data any, conf cache.ClusterConf, opts ...cache.Option) *GCache {
	cluster := cache.New(conf, syncx.NewSingleFlight(), cache.NewStat("cluster"), xerr.ErrNotFound, opts...)
	return &GCache{
		Data:   data,
		Cache:  cluster,
		Expiry: defaultExpiry,
	}
}
