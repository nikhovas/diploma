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

func (n *Common) CdNewMsgUniqueId() *distvars.RedisCounter {
	return n.DirInfo.GetRedisCounterChild("new-msg-unique-id")
}

func (n *Common) CdGoods() *distvars.RedisDictStringString {
	return n.DirInfo.GetRedisDictStringStringChild("goods")
}

func (n *Common) CdGoodsUpdated() *distvars.RedisTimeFlag {
	return n.DirInfo.GetRedisTimeFlag("goods-updated")
}
