package eventobserver

import (
	"bytes"
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/nikhovas/diploma/go/lib/proto/consumer_actions"
	actions "github.com/nikhovas/diploma/go/lib/proto/consumer_actions"
	"github.com/nikhovas/diploma/go/lib/utils/distfs"
	"github.com/nikhovas/diploma/go/lib/utils/distfs/bots/meta_service_name/meta_group_id"
	"github.com/nikhovas/diploma/go/lib/utils/distvars"
	"github.com/nikhovas/diploma/go/lib/utils/foundation/modules/norun"
	"log"
	configExamples "state_machine_executor/config_examples"
	"state_machine_executor/coremodules"
	stateMachine "state_machine_executor/state_machine"
	"state_machine_executor/state_machine/distributedStorage"
	"state_machine_executor/utils"
	"time"
)

type EventObserver struct {
	norun.NoRun
	cm *coremodules.CoreModules
}

func FromConfig(config Config, cm *coremodules.CoreModules) *EventObserver {
	return &EventObserver{cm: cm}
}

func (eo *EventObserver) Observe(ctx context.Context, ae *actions.ActionEvent) {
	userDir := distfs.NewRoot(eo.cm.RedisClient, eo.cm.ConsulClient).CdBots().MetaCdServiceName(ae.ServiceName).
		MetaCdGroupId(ae.BotId).MetaCdUserId(ae.UserId)

	eo.ProcessLockedPart(ctx, userDir, ae)

	// send to self if any values exists
	actionsStartFromVar := userDir.CdActions()
	value, err := actionsStartFromVar.GetNewestTimeValue(ctx)
	if err != nil {
		return
	}
	if value == nil {
		return
	}

	ae.Time = uint64(*value)

	err = eo.cm.EventsQueueWriter.SendMessage(ae)
	if err != nil {
		log.Println(err)
	}
}

func MessageGenerator(userActionList []*comsumer_actions.UserAction) []utils.MessageInfo {
	var messages []utils.MessageInfo

	for _, userActionElem := range userActionList {
		if newMessage := userActionElem.GetNewMessage(); newMessage != nil {
			messages = append(messages, utils.MessageInfo{
				Text: newMessage.Text,
				Id:   int(newMessage.Id),
			})
		}
	}

	return messages
}


func (eo *EventObserver) ProcessLockedPart(
	ctx context.Context,
	userDir *meta_group_id.MetaUserId,
	ae *actions.ActionEvent,
) {
	lockVar := userDir.GetLock()
	locked, err := lockVar.TryLock()
	if err != nil {
		panic(err)
	} else if !locked {
		return
	}

	defer func(lock *distvars.ConsulLock) {
		err := lock.Unlock()
		if err != nil {
			fmt.Println(err)
		}
	}(lockVar)

	actionsVar := userDir.CdActions()

	actionStrings, err := actionsVar.Pop(context.Background(), 5*time.Second)
	if err != nil {
		return
	}

	var userActionList []*actions.UserAction
	for _, actionString := range actionStrings {
		var action actions.UserAction
		if err = jsonpb.Unmarshal(bytes.NewReader([]byte(actionString)), &action); err != nil {
			panic(err)
		}
		userActionList = append(userActionList, &action)
	}

	messages := MessageGenerator(userActionList)
	externalStorage := distributedStorage.ConsulStorage{
		Storage:      eo.cm.ConsulClient.KV(),
		DataBasePath: userDir.GetStateMachinePath(),
	}

	//config := configExamples.QuestionsOnlyInternal
	config := configExamples.MarketInternal

	sm := stateMachine.StateMachine{}
	sm.Init(
		eo.cm,
		config.States,
		ae.ServiceName,
		ae.BotId,
		int(ae.ShopId),
		ae.UserId,
		config.Data,
		&externalStorage,
		userDir.GetStateMachinePath(),
	)

	sm.Process(ctx, eo.cm, messages)
	sm.Finish()
}