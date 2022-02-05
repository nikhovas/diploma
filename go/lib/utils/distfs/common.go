package distfs

import (
	"github.com/nikhovas/diploma/go/lib/utils/distfs/common"
	"github.com/nikhovas/diploma/go/lib/utils/distvars"
)

type Common struct {
	DirInfo distvars.MetaDir
}

func (dv *Common) MetaCdShopId(shopId int) *common.MetaShopId {
	return &common.MetaShopId{
		DirInfo: dv.DirInfo.GetChild(shopId),
	}
}
