package application

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func (a *Application) SetUpAmqp() error {
	const amqpUrl = "amqp://guest:guest@localhost:5672/"
	const queueName = "action_events"

	var ampqConn *amqp.Connection
	var err error

	for i := 0; i < 50; i++ {
		ampqConn, err = amqp.Dial(amqpUrl)
		if err != nil {
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}

	if err != nil {
		return err
	} else {
	}

	a.AmqpInputChannel, err = ampqConn.Channel()
	if err != nil {
		return err
	}
	a.AmqpInputQueue, err = a.AmqpInputChannel.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return err
	}

	return nil
}
