package application

import (
	"github.com/nikhovas/diploma/go/lib/utils/foundation/modules/queuewriter"
	"state_machine_executor/modules/eventobserver"
	"state_machine_executor/modules/eventsqueuereader"
)

type Config struct {
	ConsumerMessageSender queuewriter.Config `yaml:"consumer-message-sender"`
	OrdersQueueWriter queuewriter.Config `yaml:"orders-queue-writer"`
	EventsQueueWriter queuewriter.Config `yaml:"events-queue-writer"`
	EventsQueueReader eventsqueuereader.Config `yaml:"events-queue-reader"`
	EventObserver eventobserver.Config `yaml:"event-observer"`
}
