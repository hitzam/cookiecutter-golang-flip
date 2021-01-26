package repository

import (
	"context"
	"time"
)

// ICacheRepo interface for cache repo
type ICacheRepo interface {
	WriteCache(ctx context.Context, key string, data interface{}, ttl time.Duration) (err error)
	WriteCacheIfEmpty(ctx context.Context, key string, data interface{}, ttl time.Duration) (err error)
}

type cacheRepo struct {
	opt Option
}

// NewCacheRepository initiate cache repo
func NewCacheRepository(opt Option) ICacheRepo {
	return &cacheRepo{
		opt: opt,
	}
}

// WriteCache this will and must write the data to cache with corresponding key using locking
func (c *cacheRepo) WriteCache(ctx context.Context, key string, data interface{}, ttl time.Duration) (err error) {
	return c.opt.CacheClient.SetEX(ctx, key, data, ttl).Err()
}

// WriteCacheIfEmpty will try to write to cache, if the data still empty after locking
func (c *cacheRepo) WriteCacheIfEmpty(ctx context.Context, key string, data interface{}, ttl time.Duration) (err error) {
	return c.opt.CacheClient.SetNX(ctx, key, data, ttl).Err()
}
