package messageProcess

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	actions "github.com/nikhovas/diploma/go/lib/proto/consumer_actions"
	"github.com/nikhovas/diploma/go/lib/utils/distfs"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"state_machine_executor/application"
	"sync"
)

func Process(ctx context.Context, wg *sync.WaitGroup, app *application.Application, ae *actions.ActionEvent) {
	defer wg.Done()
	userDir := distfs.NewRoot(app.RedisClient, app.ConsulClient).CdBots().MetaCdServiceName(ae.ServiceName).
		MetaCdGroupId(ae.BotId).MetaCdUserId(ae.UserId)

	ProcessLockedPart(ctx, app, userDir, ae)

	// send to self if any values exists
	actionsStartFromVar := userDir.CdActions()
	value, err := actionsStartFromVar.GetNewestTimeValue(ctx)
	if err != nil {
		return
	}
	if value == nil {
		return
	}

	ae.Time = uint64(*value)

	m := jsonpb.Marshaler{}
	aeString, _ := m.MarshalToString(ae)
	err = app.AmqpInputChannel.Publish(
		"",
		app.AmqpInputQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(aeString),
		},
	)
	if err != nil {
		log.Println(err)
	}
}
