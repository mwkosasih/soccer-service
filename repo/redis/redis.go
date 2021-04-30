package redis

import (
	"context"
	"time"
)

type RedisDB struct {
}

func (r *RedisDB) Get(ctx context.Context, key string) ([]byte, error) {
	return RedisClient.Get(ctx, key).Bytes()
}

func (r *RedisDB) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	return RedisClient.Set(ctx, key, value, duration).Err()
}

func NewRedis() *RedisDB {
	return &RedisDB{}
}
