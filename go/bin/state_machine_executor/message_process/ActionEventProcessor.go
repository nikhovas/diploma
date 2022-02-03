package messageProcess

import (
	actions "github.com/nikhovas/diploma/go/lib/proto/consumer_actions"
	"state_machine_executor/application"
)

type ActionEventProcessor struct {
	Application *application.Application
	ActionEvent *actions.ActionEvent
}

const vkShopBot = "vk-shop-bot"
