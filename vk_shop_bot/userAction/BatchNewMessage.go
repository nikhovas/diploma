package userAction

const BatchNewMessageActionName = "new_message_batch"

type BatchNewMessage struct {
	Messages []NewMessage `json:"messages"`
}
