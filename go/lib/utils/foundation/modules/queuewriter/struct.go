package queuewriter

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

type QueueWriter struct {
	amqpConn *amqp.Connection
	amqpChan *amqp.Channel
	amqpQueue *amqp.Queue
	exchangeName string
	queueName string
}

func FromConfig(config Config) *QueueWriter {
	if config.Url == "" {
		config.Url = "amqp://guest:guest@localhost:5672/"
	}

	qw := &QueueWriter{}

	var err error

	qw.amqpConn, err = amqp.Dial(config.Url)
	if err != nil {
		log.Fatal(err)
	}

	//for i := 0; i < 50; i++ {
	//	qw.amqpConn, err = amqp.Dial(config.Url)
	//	if err != nil {
	//		time.Sleep(10 * time.Second)
	//	} else {
	//		break
	//	}
	//}
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	qw.amqpChan, err = qw.amqpConn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	qw.queueName = ""
	if queueName := config.QueueName; queueName != "" {
		queue, err := qw.amqpChan.QueueDeclare(queueName, false, false, false, false, nil)
		if err != nil {
			log.Fatal(err)
		}
		qw.amqpQueue = &queue
		qw.queueName = queue.Name
	}

	if exchangeName := config.ExchangeName; exchangeName != "" {
		err := qw.amqpChan.ExchangeDeclare(
			config.ExchangeName,   // name
			"fanout", // type
			true,     // durable
			false,    // auto-deleted
			false,    // internal
			false,    // no-wait
			nil,      // arguments
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	qw.exchangeName = config.ExchangeName

	return qw
}

func (qw *QueueWriter) Run(ctx context.Context, wg *sync.WaitGroup) {}

func (qw *QueueWriter) Stop() {}

func (qw *QueueWriter) SendMessage(msg proto.Message) error {
	m := jsonpb.Marshaler{}
	stringMsg, _ := m.MarshalToString(msg)
	publishing := amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(stringMsg),
	}

	err := qw.amqpChan.Publish(qw.exchangeName, qw.queueName, false, false, publishing)
	if err != nil {
		return err
	}
	return nil
}