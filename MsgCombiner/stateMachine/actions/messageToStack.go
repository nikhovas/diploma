package actions

import (
	"MsgCombiner/Application"
	"MsgCombiner/stateMachine/localStorage"
	"context"
)

type MessageToStack struct {
	GenericAction
}

func NewMessageToStack(genericAction *GenericAction) ActionInterface {
	return &MessageToStack{GenericAction: *genericAction}
}

func (a *MessageToStack) Run(ctx context.Context, application *Application.Application, storage *localStorage.Storage) {
	MessageToStackFunc(storage, a.Arguments, a.Return)
}

func MessageToStackFunc(storage *localStorage.Storage, Arguments Arguments, Return Returns) {
	msg := storage.KvStorage.Get("message")
	storage.MessageDeque.PushFront(msg)
}
