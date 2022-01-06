package states

import (
	"MsgCombiner/stateMachine/actions"
	"MsgCombiner/stateMachine/localStorage"
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

func (s *Action) Process(storage *localStorage.Storage) (string, bool) {
	for _, action := range s.Actions {
		if action != nil {
			action.Run(storage)
		}
	}

	return s.Next, false
}
