package main

import (
	"bytes"
	"github.com/golang/protobuf/jsonpb"
	actions "github.com/nikhovas/diploma/go/lib/proto/consumer_actions"
	cb "github.com/nikhovas/diploma/go/lib/proto/consumer_bot"
	ctrl "github.com/nikhovas/diploma/go/lib/proto/controller"
	qw "github.com/nikhovas/diploma/go/lib/proto/question_worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

			aep := &messageProcess.ActionEventProcessor{
				Application: a,
				ActionEvent: &ae,
			}
			a.ReadQueueWg.Add(1)
			go func() {
				aep.Process(&a.ReadQueueWg)
			}()
		}
	}
}

func createControllerClient() (*grpc.ClientConn, ctrl.ControllerClient) {
	conn, err := grpc.Dial("localhost:7777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := ctrl.NewControllerClient(conn)
	return conn, client
}

func createQwClient() (*grpc.ClientConn, qw.QuestionWorkerClient) {
	conn, err := grpc.Dial("localhost:7777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := qw.NewQuestionWorkerClient(conn)
	return conn, client
}

func createVkServerClient() (*grpc.ClientConn, cb.VkServerClient) {
	conn, err := grpc.Dial("localhost:5555", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := cb.NewVkServerClient(conn)
	return conn, client
}

func main() {
	var app application.Application

	ctrlConn, ctrlClient := createControllerClient()
	defer ctrlConn.Close()
	vksConn, vksClient := createVkServerClient()
	defer vksConn.Close()
	qwConn, qwClient := createQwClient()
	defer qwConn.Close()

	app.VksClient = vksClient
	app.ControlClient = ctrlClient
	app.QwClient = qwClient

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
