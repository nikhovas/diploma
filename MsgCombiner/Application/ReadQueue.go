package Application

import (
	"MsgCombiner/messageProcess"
	"bytes"
	"github.com/golang/protobuf/jsonpb"
	actions "github.com/nikhovas/diploma/proto/data/actionEvent"
)

func (a *Application) ReadQueue() error {
	messages, _ := a.AmqpInputChannel.Consume(a.AmqpInputQueue.Name, "", true, false, false, false, nil)

	for {
		if err := a.Semaphore.Acquire(a.Context, 1); err != nil {
			return err
		}

		select {
		case d := <-messages:
			var ae actions.ActionEvent
			err := jsonpb.Unmarshal(bytes.NewReader(d.Body), &ae)
			if err != nil {
				continue
			}

			aep := &messageProcess.ActionEventProcessor{
				Application: a,
				ActionEvent: &ae,
			}
			go func() {
				if aep.Process(&a.ReadQueueWg) != nil {
				}
			}()
		}
	}
}
