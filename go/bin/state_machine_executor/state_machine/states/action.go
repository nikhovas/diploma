package states

import (
	"context"
	"state_machine_executor/coremodules"
	"state_machine_executor/state_machine/actions"
	"state_machine_executor/state_machine/localStorage"
)

type Action struct {
	Actions []actions.ActionInterface
	Next    string
}

func NewAction(actionElements []actions.GenericAction, next string) *Action {
	var internalActionElements []actions.ActionInterface
	for _, act := range actionElements {
		internalActionElements = append(internalActionElements, actions.ActionBuilders[act.Name](&act))
	}

	return &Action{
		Actions: internalActionElements,
		Next:    next,
	}

}

func (s *Action) Process(ctx context.Context, cm *coremodules.CoreModules, ls *localStorage.Storage) (string, bool) {
	for _, action := range s.Actions {
		if action != nil {
			action.Run(ctx, cm, ls)
		}
	}

	return s.Next, false
}
