package meta_service_name

import (
	"github.com/nikhovas/diploma/go/lib/utils/distfs/bots/meta_service_name/meta_group_id"
	"github.com/nikhovas/diploma/go/lib/utils/distvars"
)

type MetaGroupId struct {
	DirInfo distvars.MetaDir
}

func (n *MetaGroupId) MetaCdUserId(userId interface{}) *meta_group_id.MetaUserId {
	return &meta_group_id.MetaUserId{
		DirInfo: n.DirInfo.GetChild(userId),
	}
}

func (n *MetaGroupId) CdCommon() *meta_group_id.Common {
	return &meta_group_id.Common{
		DirInfo: n.DirInfo.GetChild("common"),
	}
}
