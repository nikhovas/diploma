package stateMachine

import (
	"context"
	"state_machine_executor/application"
	"state_machine_executor/state_machine/localStorage"
	"state_machine_executor/state_machine/states"
	"state_machine_executor/state_machine/storageInterfaces"
	"state_machine_executor/utils"
	"strconv"
)

type StateMachine struct {
	application     *application.Application
	states          map[string]states.IState
	dataBasePath    string
	storage         localStorage.Storage
	externalStorage storageInterfaces.IStorage
}

func (sm *StateMachine) Init(
	application *application.Application,
	states map[string]states.IState,
	botService string,
	groupId string,
	shopId int,
	userId string,
	storageDescription map[string]localStorage.DataElement,
	externalStorage storageInterfaces.IStorage,
	userIdStateMachinePath string,
) {
	sm.application = application
	sm.states = states

	sm.dataBasePath = userIdStateMachinePath

	storageDescription["botService"] = localStorage.DataElement{Memory: "const", Type: "string", Default: botService}
	storageDescription["shopId"] = localStorage.DataElement{Memory: "const", Type: "string", Default: strconv.Itoa(shopId)}
	storageDescription["userId"] = localStorage.DataElement{Memory: "const", Type: "string", Default: userId}
	storageDescription["state"] = localStorage.DataElement{Memory: "long", Type: "string", Default: "initial"}
	storageDescription["message"] = localStorage.DataElement{Memory: "short", Type: "string", Default: ""}
	storageDescription["messageId"] = localStorage.DataElement{Memory: "short", Type: "string", Default: ""}
	storageDescription["groupId"] = localStorage.DataElement{Memory: "const", Type: "string", Default: groupId}
	sm.storage.KvStorage.Init(storageDescription, externalStorage)

	sm.externalStorage = externalStorage
}

func (sm *StateMachine) Finish() {
	longMemory := sm.storage.KvStorage.GetUpdatedLongMemory()
	for key, value := range longMemory {
		sm.externalStorage.Set(key, value)
	}
}

func (sm *StateMachine) Process(ctx context.Context, app *application.Application, newMessages []utils.MessageInfo) {
	for _, msg := range newMessages {
		sm.storage.MessageDeque.PushBack(&msg)
	}

	currentState := sm.storage.KvStorage.Get("state")
	breakContext := false

	for !breakContext {
		currentState, breakContext = sm.states[currentState.(string)].Process(ctx, app, &sm.storage)
		sm.storage.KvStorage.Set("state", currentState)
	}
}
