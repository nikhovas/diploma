package messageProcess

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"state_machine_executor/utils"
)

func (aep *ActionEventProcessor) ProcessLockedPart(messagesKey string) int {
	lockKey := utils.GetLockKey(vkShopBot, aep.ActionEvent.BotId, aep.ActionEvent.UserId)

	lock, err := aep.Application.ConsulClient.LockOpts(&api.LockOptions{
		Key:         lockKey,
		LockTryOnce: true,
	})
	if err != nil {
		panic(err)
	}

	ch, err := lock.Lock(nil)
	if err != nil {
		panic(err)
	}

	// no lock
	if ch == nil {
		return 0
	}

	defer func(lock *api.Lock) {
		err := lock.Unlock()
		if err != nil {
			// logging
		}
	}(lock)

	newestTs, userActionList, err := aep.GetAllUserActions(messagesKey)
	if err != nil {
		panic(err)
	}

	messages := MessageGenerator(userActionList)

	for _, msg := range messages {
		fmt.Println(msg)
	}

	aep.RunStateMachine(aep.Application, messages)

	return newestTs
}
