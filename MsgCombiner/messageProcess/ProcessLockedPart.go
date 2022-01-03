package messageProcess

import (
	"MsgCombiner/utils"
	"github.com/hashicorp/consul/api"
)

func (aep *ActionEventProcessor) ProcessLockedPart(messagesKey string) (int, error) {
	lockKey := utils.GetLockKey(vkShopBot, aep.ActionEvent.BotId, aep.ActionEvent.UserId)

	locked := true

	lock, err := aep.Application.ConsulClient.LockKey(lockKey)
	if err != nil {
		return 0, err
	}

	if !locked {
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

	// main part: call state machine
	//stateManager(messages, vkShopBot, ae.GetBotId(), ae.GetUserId())

	return newestTs, nil
}
