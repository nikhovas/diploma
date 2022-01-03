package messageProcess

import (
	"MsgCombiner/utils"
	"sync"
)

func (aep *ActionEventProcessor) Process(wg *sync.WaitGroup) error {
	defer wg.Done()

	messagesKey := utils.GetMessagesKey(vkShopBot, aep.ActionEvent.BotId, aep.ActionEvent.UserId)

	newestTs, err := aep.ProcessLockedPart(messagesKey)
	if err != nil {
		return err
	}

	aep.SendEventIfUserActionsExist(messagesKey, newestTs)

	return nil
}
