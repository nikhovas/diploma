package queuereader

import (
	"bytes"
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/nikhovas/diploma/go/lib/proto/common"
	"github.com/nikhovas/diploma/go/lib/utils/foundation/modules/queuereader"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
	"telegramordersnotifier/modules/databasesync"
	"telegramordersnotifier/modules/telegramsender"
)

type QueueReader struct {
	queueReader *queuereader.QueueReader
	databaseSync *databasesync.DatabaseSync
	telegramSender *telegramsender.TelegramSender
}

func FromConfig(config Config, databaseSync *databasesync.DatabaseSync, telegramSender *telegramsender.TelegramSender) *QueueReader {
	qr := &QueueReader{
		databaseSync: databaseSync,
		telegramSender: telegramSender,
	}
	qr.queueReader = queuereader.FromConfig(config.QueueReader, qr.worker)
	return qr
}

func (qr *QueueReader) Run(ctx context.Context, wg *sync.WaitGroup) {
	qr.queueReader.Run(ctx, wg)
}

func (qr *QueueReader) worker(d *amqp.Delivery) {
	var order common.Order
	err := jsonpb.Unmarshal(bytes.NewReader(d.Body), &order)
	if err != nil {
		return
	}

	groupId, found := qr.databaseSync.GetTelegramGroup(int(order.ShopId))
	if !found {
		return
	}

	text := fmt.Sprintf("Новый заказ\nПользователь: %s\nАдрес: %s\nТовары:", order.UserId, order.Address)
	for i, item := range order.Item {
		text += fmt.Sprintf("\n%d: %s", i + 1, item)
	}

	qr.telegramSender.SendMessage(text, groupId)
}

func (qr *QueueReader) Stop() {
	qr.queueReader.Stop()
}
