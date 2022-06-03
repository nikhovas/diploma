package stateMachine

import (
	"context"
	"fmt"
	"state_machine_executor/coremodules"
	"state_machine_executor/state_machine/localStorage"
	"state_machine_executor/state_machine/states"
	"state_machine_executor/state_machine/storageInterfaces"
	"state_machine_executor/utils"
	"strconv"
)

type StateMachine struct {
	cm *coremodules.CoreModules
	states          map[string]states.IState
	dataBasePath    string
	storage         localStorage.Storage
	externalStorage storageInterfaces.IStorage
}

func (sm *StateMachine) Init(
	cm *coremodules.CoreModules,
	states map[string]states.IState,
	botService string,
	groupId string,
	shopId int,
	userId string,
	storageDescription map[string]localStorage.DataElement,
	externalStorage storageInterfaces.IStorage,
	userIdStateMachinePath string,
) {
	sm.cm = cm
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

func (sm *StateMachine) Process(ctx context.Context, cm *coremodules.CoreModules, newMessages []utils.MessageInfo) {
	for _, msg := range newMessages {
		sm.storage.MessageDeque.PushBack(&msg)
	}

	currentState := sm.storage.KvStorage.Get("state")
	breakContext := false

	for !breakContext {
		state, exists := sm.states[currentState.(string)]
		if !exists {
			state = sm.states["initial"]
			fmt.Printf("No state with name %s, using inital state", currentState.(string))
		}

		fmt.Printf("Executing state: %s", currentState)
		currentState, breakContext = state.Process(ctx, cm, &sm.storage)
		sm.storage.KvStorage.Set("state", currentState)
		fmt.Printf("New state: %s", currentState)
	}
}
