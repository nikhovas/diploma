package workingbotsupdater

import (
	"context"
	"fmt"
	"github.com/nikhovas/diploma/go/lib/utils/distfs/bots"
	"github.com/nikhovas/diploma/go/lib/utils/foundation"
	"github.com/nikhovas/diploma/go/lib/utils/log"
	"strconv"
	"sync"
	"time"
	"vk_consumer_bot/modules/combinedbot"
)

type WorkingBotsUpdater struct {
	vkDistFsBase *bots.MetaServiceName
	combinedBot *combinedbot.CombinedBot
}

func FromConfig(config Config, vkDistFsBase *bots.MetaServiceName, combinedBot *combinedbot.CombinedBot) *WorkingBotsUpdater {
	return &WorkingBotsUpdater{
		vkDistFsBase: vkDistFsBase,
		combinedBot:  combinedBot,
	}
}

func (wbu *WorkingBotsUpdater) Run(ctx context.Context, wg *sync.WaitGroup) {
	updateTime := 10 * time.Second
	foundation.RunPeriodic(ctx, wg, wbu.worker, updateTime)
}

func (wbu *WorkingBotsUpdater) worker(ctx context.Context) {
	log.Log(ctx, "Start read bot state updates from redis")
	enabledBotsDistFs := wbu.vkDistFsBase.CdCommon().CdEnabledBots()
	for update := range enabledBotsDistFs.GetUpdatesChan(ctx, make(map[string]struct{}), 7*time.Second) {
		log.Log(ctx, "Getting updates")
		if update.Error != nil {
			fmt.Println(update.Error)
		}

		log.Log(ctx, "Start deleting bots")
		for _, botStr := range update.Deleted {
			bot, _ := strconv.Atoi(botStr)
			wbu.combinedBot.RemoveBot(bot)
		}
		log.Log(ctx, "End deleting bots")

		log.Log(ctx, "Start adding bots")
		for _, botStr := range update.Added {
			bot, _ := strconv.Atoi(botStr)
			if err := wbu.combinedBot.AddBot(ctx, bot); err != nil {
				continue
			}
		}
		log.Log(ctx, "End adding bots")
	}
}

func (wbu *WorkingBotsUpdater) Stop() {}
