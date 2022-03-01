package products

import (
	"github.com/nikhovas/diploma/go/lib/utils/distvars"
)

type MetaVkGroupId struct {
	DirInfo distvars.MetaDir
}

func (n *MetaVkGroupId) CdProducts() *distvars.RedisDictStringString {
	return n.DirInfo.GetRedisDictStringStringChild("products")
}

func (n *MetaVkGroupId) CdUpdateFlag() *distvars.RedisTimeFlag {
	return n.DirInfo.GetRedisTimeFlag("updated")
}
