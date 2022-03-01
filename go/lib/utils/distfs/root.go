package distfs

import (
	"github.com/go-redis/redis/v8"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/nikhovas/diploma/go/lib/utils/distvars"
)

type Root struct {
	DirInfo distvars.MetaDir
}

func NewRoot(redis *redis.Client, consul *consulApi.Client) *Root {
	return &Root{
		DirInfo: distvars.MetaDir{
			Redis:  redis,
			Consul: consul,
			Path:   "",
		},
	}
}

func (n *Root) CdBots() *Bots {
	return &Bots{
		DirInfo: n.DirInfo.GetChild("bots"),
	}
}

func (n *Root) CdCommon() *Common {
	return &Common{
		DirInfo: n.DirInfo.GetChild("common"),
	}
}

func (n *Root) CdProducts() *Products {
	return &Products{
		DirInfo: n.DirInfo.GetChild("products"),
	}
}
