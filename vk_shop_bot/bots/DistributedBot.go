package vkDistributedWrapper

import (
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"log"
	"math/big"
	"strconv"
	"vk_shop_bot/vkApi"
	"vk_shop_bot/vkApi/VkApiServer"
	"vk_shop_bot/vkApi/VkLongPullServer"
)

type Bot struct {
	vkApi.VkBot
	//coordinator coordinator.ICoordinator
	coordinator *consulApi.KV
}

func (bot *Bot) GetCoordinatorBasePath() string {
	return fmt.Sprintf("services/vk-shop-bot/%d", bot.GroupId)
}

func (bot *Bot) GetStKvKey() string {
	return fmt.Sprintf("%s/ts", bot.GetCoordinatorBasePath())
}

func (bot *Bot) GetToken() string {
	return fmt.Sprintf("%s/token", bot.GetCoordinatorBasePath())
}

func (bot *Bot) GetTokenValue() (string, error) {
	kv, _, err := bot.coordinator.Get(bot.GetToken(), nil)
	if err != nil {
		return "", err
	}
	if kv == nil {
		return "", big.ErrNaN{}
	}

	return string(kv.Value), err
}

func (bot *Bot) Init(token string, groupId int, vkApiServer *VkApiServer.VkApiServer, coordinator *consulApi.KV) error {
	bot.GroupId = groupId
	bot.coordinator = coordinator
	var err error

	if token == "" {
		token, err = bot.GetTokenValue()
		if err != nil {
			return err
		}
	} else {
		_, _ = bot.coordinator.Put(&consulApi.KVPair{Key: bot.GetToken(), Value: []byte(token)}, nil)
	}

	return bot.VkBot.Init(token, groupId, vkApiServer)
}

func (bot *Bot) Authorize() (err error) {
	err = bot.VkBot.Authorize()
	if err != nil {
		return
	}

	stKv := &consulApi.KVPair{Key: bot.GetStKvKey(), Value: []byte(strconv.Itoa(bot.CurrentTs)), ModifyIndex: 0}
	_, _, err = bot.coordinator.CAS(stKv, nil)
	if err != nil {
		log.Fatal(err)
	}

	bot.CurrentTs, err = strconv.Atoi(string(stKv.Value))
	return
}

func CleanKv(kv *consulApi.KVPair) *consulApi.KVPair {
	return &consulApi.KVPair{Key: kv.Key, Value: kv.Value, ModifyIndex: kv.ModifyIndex}
}

func (bot *Bot) GetUpdates() (updates []VkLongPullServer.UpdateObject, err error) {
	stValueNodeKey := fmt.Sprintf("services/vk-shop-bot/%d/ts", bot.GroupId)
	stKv, _, err := bot.coordinator.Get(stValueNodeKey, nil)
	if err != nil {
		return
	}

	startedFromTs, err := strconv.Atoi(string(stKv.Value))
	if err != nil {
		return
	}

	bot.VkBot.CurrentTs = startedFromTs

	updates, err = bot.VkBot.GetUpdates()

	realStartFromTs := startedFromTs

	stKv.Value = []byte(strconv.Itoa(bot.CurrentTs))

	for {
		stKv = CleanKv(stKv)
		success, _, err := bot.coordinator.CAS(stKv, nil)
		if err != nil {
			log.Fatal(err)
		}

		if success {
			break
		}

		stKv, _, err = bot.coordinator.Get(stValueNodeKey, nil)
		var newTsValue int
		newTsValue, err = strconv.Atoi(string(stKv.Value))
		if bot.CurrentTs <= newTsValue {
			return []VkLongPullServer.UpdateObject{}, nil
		}
	}

	updates = updates[(realStartFromTs - startedFromTs):]
	return
}
