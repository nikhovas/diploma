package actions

import (
	"MsgCombiner/Application"
	"MsgCombiner/stateMachine/localStorage"
	"context"
)

type ActionInterface interface {
	Run(ctx context.Context, application *Application.Application, storage *localStorage.Storage)
}
