package distQueue

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"reflect"
	"time"
)

type DistQueue struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	queue     amqp.Queue
	marshaler jsonpb.Marshaler
	elemType  reflect.Type
}

func NewDistQueue(url string, queue string, elemType reflect.Type) (*DistQueue, error) {
	dq := DistQueue{}
	var err error

	dq.conn, err = amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	dq.channel, err = dq.conn.Channel()
	if err != nil {
		return nil, err
	}
	dq.queue, err = dq.channel.QueueDeclare(queue, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	dq.marshaler = jsonpb.Marshaler{}
	dq.elemType = elemType

	return &dq, nil
}

type DistQueueConfig struct {
	Url  string `yaml:"url"`
	Name string `yaml:"name"`
}

func NewDistQueueFromCfg(config DistQueueConfig, elemType reflect.Type) (*DistQueue, error) {
	return NewDistQueue(config.Url, config.Name, elemType)
}

func (dq *DistQueue) Push(msg proto.Message, delay int) error {
	res, err := dq.marshaler.MarshalToString(msg)
	if err != nil {
		return err
	}

	headers := make(amqp.Table)

	if delay != 0 {
		headers["x-delay"] = delay
	}

	err = dq.channel.Publish(
		"",
		dq.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         []byte(res),
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
			Headers:      headers,
		},
	)
	return err
}

func (dq *DistQueue) GetChan() (chan proto.Message, error) {
	deliveries := make(chan proto.Message)

	messages, err := dq.channel.Consume(
		dq.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for msg := range messages {
			val1 := reflect.New(dq.elemType)
			val2 := val1.Interface()
			val := val2.(proto.Message)
			d := string([]byte(msg.Body))
			fmt.Print(d)
			err := jsonpb.Unmarshal(bytes.NewReader([]byte(msg.Body)), val)
			if err != nil {
				continue
			}

			deliveries <- val
		}
		close(deliveries)
	}()
	return deliveries, err
}
