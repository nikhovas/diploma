package main

import (
	"MsgCombiner/Application"
	"MsgCombiner/messageProcess"
	"bytes"
	"github.com/golang/protobuf/jsonpb"
	actions "github.com/nikhovas/diploma/proto/data/actionEvent"
	"log"
)

func ReadQueue(a *Application.Application) error {
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
			a.ReadQueueWg.Add(1)
			go func() {
				if aep.Process(&a.ReadQueueWg) != nil {
				}
			}()
		}
	}
}

func main() {
	var app Application.Application

	app.Init()

	if err := app.SetUpCoordinator(); err != nil {
		log.Fatal(err)
	}

	if err := app.SetUpAmqp(); err != nil {
		log.Fatal(err)
	}

	if err := ReadQueue(&app); err != nil {
		log.Fatal(err)
	}
}
