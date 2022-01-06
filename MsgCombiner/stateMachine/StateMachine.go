package stateMachine

import (
	"MsgCombiner/Application"
	"MsgCombiner/stateMachine/localStorage"
	"MsgCombiner/stateMachine/states"
	"MsgCombiner/stateMachine/storageInterfaces"
	"MsgCombiner/utils"
	"fmt"
)

type StateMachine struct {
	application     *Application.Application
	states          map[string]states.IState
	dataBasePath    string
	storage         localStorage.Storage
	externalStorage storageInterfaces.IStorage
}

func (sm *StateMachine) Init(
	application *Application.Application,
	states map[string]states.IState,
	botService string,
	groupId string,
	userId string,
	storageDescription map[string]localStorage.DataElement,
	externalStorage storageInterfaces.IStorage,
) {
	sm.application = application
	sm.states = states

	sm.dataBasePath = fmt.Sprintf("%s/state-machine", utils.GetDataBasePath(botService, groupId, userId))

	storageDescription["botService"] = localStorage.DataElement{Memory: "const", Type: "string", Default: botService}
	storageDescription["groupId"] = localStorage.DataElement{Memory: "const", Type: "string", Default: groupId}
	storageDescription["userId"] = localStorage.DataElement{Memory: "const", Type: "string", Default: userId}
	storageDescription["state"] = localStorage.DataElement{Memory: "long", Type: "string", Default: "initial"}
	storageDescription["message"] = localStorage.DataElement{Memory: "short", Type: "string", Default: ""}
	sm.storage.KvStorage.Init(storageDescription, externalStorage)

	sm.externalStorage = externalStorage
}

func (sm *StateMachine) GetDataBasePath() string {
	return sm.dataBasePath
}

func (sm *StateMachine) Finish() {
	longMemory := sm.storage.KvStorage.GetUpdatedLongMemory()
	for key, value := range longMemory {
		sm.externalStorage.Set(key, value)
	}
}

func (sm *StateMachine) Process(newMessages []string) {
	for _, msg := range newMessages {
		sm.storage.MessageDeque.PushBack(msg)
	}

	currentState := sm.storage.KvStorage.Get("state")
	breakContext := false

	for !breakContext {
		currentState, breakContext = sm.states[currentState.(string)].Process(&sm.storage)
		sm.storage.KvStorage.Set("state", currentState)
	}
}
