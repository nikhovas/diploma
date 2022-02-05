package bots

import (
	"github.com/nikhovas/diploma/go/lib/utils/distfs/bots/meta_service_name"
	"github.com/nikhovas/diploma/go/lib/utils/distvars"
)

type MetaServiceName struct {
	DirInfo distvars.MetaDir
}

func (n *MetaServiceName) MetaCdGroupId(groupId interface{}) *meta_service_name.MetaGroupId {
	return &meta_service_name.MetaGroupId{
		DirInfo: n.DirInfo.GetChild(groupId),
	}
}

func (n *MetaServiceName) CdCommon() *meta_service_name.Common {
	return &meta_service_name.Common{
		DirInfo: n.DirInfo.GetChild("common"),
	}
}
