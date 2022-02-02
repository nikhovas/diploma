package bots

import (
	"context"
	"github.com/cornelk/hashmap"
	"github.com/hashicorp/consul/api"
	"github.com/nikhovas/diploma/lib/go/vk/apiServer"
	"github.com/nikhovas/diploma/lib/go/vk/longPullServer"
	"log"
	"math/big"
	"sync"
	"sync/atomic"
	"time"
	ctrl "vkShopBot/grpc/control"
)

type BotWrapper struct {
	Bot        *DistributedBot
	StopSignal int64
	Finished   int64
}

type CombinedBot struct {
	control     ctrl.ControlClient
	bots        hashmap.HashMap
	coordinator *api.KV
	wg          sync.WaitGroup
	globalStop  int64
	vkApiServer *apiServer.VkApiServer
	callback    func(groupId int, ts int, update longPullServer.UpdateObject)
	stopped     int64
}

func (cb *CombinedBot) Init(coordinator *api.KV, vkApiServer *apiServer.VkApiServer, callback func(groupId int, ts int, update longPullServer.UpdateObject), control ctrl.ControlClient) {
	cb.coordinator = coordinator
	cb.vkApiServer = vkApiServer
	cb.callback = callback
	cb.stopped = 0
	cb.control = control
}

func (cb *CombinedBot) botWorker(bw *BotWrapper) {
	for {
		if atomic.LoadInt64(&bw.StopSignal) == 1 || atomic.LoadInt64(&cb.globalStop) == 1 {
			cb.wg.Done()
			atomic.StoreInt64(&bw.Finished, 1)
			return
		}

		log.Printf("Getting updates for %d", bw.Bot.GroupId)
		ts, updates, _ := bw.Bot.GetUpdates()
		for _, update := range updates {
			cb.callback(bw.Bot.GroupId, ts, update)
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

func (cb *CombinedBot) AddBotInternal(groupId int, token string) error {
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
	if err := vkBot.Init(token, groupId, cb.vkApiServer, cb.coordinator); err != nil {
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

func (cb *CombinedBot) AddBot(groupId int, token string) error {
	if err := cb.AddBotInternal(groupId, token); err != nil {
		cb.RemoveBot(groupId)
		return err
	}

	_, _ = cb.control.NotifyBotStatusChange(context.Background(), &ctrl.NotifyBotStatusChangeRequest{
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
}

func (cb *CombinedBot) SendMessage(groupId int, userId int, text string) (int, error) {
	res, exists := cb.bots.Get(groupId)
	if !exists {
		return 0, big.ErrNaN{}
	}

	botWrapper := res.(*BotWrapper)

	return botWrapper.Bot.SendMessage(userId, text)
}
