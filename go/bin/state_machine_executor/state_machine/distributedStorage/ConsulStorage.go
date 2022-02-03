package distributedStorage

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"strconv"
)

type ConsulStorage struct {
	Storage      *api.KV
	DataBasePath string
}

func (s *ConsulStorage) GetDataPath(dataName string) string {
	return fmt.Sprintf("%s/%s", s.DataBasePath, dataName)
}

func (s *ConsulStorage) Get(key string, valueType string) interface{} {
	dataPath := s.GetDataPath(key)

	stKv, _, err := s.Storage.Get(dataPath, nil)
	if err != nil {
		panic(err)
	}

	if stKv == nil {
		return nil
	}

	val := string(stKv.Value)

	switch valueType {
	case "int":
		r, _ := strconv.Atoi(val)
		return r
	case "string":
		return val
	case "stringList":
		var res []string
		_ = json.Unmarshal([]byte(val), &res)
		return res
	default:
		return nil
	}
}

func (s *ConsulStorage) Set(key string, value interface{}) {
	dataPath := s.GetDataPath(key)

	var data []byte
	switch value.(type) {
	case []string:
		data, _ = json.Marshal(value)
	default:
		data = []byte(fmt.Sprint(value))
	}

	stKv := &api.KVPair{Key: dataPath, Value: data}
	_, err := s.Storage.Put(stKv, nil)
	if err != nil {
		panic(err)
	}
}
