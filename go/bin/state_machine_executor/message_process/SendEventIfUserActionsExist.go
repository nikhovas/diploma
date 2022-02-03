package messageProcess

import (
	"bytes"
	"github.com/golang/protobuf/jsonpb"
	UserActions "github.com/nikhovas/diploma/proto/data/userActions"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func (aep *ActionEventProcessor) SendEventIfUserActionsExist(messagesKey string, newestTs int) {
	list, _, err := aep.Application.ConsulClient.KV().List(messagesKey, nil)
	if err != nil {
		return
	}

	sendNewEvent := false

	for _, kv := range list {
		var userAction UserActions.UserAction
		if err = jsonpb.Unmarshal(bytes.NewReader(kv.Value), &userAction); err != nil {
			return
		}
		ts := int(userAction.Time)

		if ts <= newestTs {
			_, _ = aep.Application.ConsulClient.KV().Delete(kv.Key, nil)
		} else {
			sendNewEvent = true
		}
	}

	if sendNewEvent {
		aep.ActionEvent.Time = uint64(newestTs)
		m := jsonpb.Marshaler{}
		aeString, _ := m.MarshalToString(aep.ActionEvent)
		err := aep.Application.AmqpInputChannel.Publish(
			"",
			aep.Application.AmqpInputQueue.Name,
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
}
