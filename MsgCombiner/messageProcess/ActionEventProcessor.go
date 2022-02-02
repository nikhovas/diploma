package messageProcess

import (
	"MsgCombiner/Application"
	ae "MsgCombiner/grpc/actionEvent"
)

type ActionEventProcessor struct {
	Application *Application.Application
	ActionEvent *ae.ActionEvent
}

const vkShopBot = "vk-shop-bot"
