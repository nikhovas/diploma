package application

import (
	"context"
	"github.com/nikhovas/diploma/go/lib/utils/foundation"
	"state_machine_executor/coremodules"
	"state_machine_executor/modules/eventobserver"
	"state_machine_executor/modules/eventsqueuereader"
	"sync"
)

type Application struct {
	CoreModules      *coremodules.CoreModules
	EventsQueueReader     foundation.Module
	EventObserver         *eventobserver.EventObserver
}

func FromConfig(config Config) *Application {
	coreModules := coremodules.FromConfig(coremodules.Config{
		ConsumerMessageSender: config.ConsumerMessageSender,
		OrdersQueueWriter:     config.OrdersQueueWriter,
		EventsQueueWriter:     config.EventsQueueWriter,
	})
	eventObserver := eventobserver.FromConfig(config.EventObserver, coreModules)

	return &Application{
		CoreModules: coreModules      ,
		EventsQueueReader: eventsqueuereader.FromConfig(
			config.EventsQueueReader,
			&coreModules.ReadQueueWg,
			coreModules,
			eventObserver,
		),
		EventObserver:     eventObserver,
	}
}

func (app *Application) Run(ctx context.Context, wg *sync.WaitGroup) {
	app.CoreModules.Run(ctx, wg)
	app.EventsQueueReader.Run(ctx, wg)
	app.EventObserver.Run(ctx, wg)
}

func (app *Application) Stop() {
	app.CoreModules.Stop()
	app.EventsQueueReader.Stop()
	app.EventObserver.Stop()
}