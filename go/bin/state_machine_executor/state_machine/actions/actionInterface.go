package actions

import (
	"context"
	"state_machine_executor/application"
	"state_machine_executor/state_machine/localStorage"
)

type ActionInterface interface {
	Run(ctx context.Context, application *application.Application, storage *localStorage.Storage)
}
