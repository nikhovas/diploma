package consumermessagesender

import "github.com/nikhovas/diploma/go/lib/utils/foundation/modules/queuewriter"

type ConsumerMessageSender = queuewriter.QueueWriter

func FromConfig(config queuewriter.Config) *ConsumerMessageSender {
	return queuewriter.FromConfig(config)
}
