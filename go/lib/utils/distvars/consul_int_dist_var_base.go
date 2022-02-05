package distvars

import (
	"context"
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"strconv"
)

type NotFoundError struct {
	Key string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Not found: %s", e.Key)
}

type ConsulInt struct {
	Storage *consulApi.KV
	Path    string
}

func (dv *ConsulInt) Get(ctx context.Context) (int, uint64, error) {
	res, _, err := dv.Storage.Get(dv.Path, nil)
	if err != nil {
		return 0, 0, err
	} else if res == nil {
		return 0, 0, &NotFoundError{Key: dv.Path}
	}
	resStr, err := strconv.Atoi(string(res.Value))

	return resStr, res.ModifyIndex, err
}

func (dv *ConsulInt) Put(ctx context.Context, value int) error {
	_, err := dv.Storage.Put(&consulApi.KVPair{Key: dv.Path, Value: []byte(strconv.Itoa(value))}, nil)
	return err
}

func (dv *ConsulInt) CAS(ctx context.Context, value int, index uint64) (bool, error) {
	stKv := consulApi.KVPair{Key: dv.Path, Value: []byte(strconv.Itoa(value)), ModifyIndex: index}
	success, _, err := dv.Storage.CAS(&stKv, nil)
	return success, err
}

func (dv *ConsulInt) SwapIfGreater(ctx context.Context, value int) (int, int, error) {
	for {
		distVal, index, err := dv.Get(ctx)
		if err != nil {
			switch err.(type) {
			case *NotFoundError:
				index = 0
				distVal = value
			default:
				return 0, 0, err
			}
		}

		if value < distVal {
			return distVal, distVal, nil
		} else {
			success, err := dv.CAS(ctx, value, index)
			if err != nil {
				return 0, 0, err
			} else if success {
				return value, distVal, nil
			}
		}
	}
}
