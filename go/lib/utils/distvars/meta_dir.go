package distvars

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	consulApi "github.com/hashicorp/consul/api"
)

type MetaDir struct {
	Redis  *redis.Client
	Consul *consulApi.Client
	Path   string
}

func (metaDir *MetaDir) GetChild(nodeName interface{}) MetaDir {
	return MetaDir{
		Redis:  metaDir.Redis,
		Consul: metaDir.Consul,
		Path:   metaDir.GetChildPath(nodeName),
	}
}

func (metaDir *MetaDir) GetChildPath(nodeName interface{}) string {
	if metaDir.Path == "" {
		return fmt.Sprintf("%v", nodeName)
	} else {
		return fmt.Sprintf("%s/%v", metaDir.Path, nodeName)
	}
}

func (metaDir *MetaDir) GetRedisStringChild(nodeName interface{}) *RedisString {
	return &RedisString{
		Storage: metaDir.Redis,
		Path:    metaDir.GetChildPath(nodeName),
	}
}

func (metaDir *MetaDir) GetRedisStringSetChild(nodeName interface{}) *RedisStringSet {
	return &RedisStringSet{
		Storage: metaDir.Redis,
		Path:    metaDir.GetChildPath(nodeName),
	}
}

func (metaDir *MetaDir) GetConsulDictStringStringChild(nodeName interface{}) *ConsulDictStringString {
	return &ConsulDictStringString{
		Storage: metaDir.Consul.KV(),
		Path:    metaDir.GetChildPath(nodeName),
	}
}

func (metaDir *MetaDir) GetConsulIntChild(nodeName interface{}) *ConsulInt {
	return &ConsulInt{
		Storage: metaDir.Consul.KV(),
		Path:    metaDir.GetChildPath(nodeName),
	}
}

func (metaDir *MetaDir) GetConsulLockChild(nodeName interface{}) *ConsulLock {
	return &ConsulLock{
		Storage: metaDir.Consul,
		Path:    metaDir.GetChildPath(nodeName),
		lock:    nil,
	}
}

func (metaDir *MetaDir) GetRedisDictStringStringChild(nodeName interface{}) *RedisDictStringString {
	return &RedisDictStringString{
		Storage: metaDir.Redis,
		Path:    metaDir.GetChildPath(nodeName),
	}
}

func (metaDir *MetaDir) GetConsulTimeMpsqQueueString(nodeName interface{}) *ConsulTimeMpsqQueueString {
	path := metaDir.GetChildPath(nodeName)
	startFromPath := fmt.Sprintf("%s/start-from", path)
	dataPath := fmt.Sprintf("%s/data", path)
	return &ConsulTimeMpsqQueueString{
		Storage: metaDir.Consul.KV(),
		Path:    path,
		StartFrom: ConsulInt{
			Storage: metaDir.Consul.KV(),
			Path:    startFromPath,
		},
		Data: ConsulDictStringString{
			Storage: metaDir.Consul.KV(),
			Path:    dataPath,
		},
	}
}

func (metaDir *MetaDir) GetRedisCounterChild(nodeName interface{}) *RedisCounter {
	return &RedisCounter{
		Storage: metaDir.Redis,
		Path:    metaDir.GetChildPath(nodeName),
	}
}

func (metaDir *MetaDir) GetRedisTimeFlag(nodeName interface{}) *RedisTimeFlag {
	return &RedisTimeFlag{
		Storage: metaDir.Redis,
		Path:    metaDir.GetChildPath(nodeName),
	}
}
