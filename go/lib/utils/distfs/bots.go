package distfs

import (
	"github.com/nikhovas/diploma/go/lib/utils/distfs/bots"
	"github.com/nikhovas/diploma/go/lib/utils/distvars"
)

type Bots struct {
	DirInfo distvars.MetaDir
}

func (n *Bots) MetaCdServiceName(service string) *bots.MetaServiceName {
	return &bots.MetaServiceName{
		DirInfo: n.DirInfo.GetChild(service),
	}
}
