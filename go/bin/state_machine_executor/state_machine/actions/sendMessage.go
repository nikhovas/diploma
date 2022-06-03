package actions

import (
	"context"
	"fmt"
	"github.com/nikhovas/diploma/go/lib/proto/common"
	"github.com/nikhovas/diploma/go/lib/proto/consumer_messages"
	"state_machine_executor/coremodules"
	"state_machine_executor/state_machine/localStorage"
	"strconv"
)

type SendMessage struct {
	GenericAction
}

func NewSendMessage(genericAction *GenericAction) ActionInterface {
	return &SendMessage{GenericAction: *genericAction}
}

func (a *SendMessage) Run(ctx context.Context, cm *coremodules.CoreModules, storage *localStorage.Storage) {
	SendMessageFunc(ctx, cm, storage, a.Arguments, a.Return)
}

func SendMessageFunc(
	ctx context.Context,
	cm *coremodules.CoreModules,
	storage *localStorage.Storage,
	Arguments Arguments,
	Return Returns,
) {
	textKey := Arguments["text"]
	botService := storage.KvStorage.Get("botService").(string)
	groupId, _ := strconv.Atoi(storage.KvStorage.Get("groupId").(string))
	userId, _ := strconv.Atoi(storage.KvStorage.Get("userId").(string))

	var text string
	preText := storage.KvStorage.Get(textKey)
	if t, ok := preText.(string); ok {
		text = t
	}

	err := cm.ConsumerMessageSender.SendMessage(&consumer_messages.MessageToConsumer{
		Uuid: "",
		MsgLocation: &common.MsgLocation{
			Service: botService,
			GroupId: int64(groupId),
			UserId:  int64(userId),
		},
		Text:  text,
		Reply: nil,
	})
	if err != nil {
		fmt.Println(err)
	}
}
