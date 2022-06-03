package states

import (
	"context"
	"state_machine_executor/coremodules"
	"state_machine_executor/state_machine/localStorage"
)

type IState interface {
	Process(ctx context.Context, cm *coremodules.CoreModules, ls *localStorage.Storage) (string, bool)
}
