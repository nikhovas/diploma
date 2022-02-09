package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	consumerActions "github.com/nikhovas/diploma/go/lib/proto/consumer_actions"
	ctrl "github.com/nikhovas/diploma/go/lib/proto/controller"
	"github.com/nikhovas/diploma/go/lib/utils/consts"
	longPullServer "github.com/nikhovas/diploma/go/lib/vk/long_pull_server"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strconv"
	"vk_consumer_bot/application"
)

func callback(
	app *application.Application,
	groupId int,
	ts int,
	update longPullServer.UpdateObject,
) {
	fmt.Printf("Start callback for %d\n", groupId)

	ro := update.Object

	var userId int
	ae := &consumerActions.ActionEvent{
		UserId: "",
		BotId:  strconv.Itoa(groupId),
		Time:   uint64(ts),
	}
	userAction := consumerActions.UserAction{
		Time:   uint64(ts),
		Object: nil,
	}

	switch v := ro.(type) {
	case *longPullServer.NewMessageObject:
		userAction.Object = &consumerActions.UserAction_NewMessage{
			NewMessage: &consumerActions.NewMessage{Text: v.Body, Id: uint64(v.Id)},
		}

		userId = v.UserId
		fmt.Printf("Found message with text %s\n", v.Body)
	default:
		fmt.Println("Unsupported message type")
		return
	}

	ae.UserId = strconv.Itoa(userId)
	ae.ServiceName = consts.VkConsumerBotServiceName

	res, err := app.CtrlClient.GetShopIdByKey(context.Background(), &ctrl.ShopKey{
		Key: &ctrl.ShopKey_Common{
			Common: &ctrl.CommonShopKey{
				CommonKey: &ctrl.CommonShopKey_VkGroupId{
					VkGroupId: int64(groupId),
				},
			},
		},
	})
	if err != nil {
		log.Println(err)
		return
	}
	ae.ShopId = res.ShopId

	m := jsonpb.Marshaler{}
	userActionString, _ := m.MarshalToString(&userAction)

	groupDistFs := app.VkDistFsBase.MetaCdGroupId(groupId)
	userDistFs := groupDistFs.MetaCdUserId(userId)
	actionsDistFs := userDistFs.CdActions()

	if err := actionsDistFs.Set(context.Background(), ts, userActionString); err != nil {
		log.Println(err)
		return
	}

	aeString, _ := m.MarshalToString(ae)
	err = amqpChannel.Publish("", amqpQueue.Name, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(aeString),
		})
	if err != nil {
		log.Println(err)
	}
}
