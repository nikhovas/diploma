package distvars

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisTimeFlag struct {
	Storage *redis.Client
	Path    string
}

func (dv *RedisTimeFlag) Set(ctx context.Context, expiration time.Duration) (bool, error) {
	_, err := dv.Storage.SetNX(ctx, dv.Path, true, expiration).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
