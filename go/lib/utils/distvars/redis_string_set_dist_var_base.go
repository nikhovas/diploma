package distvars

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisStringSetUpdate struct {
	Added   []string
	Deleted []string
	Error   error
}

type RedisStringSet struct {
	Storage *redis.Client
	Path    string
}

func (dv *RedisStringSet) Get(ctx context.Context) ([]string, error) {
	res, err := dv.Storage.LRange(ctx, dv.Path, 0, -1).Result()
	if err != nil {
		return []string{}, err
	}
	return res, nil
}

func (dv *RedisStringSet) Add(ctx context.Context, value string) error {
	return dv.Storage.LPush(ctx, dv.Path, value).Err()
}

func (dv *RedisStringSet) Delete(ctx context.Context, value string) error {
	return dv.Storage.LRem(ctx, dv.Path, 0, value).Err()
}

func (dv *RedisStringSet) GetUpdates(
	ctx context.Context,
	baseValue map[string]struct{},
) RedisStringSetUpdate {
	res, err := dv.Get(ctx)
	if err != nil {
		return RedisStringSetUpdate{Error: err}
	}

	newValue := make(map[string]struct{})
	for _, item := range res {
		newValue[item] = struct{}{}
	}

	deleted := make([]string, 0)
	for val, _ := range baseValue {
		if _, exists := newValue[val]; !exists {
			deleted = append(deleted, val)
			delete(baseValue, val)
		}
	}

	added := make([]string, 0)
	for val, _ := range newValue {
		if _, exists := baseValue[val]; !exists {
			added = append(added, val)
			baseValue[val] = struct{}{}
		}
	}

	return RedisStringSetUpdate{
		Added:   added,
		Deleted: deleted,
		Error:   nil,
	}
}

func (dv *RedisStringSet) getUpdatesChanWorker(
	ctx context.Context,
	baseValue map[string]struct{},
	checkDelay time.Duration,
	channel chan<- RedisStringSetUpdate,
) chan RedisStringSetUpdate {
	for {
		closeChan := false
		select {
		case _ = <-ctx.Done():
			closeChan = true
		default:
		}
		if ctx.Err() != nil {
			closeChan = true
		}

		if closeChan {
			close(channel)
		}

		channel <- dv.GetUpdates(ctx, baseValue)
		time.Sleep(checkDelay)
	}
}

func (dv *RedisStringSet) GetUpdatesChan(
	ctx context.Context,
	baseValue map[string]struct{},
	checkDelay time.Duration,
) chan RedisStringSetUpdate {
	channel := make(chan RedisStringSetUpdate, 100)
	go dv.getUpdatesChanWorker(ctx, baseValue, checkDelay, channel)
	return channel
}
