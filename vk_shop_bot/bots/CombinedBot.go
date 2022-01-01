package bots

import (
	"github.com/cornelk/hashmap"
	"github.com/hashicorp/consul/api"
	"log"
	"math/big"
	"sync"
	"sync/atomic"
	"time"
	"vk_shop_bot/vkApi/VkApiServer"
	"vk_shop_bot/vkApi/VkLongPullServer"
)

type BotWrapper struct {
	Bot        *DistributedBot
	StopSignal int64
	Finished   int64
}

type CombinedBot struct {
	bots        hashmap.HashMap
	coordinator *api.KV
	wg          sync.WaitGroup
	globalStop  int64
	vkApiServer *VkApiServer.VkApiServer
	callback    func(groupId int, update VkLongPullServer.UpdateObject)
	stopped     int64
}

func (cb *CombinedBot) Init(coordinator *api.KV, vkApiServer *VkApiServer.VkApiServer, callback func(groupId int, update VkLongPullServer.UpdateObject)) {
	cb.coordinator = coordinator
	cb.vkApiServer = vkApiServer
	cb.callback = callback
	cb.stopped = 0
}

func (cb *CombinedBot) botWorker(bw *BotWrapper) {
	for {
		if atomic.LoadInt64(&bw.StopSignal) == 1 || atomic.LoadInt64(&cb.globalStop) == 1 {
			cb.wg.Done()
			atomic.StoreInt64(&bw.Finished, 1)
			return
		}

		log.Printf("Getting updates for %d", bw.Bot.GroupId)
		updates, _ := bw.Bot.GetUpdates()
		for _, update := range updates {
			cb.callback(bw.Bot.GroupId, update)
		}

		time.Sleep(1 * time.Second)
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

func (cb *CombinedBot) AddBot(groupId int, token string) error {
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
	vkBot.Init(token, groupId, cb.vkApiServer, cb.coordinator)
	if err := vkBot.Authorize(); err != nil {
		return err
	}

	botWrapper.Bot = &vkBot

	cb.wg.Add(1)
	go cb.botWorker(botWrapper)

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
