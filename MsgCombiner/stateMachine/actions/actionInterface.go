package actions

import "MsgCombiner/stateMachine/localStorage"

type ActionInterface interface {
	Run(storage *localStorage.Storage)
}
