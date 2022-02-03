package bots

import (
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	apiServer "github.com/nikhovas/diploma/go/lib/vk/api_server"
	longPullServer "github.com/nikhovas/diploma/go/lib/vk/long_pull_server"
	"log"
	"math/big"
	"strconv"
)

type DistributedBot struct {
	VkBot
	coordinator *consulApi.KV
}

func (bot *DistributedBot) GetCoordinatorBasePath() string {
	return fmt.Sprintf("services/vk-shop-bot/%d", bot.GroupId)
}

func (bot *DistributedBot) GetStKvKey() string {
	return fmt.Sprintf("%s/ts", bot.GetCoordinatorBasePath())
}

func (bot *DistributedBot) GetToken() string {
	return fmt.Sprintf("%s/token", bot.GetCoordinatorBasePath())
}

func (bot *DistributedBot) GetTokenValue() (string, error) {
	kv, _, err := bot.coordinator.Get(bot.GetToken(), nil)
	if err != nil {
		return "", err
	}
	if kv == nil {
		return "", big.ErrNaN{}
	}

	return string(kv.Value), err
}

func (bot *DistributedBot) Init(token string, groupId int, vkApiServer *apiServer.VkApiServer, coordinator *consulApi.KV) error {
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

func (bot *DistributedBot) Authorize() (err error) {
	err = bot.VkBot.Authorize()
	if err != nil {
		return
	}

	stKv := &consulApi.KVPair{Key: bot.GetStKvKey(), Value: []byte(strconv.Itoa(bot.CurrentTs)), ModifyIndex: 0}

	var changed bool
	changed, _, err = bot.coordinator.CAS(stKv, nil)
	if err != nil {
		panic(err)
	}
	if changed {
		return
	}

	stop := false
	for !stop {
		stKv, _, err = bot.coordinator.Get(bot.GetStKvKey(), nil)
		coordinatorSt, _ := strconv.Atoi(string(stKv.Value))
		if bot.CurrentTs <= coordinatorSt {
			bot.CurrentTs = coordinatorSt
			stop = true
		} else {
			stKv = &consulApi.KVPair{Key: stKv.Key, Value: []byte(strconv.Itoa(bot.CurrentTs)), ModifyIndex: stKv.ModifyIndex}
			stop, _, err = bot.coordinator.CAS(stKv, nil)
		}
	}

	//bot.CurrentTs, err = strconv.Atoi(string(stKv.Value))
	return
}

func CleanKv(kv *consulApi.KVPair) *consulApi.KVPair {
	return &consulApi.KVPair{Key: kv.Key, Value: kv.Value, ModifyIndex: kv.ModifyIndex}
}

func (bot *DistributedBot) GetUpdates() (updTs int, updates []longPullServer.UpdateObject, err error) {
	stValueNodeKey := fmt.Sprintf("services/vk-shop-bot/%d/ts", bot.GroupId)
	stKv, _, err := bot.coordinator.Get(stValueNodeKey, nil)
	if err != nil {
		return
	}

	startedFromTs, err := strconv.Atoi(string(stKv.Value))
	if err != nil {
		return
	}

	if bot.VkBot.CurrentTs < startedFromTs {
		bot.VkBot.CurrentTs = startedFromTs
	}

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
			return 0, []longPullServer.UpdateObject{}, nil
		}
	}

	updates = updates[(realStartFromTs - startedFromTs):]
	updTs = realStartFromTs
	return
}
