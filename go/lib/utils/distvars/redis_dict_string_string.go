package distvars

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
)

type RedisDictStringString struct {
	Storage *redis.Client
	Path    string
}

func (dv *RedisDictStringString) Get(ctx context.Context, key string) (string, error) {
	key = fmt.Sprintf("%s/%s", dv.Path, key)
	res, err := dv.Storage.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

func (dv *RedisDictStringString) Set(ctx context.Context, key string, value string) error {
	key = fmt.Sprintf("%s/%s", dv.Path, key)
	err := dv.Storage.Set(ctx, key, value, 0).Err()
	return err
}

func (dv *RedisDictStringString) Keys(ctx context.Context) ([]string, error) {
	keys := fmt.Sprintf("%s/*", dv.Path)
	list, err := dv.Storage.Keys(ctx, keys).Result()
	if err != nil {
		return []string{}, err
	}

	for i, kv := range list {
		list[i] = strings.TrimPrefix(kv, dv.Path)
	}

	return list, nil
}

func (dv *RedisDictStringString) Delete(ctx context.Context, key string) error {
	key = fmt.Sprintf("%s/%s", dv.Path, key)
	return dv.Storage.Del(ctx, key).Err()
}
