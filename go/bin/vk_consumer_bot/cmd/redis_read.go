package main

import (
	"context"
	"fmt"
	"github.com/nikhovas/diploma/go/lib/utils/log"
	"strconv"
	"time"
	"vk_consumer_bot/application"
	"vk_consumer_bot/bots"
)

func redisReadWorker(ctx context.Context, app *application.Application, combinedBot *bots.CombinedBot) {
	log.Log(ctx, "Start read bot state updates from redis")
	enabledBotsDistFs := app.VkDistFsBase.CdCommon().CdEnabledBots()
	for update := range enabledBotsDistFs.GetUpdatesChan(ctx, make(map[string]struct{}), 7*time.Second) {
		log.Log(ctx, "Getting updates")
		if update.Error != nil {
			fmt.Println(update.Error)
		}

		log.Log(ctx, "Start deleting bots")
		for _, botStr := range update.Deleted {
			bot, _ := strconv.Atoi(botStr)
			combinedBot.RemoveBot(bot)
		}
		log.Log(ctx, "End deleting bots")

		log.Log(ctx, "Start adding bots")
		for _, botStr := range update.Added {
			bot, _ := strconv.Atoi(botStr)
			if err := combinedBot.AddBot(ctx, bot); err != nil {
				continue
			}
		}
		log.Log(ctx, "End adding bots")
	}
}
