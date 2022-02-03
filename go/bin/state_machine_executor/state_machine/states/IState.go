package states

import (
	"context"
	"state_machine_executor/application"
	"state_machine_executor/state_machine/localStorage"
)

type IState interface {
	Process(ctx context.Context, app *application.Application, ls *localStorage.Storage) (string, bool)
}
