package distvars

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisCounter struct {
	Storage *redis.Client
	Path    string
}

func (dv *RedisCounter) Incr(ctx context.Context) (int, error) {
	res, err := dv.Storage.Incr(ctx, dv.Path).Result()
	if err != nil {
		return 0, err
	}
	return int(res), nil
}
