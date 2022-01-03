package messageProcess

import UserActions "github.com/nikhovas/diploma/proto/data/userActions"

func MessageGenerator(userActionList []*UserActions.UserAction) ([]string, error) {
	var messages []string

	// TODO: advavnced message generator

	for _, userActionElem := range userActionList {
		if newMessage := userActionElem.GetNewMessage(); newMessage != nil {
			messages = append(messages, newMessage.Text)
		}
	}

	return messages, nil
}
