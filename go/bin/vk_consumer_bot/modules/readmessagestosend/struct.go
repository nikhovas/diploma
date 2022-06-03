package readmessagestosend

import (
	"bytes"
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/nikhovas/diploma/go/lib/proto/consumer_messages"
	"github.com/nikhovas/diploma/go/lib/utils/foundation/modules/queuereader"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
	"vk_consumer_bot/modules/combinedbot"
)

type ReadMessagesToSend struct {
	queueReader *queuereader.QueueReader
	combinedBot *combinedbot.CombinedBot
}

func FromConfig(config Config, combinedBot *combinedbot.CombinedBot) *ReadMessagesToSend {
	module := &ReadMessagesToSend{}
	module.queueReader = queuereader.FromConfig(config.QueueReader, module.worker)
	module.combinedBot = combinedBot
	return module
}

func (module *ReadMessagesToSend) Run(ctx context.Context, wg *sync.WaitGroup) {
	module.queueReader.Run(ctx, wg)
}

func (module *ReadMessagesToSend) worker(d *amqp.Delivery) {
	var msg consumer_messages.MessageToConsumer
	err := jsonpb.Unmarshal(bytes.NewReader(d.Body), &msg)
	if err != nil {
		log.Fatal(err)
	}

	var replyTo *int
	replyTo = nil

	if reply := msg.Reply; reply != nil {
		data := int(reply.ReplyMessageId)
		replyTo = &data
	}

	_, err = module.combinedBot.SendMessage(
		int(msg.MsgLocation.GroupId),
		int(msg.MsgLocation.UserId),
		msg.Text,
		replyTo,
	)
}

func (module *ReadMessagesToSend) Stop() {
	module.queueReader.Stop()
}
