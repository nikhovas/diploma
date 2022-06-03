package queuereader

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/sync/semaphore"
	"log"
	"sync"
)

type QueueReader struct {
	amqpConn *amqp.Connection
	amqpChan *amqp.Channel
	amqpQueue amqp.Queue
	worker    func(msg *amqp.Delivery)
	semaphore *semaphore.Weighted
}

func FromConfig(config Config, worker func(msg *amqp.Delivery)) *QueueReader {
	if config.Url == "" {
		config.Url = "amqp://guest:guest@localhost:5672/"
	}

	qr := &QueueReader{}

	var err error
	qr.amqpConn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalln(err)
	}

	qr.amqpChan, err = qr.amqpConn.Channel()
	if err != nil {
		log.Fatalln(err)
	}

	if config.ExchangeName != "" {
		err = qr.amqpChan.ExchangeDeclare(
			config.ExchangeName, // name
			"fanout",            // type
			true,                // durable
			false,               // auto-deleted
			false,               // internal
			false,               // no-wait
			nil,                 // arguments
		)
		if err != nil {
			log.Fatalln(err)
		}
	}

	q, err := qr.amqpChan.QueueDeclare(
		config.QueueName,    // name
		false, // durable
		false, // delete when unused
		false,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalln(err)
	}

	if config.ExchangeName != "" {
		err = qr.amqpChan.QueueBind(
			q.Name,              // queuereader name
			"",                  // routing key
			config.ExchangeName, // exchange
			false,
			nil,
		)
		if err != nil {
			log.Fatalln(err)
		}
	}

	qr.worker = worker

	if config.MaxWorkers != 0 {
		qr.semaphore = semaphore.NewWeighted(int64(config.MaxWorkers))
	} else {
		qr.semaphore = nil
	}

	return qr
}

func (qr *QueueReader) Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go qr.RunBlock(ctx, wg)
}

func (qr *QueueReader) RunBlock(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	msgs, err := qr.amqpChan.Consume(
		qr.amqpQueue.Name, // queuereader
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {
		case d := <-msgs:
			if qr.semaphore != nil {
				if !qr.semaphore.TryAcquire(1) {
					continue
				}
			}

			go qr.preWorker(&d)


		case <-ctx.Done():
			break
		}
	}
}

func (qr *QueueReader) preWorker(d *amqp.Delivery) {
	qr.worker(d)

	if qr.semaphore != nil {
		qr.semaphore.Release(1)
	}

	if d.Acknowledger != nil {
		if err := d.Ack(false); err != nil {
			fmt.Println(err)
		}
	}
}

func (qr *QueueReader) Stop() {
	err := qr.amqpConn.Close()
	if err != nil {}

	err = qr.amqpChan.Close()
	if err != nil {}
}

