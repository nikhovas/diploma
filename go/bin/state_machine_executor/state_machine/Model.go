package stateMachine

import (
	"state_machine_executor/state_machine/localStorage"
	"state_machine_executor/state_machine/states"
)

type Model struct {
	StartState string
	States     map[string]states.IState
	Data       map[string]localStorage.DataElement
}
