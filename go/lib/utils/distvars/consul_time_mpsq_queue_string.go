package distvars

import (
	"context"
	consulApi "github.com/hashicorp/consul/api"
	"sort"
	"strconv"
	"time"
)

type ConsulTimeMpsqQueueString struct {
	Storage   *consulApi.KV
	Path      string
	StartFrom ConsulInt
	Data      ConsulDictStringString
}

func (dv *ConsulTimeMpsqQueueString) Pop(ctx context.Context, loopTime time.Duration) ([]string, error) {
	startFrom := 0
	startFrom, _, err := dv.StartFrom.Get(context.Background())
	if err != nil {
		switch err.(type) {
		case *NotFoundError:
			startFrom = 0
		default:
			return nil, err
		}
	}

	newestTs := 0
	listLength := 0
	var result map[int]string
	var times []int

	for {
		data, err := dv.Data.GetAll(context.Background())
		if err != nil {
			return nil, err
		}

		newListLength := len(data)
		if newListLength == listLength {
			break
		}

		listLength = newListLength

		for key, value := range data {
			ts, _ := strconv.Atoi(key)

			if ts < startFrom {
				_ = dv.Data.Delete(context.Background(), key)
				listLength--
			} else {
				result[ts] = value
				times = append(times, ts)
				if ts > newestTs {
					newestTs = ts
				}
			}
		}

		if loopTime == 0 {
			break
		}
		time.Sleep(loopTime)
	}

	err = dv.StartFrom.Put(ctx, startFrom)
	if err != nil {
		return nil, err
	}

	sort.Ints(times)

	var resultList []string
	for _, i := range times {
		resultList = append(resultList, result[i])
	}

	return resultList, nil
}

func (dv *ConsulTimeMpsqQueueString) Set(ctx context.Context, valueTime int, value string) error {
	return dv.Data.Set(ctx, strconv.Itoa(valueTime), value)
}

func (dv *ConsulTimeMpsqQueueString) GetNewestTimeValue(ctx context.Context) (*int, error) {
	startFrom, _, err := dv.StartFrom.Get(context.Background())
	if err != nil {
		switch err.(type) {
		case *NotFoundError:
			return nil, nil
		default:
			return nil, err
		}
	}

	data, err := dv.Data.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	exists := false
	newestTime := 0

	for key, _ := range data {
		ts, _ := strconv.Atoi(key)
		if ts <= startFrom {
			_ = dv.Data.Delete(ctx, key)
		} else {
			exists = true
			if ts > newestTime {
				newestTime = ts
			}
		}
	}

	if !exists {
		return nil, nil
	} else {
		return &newestTime, nil
	}
}
