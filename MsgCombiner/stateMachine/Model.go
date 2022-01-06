package stateMachine

import (
	"MsgCombiner/stateMachine/localStorage"
	"MsgCombiner/stateMachine/states"
)

type Model struct {
	StartState string
	States     map[string]states.IState
	Data       map[string]localStorage.DataElement
}
