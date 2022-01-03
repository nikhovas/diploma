package messageProcess

import (
	"MsgCombiner/Application"
	actions "github.com/nikhovas/diploma/proto/data/actionEvent"
)

type ActionEventProcessor struct {
	Application *Application.Application
	ActionEvent *actions.ActionEvent
}

const vkShopBot = "vk-shop-bot"
