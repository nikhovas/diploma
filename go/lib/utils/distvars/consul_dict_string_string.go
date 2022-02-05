package distvars

import (
	"context"
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"strings"
)

type ConsulDictStringString struct {
	Storage *consulApi.KV
	Path    string
}

func (dv *ConsulDictStringString) Get(ctx context.Context, key string) (string, error) {
	key = fmt.Sprintf("%s/%s", dv.Path, key)
	res, _, err := dv.Storage.Get(key, nil)
	if err != nil {
		return "", err
	}
	return string(res.Value), nil
}

func (dv *ConsulDictStringString) Set(ctx context.Context, key string, value string) error {
	key = fmt.Sprintf("%s/%s", dv.Path, key)
	_, err := dv.Storage.Put(&consulApi.KVPair{Key: key, Value: []byte(value)}, nil)
	return err
}

func (dv *ConsulDictStringString) Keys(ctx context.Context) ([]string, error) {
	basePath := dv.Path + "/"

	list, _, err := dv.Storage.List(basePath, nil)
	if err != nil {
		return []string{}, err
	}

	result := make([]string, 0, len(list))
	for _, kv := range list {
		key := strings.TrimPrefix(kv.Key, basePath)
		result = append(result, key)
	}

	return result, nil
}

func (dv *ConsulDictStringString) GetAll(ctx context.Context) (map[string]string, error) {
	basePath := dv.Path + "/"

	list, _, err := dv.Storage.List(basePath, nil)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, kv := range list {
		key := strings.TrimPrefix(kv.Key, basePath)
		result[key] = string(kv.Value)
	}

	return result, nil
}

func (dv *ConsulDictStringString) Delete(ctx context.Context, key string) error {
	key = fmt.Sprintf("%s/%s", dv.Path, key)
	_, err := dv.Storage.Delete(key, nil)
	return err
}
