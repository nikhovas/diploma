package messageProcess

import (
	"MsgCombiner/configExamples"
	"MsgCombiner/stateMachine"
	"MsgCombiner/stateMachine/distributedStorage"
	"fmt"
)

func (aep *ActionEventProcessor) RunStateMachine(messages []string) {
	externalStorage := distributedStorage.ConsulStorage{
		Storage: aep.Application.ConsulClient.KV(),
		DataBasePath: fmt.Sprintf(
			"messages/%s/%s/%s/data",
			aep.ActionEvent.ServiceName,
			aep.ActionEvent.BotId,
			aep.ActionEvent.UserId,
		),
	}

	sm := stateMachine.StateMachine{}
	sm.Init(
		aep.Application,
		configExamples.QuestionsOnlyInternal.States,
		aep.ActionEvent.ServiceName,
		aep.ActionEvent.BotId,
		int(aep.ActionEvent.ShopId),
		aep.ActionEvent.UserId,
		configExamples.QuestionsOnlyInternal.Data,
		&externalStorage,
	)
	sm.Process(messages)
	sm.Finish()
}
