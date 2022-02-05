package meta_service_name

import (
	"github.com/nikhovas/diploma/go/lib/utils/distvars"
)

type Common struct {
	DirInfo distvars.MetaDir
}

func (n *Common) CdEnabledBots() *distvars.RedisStringSet {
	return n.DirInfo.GetRedisStringSetChild("enabled-bots")
}
