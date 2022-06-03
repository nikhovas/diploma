package eventsqueuewriter

import "github.com/nikhovas/diploma/go/lib/utils/foundation/modules/queuewriter"

type EventsQueueWriter = queuewriter.QueueWriter

func FromConfig(config queuewriter.Config) *EventsQueueWriter {
	return queuewriter.FromConfig(config)
}