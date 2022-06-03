package ordersqueuewriter

import "github.com/nikhovas/diploma/go/lib/utils/foundation/modules/queuewriter"

type OrdersQueueWriter = queuewriter.QueueWriter

func FromConfig(config queuewriter.Config) *OrdersQueueWriter {
	return queuewriter.FromConfig(config)
}
