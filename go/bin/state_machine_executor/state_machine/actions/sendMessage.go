package actions

import (
	"context"
	"fmt"
	"github.com/nikhovas/diploma/go/lib/proto/common"
	"github.com/nikhovas/diploma/go/lib/proto/consumer_bot"
	"state_machine_executor/application"
	"state_machine_executor/state_machine/localStorage"
	"strconv"
)

type SendMessage struct {
	GenericAction
}

func NewSendMessage(genericAction *GenericAction) ActionInterface {
	return &SendMessage{GenericAction: *genericAction}
}

func (a *SendMessage) Run(ctx context.Context, app *application.Application, storage *localStorage.Storage) {
	SendMessageFunc(ctx, app, storage, a.Arguments, a.Return)
}

func SendMessageFunc(
	ctx context.Context,
	app *application.Application,
	storage *localStorage.Storage,
	Arguments Arguments,
	Return Returns,
) {
	textKey := Arguments["text"]
	botService := storage.KvStorage.Get("botService").(string)
	groupId, _ := strconv.Atoi(storage.KvStorage.Get("groupId").(string))
	userId, _ := strconv.Atoi(storage.KvStorage.Get("userId").(string))

	_, err := app.VksClient.SendSimpleMessage(
		ctx,
		&consumer_bot.SendSimpleMessageRequest{
			Uuid: "",
			Info: &consumer_bot.SimpleMessageInformation{
				MsgLocation: &common.MsgLocation{
					Service: botService,
					GroupId: int64(groupId),
					UserId:  int64(userId),
				},
				Text: textKey,
			},
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
}
