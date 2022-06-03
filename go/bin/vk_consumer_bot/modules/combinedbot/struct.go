package combinedbot

import (
	"context"
	"github.com/cornelk/hashmap"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	ctrl "github.com/nikhovas/diploma/go/lib/proto/controller"
	"github.com/nikhovas/diploma/go/lib/utils/distfs/bots"
	"github.com/nikhovas/diploma/go/lib/utils/log"
	apiServer "github.com/nikhovas/diploma/go/lib/vk/api_server"
	"math/big"
	"sync"
	"sync/atomic"
	"time"
	//"vk_consumer_bot/application"
	bots2 "vk_consumer_bot/bots"
	"vk_consumer_bot/localUtils"
	"vk_consumer_bot/modules/messageobserver"
)

type BotWrapper struct {
	Bot        *bots2.DistributedBot
	StopSignal int64
	Finished   int64
}

type CombinedBot struct {
	//app               *application.Application
	ctrlClient ctrl.ControllerClient
	bots              hashmap.HashMap
	coordinator       *api.KV
	wg                sync.WaitGroup
	globalStop        int64
	vkApiServer       *apiServer.VkApiServer
	messageObserver *messageobserver.MessageObserver
	stopped           int64
	serviceDistFsBase *bots.MetaServiceName
	wgFromRun         *sync.WaitGroup
}

func FromConfig(
	vkApiServer *apiServer.VkApiServer,
	messageObserver *messageobserver.MessageObserver,
	vkDistFsBase *bots.MetaServiceName,
	ctrlClient ctrl.ControllerClient,
) *CombinedBot {
	cb := &CombinedBot{}

	cb.vkApiServer = vkApiServer
	cb.messageObserver = messageObserver
	cb.stopped = 0
	cb.serviceDistFsBase = vkDistFsBase
	cb.ctrlClient = ctrlClient

	return cb
}

func (cb *CombinedBot) Run(ctx context.Context, wg *sync.WaitGroup) {
	cb.wgFromRun = wg
	wg.Add(1)
}

func (cb *CombinedBot) Stop() {
	atomic.StoreInt64(&cb.globalStop, 1)
	cb.wg.Wait()
	atomic.StoreInt64(&cb.stopped, 1)

	cb.wgFromRun.Done()
}

func (cb *CombinedBot) botWorker(bw *BotWrapper) {
	for {
		if atomic.LoadInt64(&bw.StopSignal) == 1 || atomic.LoadInt64(&cb.globalStop) == 1 {
			cb.wg.Done()
			atomic.StoreInt64(&bw.Finished, 1)
			return
		}

		//log.Printf("Getting updates for %d", bw.Bot.GroupId)
		_ = uuid.New()
		//ctx := context.WithValue(context.Background(), "uuid", u.String())
		ts, updates, _ := bw.Bot.GetUpdates()
		for _, update := range updates {
			cb.messageObserver.Callback(bw.Bot.GroupId, ts, update)
		}

		time.Sleep(3 * time.Second)
	}
}

func (cb *CombinedBot) AddBotInternal(groupId int) error {
	botWrapper := &BotWrapper{
		Bot:        nil,
		StopSignal: 0,
		Finished:   0,
	}

	actual, loaded := cb.bots.GetOrInsert(groupId, botWrapper)
	if loaded {
		return nil
	}

	botWrapper = actual.(*BotWrapper)

	var vkBot bots2.DistributedBot
	if err := vkBot.Init(groupId, cb.vkApiServer, cb.serviceDistFsBase); err != nil {
		return err
	}
	if err := vkBot.Authorize(); err != nil {
		return err
	}

	botWrapper.Bot = &vkBot

	cb.wg.Add(1)
	go cb.botWorker(botWrapper)

	return nil
}

func (cb *CombinedBot) AddBot(ctx context.Context, groupId int) error {
	if err := cb.AddBotInternal(groupId); err != nil {
		log.Log(ctx, "Cant start bot for %d. Error: %v", groupId, err)
		cb.RemoveBot(groupId)
		return err
	}

	log.Log(ctx, "Send notify message for %d", groupId)
	_, _ = cb.ctrlClient.NotifyBotStatusChange(context.Background(), &ctrl.NotifyBotStatusChangeRequest{
		Key: &ctrl.ShopKey{
			Key: &ctrl.ShopKey_VkConsumer{
				VkConsumer: &ctrl.VkConsumerShopKey{
					Key: &ctrl.VkConsumerShopKey_GroupId{
						GroupId: int64(groupId),
					},
				},
			},
		},
		ToEnabled: true,
	})

	return nil
}

func (cb *CombinedBot) RemoveBot(groupId int) {
	res, exists := cb.bots.Get(groupId)
	if !exists {
		return
	}

	botWrapper := res.(*BotWrapper)

	atomic.StoreInt64(&botWrapper.StopSignal, 1)

	for atomic.LoadInt64(&botWrapper.Finished) == 0 {
		time.Sleep(2 * time.Second)
	}

	cb.bots.Del(groupId)

	_, _ = cb.ctrlClient.NotifyBotStatusChange(context.Background(), &ctrl.NotifyBotStatusChangeRequest{
		Key:       localUtils.NewVkGroupIdShopKey(groupId),
		ToEnabled: false,
	})
}

func (cb *CombinedBot) SendMessage(groupId int, userId int, text string, replyTo *int) (int, error) {
	res, exists := cb.bots.Get(groupId)
	if !exists {
		return 0, big.ErrNaN{}
	}

	botWrapper := res.(*BotWrapper)

	return botWrapper.Bot.SendMessage(userId, text, replyTo)
}

