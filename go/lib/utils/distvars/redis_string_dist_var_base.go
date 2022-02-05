package distvars

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisString struct {
	Storage *redis.Client
	Path    string
}

func (dv *RedisString) Get(ctx context.Context) (string, error) {
	return dv.Storage.Get(ctx, dv.Path).Result()
}

func (dv *RedisString) Set(ctx context.Context, value string) error {
	return dv.Storage.Set(ctx, dv.Path, value, 0).Err()
}
