package v0

import (
	"state_machine_executor/state_machine"
	"state_machine_executor/state_machine/actions"
	"state_machine_executor/state_machine/localStorage"
	"state_machine_executor/state_machine/states"
)

type Action struct {
	Name      string            `json:"name"`
	Arguments map[string]string `json:"arguments"`
	Return    map[string]string `json:"return"`
}

func (a *Action) ToInternal() actions.GenericAction {
	return actions.GenericAction{
		Name:      a.Name,
		Arguments: a.Arguments,
		Return:    a.Return,
	}
}

type ActionState struct {
	Actions []Action `json:"actions"`
	Next    string   `json:"next"`
}

func (s *ActionState) ToInternal() *states.Action {
	newActions := make([]actions.ActionInterface, 4)
	for _, a := range s.Actions {
		ga := a.ToInternal()
		newActions = append(newActions, actions.ActionBuilders[ga.Name](&ga))
	}

	return &states.Action{
		Actions: newActions,
		Next:    s.Next,
	}
}

type MessageWaitState struct {
	MessageToStack bool   `json:"messageToStack"`
	Next           string `json:"next"`
}

func (s *MessageWaitState) ToInternal() *states.MessageWait {
	return &states.MessageWait{
		MessageToStack: s.MessageToStack,
		Next:           s.Next,
	}
}

type State struct {
	ActionStateData      *ActionState      `json:"actionStateData"`
	MessageWaitStateData *MessageWaitState `json:"messageWaitStateData"`
	MessageClassifierStateData *states.MessageClassifier
	CheckOrderStateData *states.CheckOrder
	ConfirmOrderStateData *states.ConfirmOrder
}

func (s *State) ToInternal() states.IState {
	if s.ActionStateData != nil {
		return s.ActionStateData.ToInternal()
	} else if s.MessageWaitStateData != nil {
		return s.MessageWaitStateData.ToInternal()
	} else if s.MessageClassifierStateData != nil {
		return s.MessageClassifierStateData
	} else if s.CheckOrderStateData != nil {
		return s.CheckOrderStateData
	} else if s.ConfirmOrderStateData != nil {
		return s.ConfirmOrderStateData
	} else {
		return nil
	}
}

type States map[string]State

func (s *States) ToInternal() map[string]states.IState {
	newStates := make(map[string]states.IState, 4)
	for key, value := range *s {
		newStates[key] = value.ToInternal()
	}

	return newStates
}

type DataElement struct {
	Memory  string      `json:"memory"`
	Type    string      `json:"type"`
	Default interface{} `json:"default"`
}

func (d *DataElement) ToInternal() localStorage.DataElement {
	return localStorage.DataElement{
		Memory:  d.Memory,
		Type:    d.Type,
		Default: d.Default,
	}
}

type Model struct {
	StartState string                 `json:"startState"`
	States     States                 `json:"states"`
	Data       map[string]DataElement `json:"data"`
}

func (m *Model) ToInternal() stateMachine.Model {
	newData := make(map[string]localStorage.DataElement)
	for key, value := range m.Data {
		newData[key] = value.ToInternal()
	}

	return stateMachine.Model{
		StartState: m.StartState,
		States:     m.States.ToInternal(),
		Data:       newData,
	}
}
