package distfs

import (
	"github.com/nikhovas/diploma/go/lib/utils/distfs/products"
	"github.com/nikhovas/diploma/go/lib/utils/distvars"
)

type Products struct {
	DirInfo distvars.MetaDir
}

func (n *Products) MetaCdProducts(shopId interface{}) *products.MetaVkGroupId {
	return &products.MetaVkGroupId{
		DirInfo: n.DirInfo.GetChild(shopId),
	}
}
