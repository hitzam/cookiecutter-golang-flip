package repository

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// ICacheRepo interface for cache repo
type ICacheRepo interface {
	WriteCache(key string, data interface{}, ttl time.Duration) (err error)
	WriteCacheIfEmpty(key string, data interface{}, ttl time.Duration) (err error)
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
func (c *cacheRepo) WriteCache(key string, data interface{}, ttl time.Duration) (err error) {
	// write data to cache
	conn := c.opt.CachePool.Get()
	defer conn.Close()

	_, err = conn.Do("SETEX", key, ttl.Seconds(), data)
	if err != nil {
		return err
	}

	return nil
}

// WriteCacheIfEmpty will try to write to cache, if the data still empty after locking
func (c *cacheRepo) WriteCacheIfEmpty(key string, data interface{}, ttl time.Duration) (err error) {
	// check whether cache value is empty
	conn := c.opt.CachePool.Get()
	defer conn.Close()

	_, err = conn.Do("GET", key)
	if err != nil {
		if err == redis.ErrNil {
			return nil //return nil as the data already set, no need to overwrite
		}

		return err
	}

	// write data to cache
	_, err = conn.Do("SETEX", key, ttl.Seconds(), data)
	if err != nil {
		return err
	}

	return nil
}
