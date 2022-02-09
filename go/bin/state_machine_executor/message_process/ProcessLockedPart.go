package messageProcess

import (
	"bytes"
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	actions "github.com/nikhovas/diploma/go/lib/proto/consumer_actions"
	"github.com/nikhovas/diploma/go/lib/utils/distfs/bots/meta_service_name/meta_group_id"
	"github.com/nikhovas/diploma/go/lib/utils/distvars"
	"state_machine_executor/application"
	configExamples "state_machine_executor/config_examples"
	stateMachine "state_machine_executor/state_machine"
	"state_machine_executor/state_machine/distributedStorage"
	"time"
)

func ProcessLockedPart(
	ctx context.Context,
	app *application.Application,
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
		Storage:      app.ConsulClient.KV(),
		DataBasePath: userDir.GetStateMachinePath(),
	}

	sm := stateMachine.StateMachine{}
	sm.Init(
		app,
		configExamples.QuestionsOnlyInternal.States,
		ae.ServiceName,
		ae.BotId,
		int(ae.ShopId),
		ae.UserId,
		configExamples.QuestionsOnlyInternal.Data,
		&externalStorage,
		userDir.GetStateMachinePath(),
	)

	sm.Process(ctx, app, messages)
	sm.Finish()
}
