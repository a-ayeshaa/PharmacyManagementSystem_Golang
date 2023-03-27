package cache

import (
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"PharmaProject/internal/conn"
)

// RedisCache represents the redis driver
type RedisCache struct {
	rds        conn.RedisClient
	prefix     string
	defaultTTL time.Duration
}

// NewRedisCache is the factory function for the  redis cache instance
func NewRedisCache(r conn.RedisClient, pfx string, dtl time.Duration) Cacher {
	return &RedisCache{
		rds:        r,
		prefix:     pfx,
		defaultTTL: dtl,
	}
}

// SetCache set cache on redis
func (r *RedisCache) SetCache(key string, val string, ttl ...time.Duration) error {
	var cacheTTL time.Duration

	if len(ttl) == 0 { // if optional ttl for cache is not provided, read the ttl from consul as default
		cacheTTL = r.defaultTTL
	} else {
		cacheTTL = ttl[0]
		if cacheTTL == (0 * time.Second) { //if 0 second cache is provided assign the default
			cacheTTL = r.defaultTTL
		}
	}

	if !strings.HasPrefix(key, r.prefix) {
		key = r.BuildKey(key)
	}

	if err := r.rds.Set(key, val, cacheTTL).
		Err(); err != nil {
		return err
	}

	return nil
}

// GetCache returns a cache value against the key provided from redis
func (r *RedisCache) GetCache(key string) (string, error) {
	if !strings.HasPrefix(key, r.prefix) {
		key = r.BuildKey(key)
	}

	bb, erR := r.rds.Get(key).Result()
	if erR != nil {
		if erR != redis.Nil {
			return "", erR
		}
	}

	return bb, nil
}

// ClearCache clears a cache matching the patter of the redis key
func (r *RedisCache) ClearCache(pattern string) error {
	if !strings.HasPrefix(pattern, r.prefix) {
		pattern = r.BuildKey(pattern)
	}

	keys, err := r.rds.Keys(pattern).Result()

	if err != nil {
		if err != redis.Nil {
			log.Println("redis failed: ", err)
			return err
		}
	}

	if len(keys) > 0 {
		if err := r.rds.Del(keys...).Err(); err != nil {
			return err
		}
	}

	return nil
}

// Exists checks if a redis key exists or not, because the function redis package provides is sh*t
func (r *RedisCache) Exists(key string) (bool, error) {
	if !strings.HasPrefix(key, r.prefix) {
		key = r.BuildKey(key)
	}

	_, err := r.rds.Get(key).Result()

	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// BuildKey builds a valid key for redis cache
func (r *RedisCache) BuildKey(keywords ...string) string {
	kk := []string{
		strings.TrimSuffix(r.prefix, "_"),
	}

	for _, kw := range keywords {
		if kw != "" && kw != r.prefix {
			kk = append(kk, kw)
		}
	}

	key := strings.Join(kk, "_")

	return key
}
