package actions

import (
	"context"
	"fmt"
	"state_machine_executor/application"
	"state_machine_executor/state_machine/localStorage"
	"strconv"
	"time"
)

type SendMessage struct {
	GenericAction
}

func NewSendMessage(genericAction *GenericAction) ActionInterface {
	return &SendMessage{GenericAction: *genericAction}
}

func (a *SendMessage) Run(ctx context.Context, application *application.Application, storage *localStorage.Storage) {
	SendMessageFunc(storage, a.Arguments, a.Return)
}

func SendMessageFunc(storage *localStorage.Storage, Arguments Arguments, Return Returns) {
	textKey := Arguments["text"]
	botService := storage.KvStorage.Get("botService").(string)
	groupId, _ := strconv.Atoi(storage.KvStorage.Get("groupId").(string))
	userId, _ := strconv.Atoi(storage.KvStorage.Get("userId").(string))
	switch botService {
	case "vk-shop-bot":
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		fmt.Print(textKey, groupId, userId, ctx)
		var err error = nil
		//_, err := grpcClient.GrpcClientSingletone.Vk.SendMessage(ctx, &VkServer.SendMessageRequest{
		//	GroupId: int64(groupId),
		//	UserId:  int64(userId),
		//	Text:    storage.KvStorage.Get(textKey).(string),
		//})
		if err != nil {
			panic(err)
		}
	}
	//fmt.Println(botService)
	//fmt.Println(storage.KvStorage.Get(textKey))
}
