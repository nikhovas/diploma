package actions

import (
	"MsgCombiner/stateMachine/localStorage"
)

type MessageToStack struct {
	GenericAction
}

func NewMessageToStack(genericAction *GenericAction) ActionInterface {
	return &MessageToStack{GenericAction: *genericAction}
}

func (a *MessageToStack) Run(storage *localStorage.Storage) {
	MessageToStackFunc(storage, a.Arguments, a.Return)
}

func MessageToStackFunc(storage *localStorage.Storage, Arguments Arguments, Return Returns) {
	msg := storage.KvStorage.Get("message")
	storage.MessageDeque.PushFront(msg)
}
