package main

import (
	"github.com/nikhovas/diploma/go/lib/utils/foundation"
	"vk_consumer_bot/application"
)

func main() {
	var config application.Config
	foundation.GetConfig(&config)
	app := application.FromConfig(config)
	foundation.RunBlock(app)
}
