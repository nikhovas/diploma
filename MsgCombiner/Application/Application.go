package Application

import (
	ctrl "MsgCombiner/grpc/control"
	qw "MsgCombiner/grpc/questionWorker"
	"context"
	consulApi "github.com/hashicorp/consul/api"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/sync/semaphore"
	"sync"
)

type Application struct {
	ConsulClient     *consulApi.Client
	AmqpInputChannel *amqp.Channel
	AmqpInputQueue   amqp.Queue
	Semaphore        *semaphore.Weighted
	Context          context.Context
	ReadQueueWg      sync.WaitGroup

	ControlClient ctrl.ControlClient
	QwClient      qw.QuestionWorkerClient
}

func (a *Application) Init() {
	a.Semaphore = semaphore.NewWeighted(int64(2000))
	a.Context = context.TODO()
}
