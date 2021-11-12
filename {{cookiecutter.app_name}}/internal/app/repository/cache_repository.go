package repositories

import (
	"context"
	"encoding/json"
	"time"
)

// ICacheRepository interface for cache repo
type ICacheRepository interface {
	Write(ctx context.Context, key string, data interface{}, ttl time.Duration) (err error)
	WriteIfEmpty(ctx context.Context, key string, data interface{}, ttl time.Duration) (err error)
	Get(ctx context.Context, key string, data interface{}) (err error)
	Delete(ctx context.Context, key string) (err error)
}

type cacheRepo struct {
	opt Option
}

// NewCacheRepository initiate cache repo
func NewCacheRepository(opt Option) ICacheRepository {
	return &cacheRepo{
		opt: opt,
	}
}

// WriteCache this will and must write the data to cache with corresponding key using locking
func (c *cacheRepo) Write(ctx context.Context, key string, data interface{}, ttl time.Duration) (err error) {
	b, err := json.Marshal(data)
    if err != nil {
       return err
    }
	return c.opt.CacheClient.SetEX(ctx, key, b, ttl).Err()
}

// WriteCacheIfEmpty will try to write to cache, if the data still empty after locking
func (c *cacheRepo) WriteIfEmpty(ctx context.Context, key string, data interface{}, ttl time.Duration) (err error) {
	return c.opt.CacheClient.SetNX(ctx, key, data, ttl).Err()
}

// Get data from cache
func (c *cacheRepo) Get(ctx context.Context, key string, data interface{}) (err error) {
	res, err := c.opt.CacheClient.Get(ctx, key).Bytes()
	if err != nil {
		return
	}

	if len(res) > 0 {
		err = json.Unmarshal(res, &data)
		if err != nil {
			return
		}
	}

	return
}

func (c *cacheRepo) Delete(ctx context.Context, key string) (err error) {
	return c.opt.CacheClient.Del(ctx, key).Err()
}
