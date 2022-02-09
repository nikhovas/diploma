package actions

import (
	"context"
	"state_machine_executor/application"
	"state_machine_executor/state_machine/localStorage"
	"state_machine_executor/utils"
	"strconv"
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
	msg := storage.KvStorage.Get("message").(string)
	msgId, _ := strconv.Atoi(storage.KvStorage.Get("messageId").(string))
	storage.MessageDeque.PushFront(&utils.MessageInfo{
		Text: msg,
		Id:   msgId,
	})
}
