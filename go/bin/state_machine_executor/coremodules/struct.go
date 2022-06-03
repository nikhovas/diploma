package coremodules

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/nikhovas/diploma/go/lib/proto/controller"
	questionWorker "github.com/nikhovas/diploma/go/lib/proto/question_worker"
	"github.com/nikhovas/diploma/go/lib/utils/clients"
	"golang.org/x/sync/semaphore"
	"google.golang.org/grpc"
	"state_machine_executor/modules/consumermessagesender"
	"state_machine_executor/modules/eventsqueuewriter"
	"state_machine_executor/modules/ordersqueuewriter"
	"sync"
)

type CoreModules struct {
	ConsulClient     *consulApi.Client
	RedisClient      *redis.Client
	Semaphore        *semaphore.Weighted
	ReadQueueWg      sync.WaitGroup

	ControlClient controller.ControllerClient
	QwClient      questionWorker.QuestionWorkerClient

	ConsumerMessageSender *consumermessagesender.ConsumerMessageSender
	OrdersQueueWriter     *ordersqueuewriter.OrdersQueueWriter
	EventsQueueWriter     *eventsqueuewriter.EventsQueueWriter

	ctrlConn *grpc.ClientConn
	qwConn *grpc.ClientConn
}

func FromConfig(config Config) *CoreModules {
	ctrlConn, ctrlClient := clients.CreateControllerClient()
	qwConn, qwClient := clients.CreateQuestionWorkerClient()

	consumerMessageSender := consumermessagesender.FromConfig(config.ConsumerMessageSender)
	ordersQueueWriter := ordersqueuewriter.FromConfig(config.OrdersQueueWriter)
	eventsQueueWriter := eventsqueuewriter.FromConfig(config.EventsQueueWriter)

	return &CoreModules{
		ConsulClient:  clients.CreateConsulClient(),
		RedisClient:   clients.CreateRedisClient(),
		Semaphore:     semaphore.NewWeighted(int64(2000)),
		ReadQueueWg:   sync.WaitGroup{},
		ControlClient: ctrlClient,
		QwClient:      qwClient,

		ConsumerMessageSender: consumerMessageSender,
		OrdersQueueWriter: ordersQueueWriter,
		EventsQueueWriter: eventsQueueWriter,

		ctrlConn: ctrlConn,
		qwConn: qwConn,
	}
}

func (cm *CoreModules) Run(ctx context.Context, wg *sync.WaitGroup) {
	cm.ConsumerMessageSender.Run(ctx, wg)
	cm.OrdersQueueWriter.Run(ctx, wg)
	cm.EventsQueueWriter.Run(ctx, wg)
}

func (cm *CoreModules) Stop() {
	err := cm.ctrlConn.Close()
	if err != nil {
		fmt.Println(err)
	}

	err = cm.qwConn.Close()
	if err != nil {
		fmt.Println(err)
	}

	cm.ConsumerMessageSender.Stop()
	cm.OrdersQueueWriter.Stop()
	cm.EventsQueueWriter.Stop()
}