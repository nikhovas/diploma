package meta_group_id

import (
	"github.com/nikhovas/diploma/go/lib/utils/distvars"
)

type Common struct {
	DirInfo distvars.MetaDir
}

func (n *Common) CdMessageId() *distvars.ConsulInt {
	return n.DirInfo.GetConsulIntChild("message-id")
}

func (n *Common) CdToken() *distvars.RedisString {
	return n.DirInfo.GetRedisStringChild("token")
}
