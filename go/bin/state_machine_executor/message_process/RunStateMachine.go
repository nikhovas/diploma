package messageProcess

import (
	"context"
	"fmt"
	"state_machine_executor/application"
	"state_machine_executor/config_examples"
	"state_machine_executor/state_machine"
	"state_machine_executor/state_machine/distributedStorage"
)

func (aep *ActionEventProcessor) RunStateMachine(app *application.Application, messages []string) {
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
	ctx := context.Background()

	sm.Process(ctx, app, messages)
	sm.Finish()
}
