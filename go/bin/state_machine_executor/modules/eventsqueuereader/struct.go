package eventsqueuereader

import (
	"bytes"
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	actions "github.com/nikhovas/diploma/go/lib/proto/consumer_actions"
	"github.com/nikhovas/diploma/go/lib/utils/foundation/modules/queuereader"
	amqp "github.com/rabbitmq/amqp091-go"
	"state_machine_executor/coremodules"
	"state_machine_executor/modules/eventobserver"
	"sync"
)

type EventsQueueReader struct {
	queueReader *queuereader.QueueReader
	readQueueWg      *sync.WaitGroup
	cm *coremodules.CoreModules
	eventObserver *eventobserver.EventObserver
}

func FromConfig(
	config Config,
	readQueueWg *sync.WaitGroup,
	cm *coremodules.CoreModules,
	eventObserver *eventobserver.EventObserver,
) *EventsQueueReader {
	eqr := &EventsQueueReader{}
	eqr.queueReader = queuereader.FromConfig(config.QueueReader, eqr.worker)
	eqr.readQueueWg = readQueueWg
	eqr.cm = cm
	eqr.eventObserver = eventObserver
	return eqr
}

func (eqr *EventsQueueReader) Run(ctx context.Context, wg *sync.WaitGroup) {
	eqr.queueReader.Run(ctx, wg)
}

func (eqr *EventsQueueReader) worker(d *amqp.Delivery) {
	eqr.cm.ReadQueueWg.Add(1)

	var ae actions.ActionEvent
	err := jsonpb.Unmarshal(bytes.NewReader(d.Body), &ae)
	if err != nil {
		return
	}

	fmt.Printf("Got action from service %s, user %s\n", ae.ServiceName, ae.UserId)

	eqr.eventObserver.Observe(context.Background(), &ae)
}

func (eqr *EventsQueueReader) Stop() {
	eqr.queueReader.Stop()
}

