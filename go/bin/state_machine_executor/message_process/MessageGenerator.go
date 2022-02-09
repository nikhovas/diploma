package messageProcess

import (
	"github.com/nikhovas/diploma/go/lib/proto/consumer_actions"
	"state_machine_executor/utils"
)

func MessageGenerator(userActionList []*comsumer_actions.UserAction) []utils.MessageInfo {
	var messages []utils.MessageInfo

	for _, userActionElem := range userActionList {
		if newMessage := userActionElem.GetNewMessage(); newMessage != nil {
			messages = append(messages, utils.MessageInfo{
				Text: newMessage.Text,
				Id:   int(newMessage.Id),
			})
		}
	}

	return messages
}
