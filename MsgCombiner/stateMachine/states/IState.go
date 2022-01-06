package states

import "MsgCombiner/stateMachine/localStorage"

type IState interface {
	Process(storage *localStorage.Storage) (string, bool)
}
