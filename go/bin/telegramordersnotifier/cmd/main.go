package main

import (
	"github.com/nikhovas/diploma/go/lib/utils/foundation"
	"telegramordersnotifier/application"
)

func main() {
	var config application.Config
	foundation.GetConfig(&config)
	app := application.FromConfig(config)
	foundation.RunBlock(app)
}
