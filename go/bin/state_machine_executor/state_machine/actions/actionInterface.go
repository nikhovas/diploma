package actions

import (
	"context"
	"state_machine_executor/coremodules"
	"state_machine_executor/state_machine/localStorage"
)

type ActionInterface interface {
	Run(ctx context.Context, cm *coremodules.CoreModules, storage *localStorage.Storage)
}
