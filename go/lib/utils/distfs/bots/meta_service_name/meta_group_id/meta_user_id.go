package meta_group_id

import (
	"fmt"
	"github.com/nikhovas/diploma/go/lib/utils/distvars"
)

type MetaUserId struct {
	DirInfo distvars.MetaDir
}

func (n *MetaUserId) CdActions() *distvars.ConsulTimeMpsqQueueString {
	return n.DirInfo.GetConsulTimeMpsqQueueString("actions")
}

func (n *MetaUserId) GetStateMachinePath() string {
	return fmt.Sprintf("%s/state-machine", n.DirInfo.Path)
}

func (n *MetaUserId) GetLock() *distvars.ConsulLock {
	return n.DirInfo.GetConsulLockChild("lock")
}
