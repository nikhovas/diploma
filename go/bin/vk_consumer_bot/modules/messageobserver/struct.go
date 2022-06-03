package messageobserver

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	consumerActions "github.com/nikhovas/diploma/go/lib/proto/consumer_actions"
	ctrl "github.com/nikhovas/diploma/go/lib/proto/controller"
	"github.com/nikhovas/diploma/go/lib/utils/consts"
	"github.com/nikhovas/diploma/go/lib/utils/distfs/bots"
	"github.com/nikhovas/diploma/go/lib/utils/foundation/modules/norun"
	longPullServer "github.com/nikhovas/diploma/go/lib/vk/long_pull_server"
	"log"
	"strconv"
	"vk_consumer_bot/localUtils"
	"vk_consumer_bot/modules/eventsqueuewriter"
)

type MessageObserver struct {
	norun.NoRun
	eventsQueueWriter *eventsqueuewriter.EventsQueueWriter
	vkDistFsBase *bots.MetaServiceName
	ctrlClient ctrl.ControllerClient
}

func FromConfig(config Config, eventsQueueWriter *eventsqueuewriter.EventsQueueWriter, vkDistFsBase *bots.MetaServiceName, ctrlClient ctrl.ControllerClient) *MessageObserver {
	return &MessageObserver{
		eventsQueueWriter: eventsQueueWriter,
		vkDistFsBase: vkDistFsBase,
		ctrlClient: ctrlClient,
	}
}

func (mo *MessageObserver) Callback(groupId int, ts int, update longPullServer.UpdateObject) {
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

	res, err := mo.ctrlClient.GetShopIdByKey(
		context.Background(),
		localUtils.NewVkGroupIdShopKey(groupId),
	)
	if err != nil {
		log.Println(err)
		return
	}
	ae.ShopId = res.ShopId

	m := jsonpb.Marshaler{}
	userActionString, _ := m.MarshalToString(&userAction)

	groupDistFs := mo.vkDistFsBase.MetaCdGroupId(groupId)
	userDistFs := groupDistFs.MetaCdUserId(userId)
	actionsDistFs := userDistFs.CdActions()

	if err := actionsDistFs.Set(context.Background(), ts, userActionString); err != nil {
		log.Println(err)
		return
	}

	err = mo.eventsQueueWriter.SendMessage(ae)
	if err != nil {
		log.Println(err)
	}
}
