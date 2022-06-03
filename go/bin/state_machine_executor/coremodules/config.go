package coremodules

import "github.com/nikhovas/diploma/go/lib/utils/foundation/modules/queuewriter"

type Config struct {
	ConsumerMessageSender queuewriter.Config `yaml:"consumer-message-sender"`
	OrdersQueueWriter queuewriter.Config `yaml:"orders-queue-writer"`
	EventsQueueWriter queuewriter.Config `yaml:"events-queue-writer"`
}
