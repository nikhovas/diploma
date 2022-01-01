package userAction

type UserAction struct {
	ActionType string      `json:"action_type"`
	Object     interface{} `json:"object"`
}
