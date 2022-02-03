package messageProcess

import (
	"state_machine_executor/utils"
	"sync"
)

func (aep *ActionEventProcessor) Process(wg *sync.WaitGroup) {
	defer wg.Done()
	messagesKey := utils.GetMessagesKey(vkShopBot, aep.ActionEvent.BotId, aep.ActionEvent.UserId)
	newestTs := aep.ProcessLockedPart(messagesKey)
	aep.SendEventIfUserActionsExist(messagesKey, newestTs)
}
