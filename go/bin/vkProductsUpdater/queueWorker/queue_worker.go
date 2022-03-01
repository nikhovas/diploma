package queueWorker

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/nikhovas/diploma/go/lib/proto/common"
	"github.com/nikhovas/diploma/go/lib/proto/vk_products_updater"
	"github.com/nikhovas/diploma/go/lib/utils/distQueue"
	"github.com/nikhovas/diploma/go/lib/utils/distfs"
	"github.com/nikhovas/diploma/go/lib/utils/distfs/bots"
	apiServer "github.com/nikhovas/diploma/go/lib/vk/api_server"
	"strconv"
	"sync"
	"time"
)

type QueueWorkerConfig struct {
	VkApiToken string `yaml:"vk_api_token"`
}

type QueueWorker struct {
	DistQueue   *distQueue.DistQueue
	VkDistFs    *bots.MetaServiceName
	VkApiServer *apiServer.VkApiServer

	marshaler jsonpb.Marshaler
	vkToken   string
}

func NewQueueWorker(config QueueWorkerConfig, dq *distQueue.DistQueue, distFsRoot *distfs.Root) (*QueueWorker, error) {
	return &QueueWorker{
		DistQueue:   dq,
		VkDistFs:    distFsRoot.CdBots().MetaCdServiceName("vk"),
		VkApiServer: apiServer.NewVkApiServer(),
		marshaler:   jsonpb.Marshaler{},
		vkToken:     config.VkApiToken,
	}, nil
}

func (qw *QueueWorker) RequestWork(ctx context.Context, req *vk_products_updater.UpdateProductsRequest) error {
	vkGroupDir := qw.VkDistFs.MetaCdGroupId(req.VkGroupId).CdCommon()
	updatedFlag := vkGroupDir.CdGoodsUpdated()
	goodsVar := vkGroupDir.CdGoods()

	flagTime := 5 * time.Minute
	set, err := updatedFlag.Set(ctx, flagTime)
	if err != nil {
		return err
	} else if !set {
		fmt.Println("Already exists")
		return nil
	}

	// work
	all, err := qw.VkApiServer.MMarketGetAll(qw.vkToken, int(req.VkGroupId))
	if err != nil {
		return err
	}

	for _, elem := range all {
		productElement := common.Product{
			Id:          int64(elem.Id),
			Title:       elem.Title,
			Description: elem.Description,
			Url:         "",
		}
		json, err := qw.marshaler.MarshalToString(&productElement)
		if err != nil {
			return err
		}
		err = goodsVar.Set(ctx, strconv.Itoa(elem.Id), json)
		if err != nil {
			return err
		}
	}

	delay := 18000
	if err := qw.DistQueue.Push(req, delay); err != nil {
		return err
	}

	return nil
}

func (qw *QueueWorker) Run(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	messages, err := qw.DistQueue.GetChan()
	if err != nil {
		return err
	}

	for {
		select {
		case msg := <-messages:
			if msg == nil {
				fmt.Println("Got nil message")
				continue
			}
			pb := msg.(*vk_products_updater.UpdateProductsRequest)
			if err := qw.RequestWork(ctx, pb); err != nil {
				_ = fmt.Errorf("error occured: %v", err)
			}
		case <-ctx.Done():
			break
		}
	}
}
