package distvars

import (
	"context"
	consulApi "github.com/hashicorp/consul/api"
)

type ConsulStringDistVarBase struct {
	storage *consulApi.KV
	path    string
}

func (dv *ConsulStringDistVarBase) Get(ctx context.Context) (string, uint64, error) {
	res, _, err := dv.storage.Get(dv.path, nil)
	if err != nil {
		return "", 0, err
	}

	return string(res.Value), res.ModifyIndex, err
}

func (dv *ConsulStringDistVarBase) Put(ctx context.Context, value string) error {
	_, err := dv.storage.Put(&consulApi.KVPair{Key: dv.path, Value: []byte(value)}, nil)
	return err
}

func (dv *ConsulStringDistVarBase) CAS(ctx context.Context, value string, index uint64) (bool, error) {
	stKv := consulApi.KVPair{Key: dv.path, Value: []byte(value), ModifyIndex: index}
	success, _, err := dv.storage.CAS(&stKv, nil)
	return success, err
}
