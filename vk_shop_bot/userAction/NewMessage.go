package userAction

const NewMessageActionName = "new_message"

type NewMessage struct {
	text string `json:"text"`
}
