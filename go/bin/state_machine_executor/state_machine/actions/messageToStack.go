package actions

import (
	"state_machine_executor/application"
	"state_machine_executor/state_machine/localStorage"
	"context"
)

type MessageToStack struct {
	GenericAction
}

func NewMessageToStack(genericAction *GenericAction) ActionInterface {
	return &MessageToStack{GenericAction: *genericAction}
}

func (a *MessageToStack) Run(ctx context.Context, application *application.Application, storage *localStorage.Storage) {
	MessageToStackFunc(storage, a.Arguments, a.Return)
}

func MessageToStackFunc(storage *localStorage.Storage, Arguments Arguments, Return Returns) {
	msg := storage.KvStorage.Get("message")
	storage.MessageDeque.PushFront(msg)
}
