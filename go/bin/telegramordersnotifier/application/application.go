package application

import (
	"context"
	"sync"
	"telegramordersnotifier/modules/databasesync"
	"telegramordersnotifier/modules/queuereader"
	"telegramordersnotifier/modules/telegramsender"
)

type Application struct {
	DatabaseSync *databasesync.DatabaseSync
	QueueReader  *queuereader.QueueReader
	TelegramSender *telegramsender.TelegramSender
}

func FromConfig(config Config) *Application {
	databaseSync := databasesync.FromConfig(config.DatabaseSync)
	telegramSender := telegramsender.FromConfig(config.TelegramSender)

	return &Application{
		DatabaseSync: databaseSync,
		QueueReader:  queuereader.FromConfig(config.QueueReader, databaseSync, telegramSender),
		TelegramSender: telegramSender,
	}
}

func (app *Application) Run(ctx context.Context, wg *sync.WaitGroup) {
	app.TelegramSender.Run(ctx, wg)
	app.DatabaseSync.Run(ctx, wg)
	app.QueueReader.Run(ctx, wg)
}

func (app *Application) Stop() {
	app.TelegramSender.Stop()
	app.QueueReader.Stop()
	app.DatabaseSync.Stop()
}
