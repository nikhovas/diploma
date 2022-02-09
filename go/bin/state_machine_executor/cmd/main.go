package main

import (
	"bytes"
	"context"
	"github.com/golang/protobuf/jsonpb"
	actions "github.com/nikhovas/diploma/go/lib/proto/consumer_actions"
	"github.com/nikhovas/diploma/go/lib/utils/clients"
	"log"
	"state_machine_executor/application"
	messageProcess "state_machine_executor/message_process"
)

func ReadQueue(a *application.Application) error {
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

			a.ReadQueueWg.Add(1)
			go func() {
				messageProcess.Process(context.Background(), &a.ReadQueueWg, a, &ae)
			}()
		}
	}
}

func main() {
	var app application.Application

	ctrlConn, ctrlClient := clients.CreateControllerClient()
	defer ctrlConn.Close()
	vksConn, vksClient := clients.CreateConsumerBotClient()
	defer vksConn.Close()
	qwConn, qwClient := clients.CreateQuestionWorkerClient()
	defer qwConn.Close()

	app.VksClient = vksClient
	app.ControlClient = ctrlClient
	app.QwClient = qwClient
	app.ConsulClient = clients.CreateConsulClient()
	app.RedisClient = clients.CreateRedisClient()

	app.Init()

	if err := app.SetUpAmqp(); err != nil {
		log.Fatal(err)
	}

	if err := ReadQueue(&app); err != nil {
		log.Fatal(err)
	}
}
