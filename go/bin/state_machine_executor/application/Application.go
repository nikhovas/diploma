package application

import (
	"context"
	"github.com/go-redis/redis/v8"
	consulApi "github.com/hashicorp/consul/api"
	consumerBot "github.com/nikhovas/diploma/go/lib/proto/consumer_bot"
	"github.com/nikhovas/diploma/go/lib/proto/controller"
	questionWorker "github.com/nikhovas/diploma/go/lib/proto/question_worker"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/sync/semaphore"
	"sync"
)

type Application struct {
	ConsulClient     *consulApi.Client
	RedisClient      *redis.Client
	AmqpInputChannel *amqp.Channel
	AmqpInputQueue   amqp.Queue
	Semaphore        *semaphore.Weighted
	Context          context.Context
	ReadQueueWg      sync.WaitGroup

	ControlClient controller.ControllerClient
	QwClient      questionWorker.QuestionWorkerClient
	VksClient     consumerBot.VkServerClient
}

func (a *Application) Init() {
	a.Semaphore = semaphore.NewWeighted(int64(2000))
	a.Context = context.TODO()
}
