package messageProcess

import (
	"MsgCombiner/utils"
	"fmt"
	"github.com/hashicorp/consul/api"
)

func (aep *ActionEventProcessor) ProcessLockedPart(messagesKey string) (int, error) {
	lockKey := utils.GetLockKey(vkShopBot, aep.ActionEvent.BotId, aep.ActionEvent.UserId)

	lock, err := aep.Application.ConsulClient.LockOpts(&api.LockOptions{
		Key:         lockKey,
		LockTryOnce: true,
	})
	if err != nil {
		return 0, err
	}

	ch, err := lock.Lock(nil)
	if err != nil {
		return 0, err
	}

	// no lock
	if ch == nil {
		return 0, nil
	}

	defer func(lock *api.Lock) {
		err := lock.Unlock()
		if err != nil {
			// logging
		}
	}(lock)

	newestTs, userActionList, err := aep.GetAllUserActions(messagesKey)
	if err != nil {
		return 0, err
	}

	// TODO: neuroprocess messages

	messages, _ := MessageGenerator(userActionList)

	for _, msg := range messages {
		fmt.Println(msg)
	}

	// main part: call state machine
	//stateManager(messages, vkShopBot, ae.GetBotId(), ae.GetUserId())

	return newestTs, nil
}
