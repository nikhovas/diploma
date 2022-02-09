package bots

import (
	"context"
	"github.com/cornelk/hashmap"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	ctrl "github.com/nikhovas/diploma/go/lib/proto/controller"
	"github.com/nikhovas/diploma/go/lib/utils/distfs/bots"
	"github.com/nikhovas/diploma/go/lib/utils/log"
	apiServer "github.com/nikhovas/diploma/go/lib/vk/api_server"
	longPullServer "github.com/nikhovas/diploma/go/lib/vk/long_pull_server"
	"math/big"
	"sync"
	"sync/atomic"
	"time"
	"vk_consumer_bot/application"
)

type BotWrapper struct {
	Bot        *DistributedBot
	StopSignal int64
	Finished   int64
}

type CombinedBot struct {
	app               *application.Application
	bots              hashmap.HashMap
	coordinator       *api.KV
	wg                sync.WaitGroup
	globalStop        int64
	vkApiServer       *apiServer.VkApiServer
	callback          func(app *application.Application, groupId int, ts int, update longPullServer.UpdateObject)
	stopped           int64
	serviceDistFsBase *bots.MetaServiceName
}

func (cb *CombinedBot) Init(
	app *application.Application,
	vkApiServer *apiServer.VkApiServer,
	callback func(app *application.Application, groupId int, ts int, update longPullServer.UpdateObject),
) {
	cb.app = app
	cb.vkApiServer = vkApiServer
	cb.callback = callback
	cb.stopped = 0
	cb.serviceDistFsBase = cb.app.VkDistFsBase
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
			cb.callback(cb.app, bw.Bot.GroupId, ts, update)
		}

		time.Sleep(3 * time.Second)
	}
}

func (cb *CombinedBot) Run(wg *sync.WaitGroup) {
	for atomic.LoadInt64(&cb.stopped) == 0 {
		time.Sleep(5 * time.Second)
	}

	wg.Done()
}

func (cb *CombinedBot) Stop() {
	atomic.StoreInt64(&cb.globalStop, 1)
	cb.wg.Wait()
	atomic.StoreInt64(&cb.stopped, 1)
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

	var vkBot DistributedBot
	if err := vkBot.Init(groupId, cb.vkApiServer, cb.app, cb.serviceDistFsBase); err != nil {
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
	_, _ = cb.app.CtrlClient.NotifyBotStatusChange(context.Background(), &ctrl.NotifyBotStatusChangeRequest{
		Key:       &ctrl.ShopKey{Key: &ctrl.ShopKey_Common{Common: &ctrl.CommonShopKey{CommonKey: &ctrl.CommonShopKey_VkGroupId{VkGroupId: int64(groupId)}}}},
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

	_, _ = cb.app.CtrlClient.NotifyBotStatusChange(context.Background(), &ctrl.NotifyBotStatusChangeRequest{
		Key:       &ctrl.ShopKey{Key: &ctrl.ShopKey_Common{Common: &ctrl.CommonShopKey{CommonKey: &ctrl.CommonShopKey_VkGroupId{VkGroupId: int64(groupId)}}}},
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
