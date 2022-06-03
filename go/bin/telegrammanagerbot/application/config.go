package application

import (
	"telegramordersnotifier/modules/databasesync"
	"telegramordersnotifier/modules/queuereader"
	"telegramordersnotifier/modules/telegramsender"
)

type Config struct {
	DatabaseSync databasesync.Config `yaml:"database-sync"`
	QueueReader  queuereader.Config `yaml:"queue-reader"`
	TelegramSender telegramsender.Config `yaml:"telegram-sender"`
}
