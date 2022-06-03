package telegramclient

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nikhovas/diploma/go/lib/utils/foundation"
	"log"
	"sync"
)

type MessageHandler interface {
	foundation.Module
	HandleMessage(upd *tgbotapi.Update)
}

type TelegramClient struct {
	bot *tgbotapi.BotAPI
	messageHandler MessageHandler
}

func FromConfig(config Config, messageHandler MessageHandler) *TelegramClient {
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic(err)
	}

	return &TelegramClient{
		bot: bot,
		messageHandler: messageHandler,
	}
}

func (tc *TelegramClient) Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go tc.RunBlock(ctx, wg)
}

func (tc *TelegramClient) RunBlock(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tc.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			tc.messageHandler.HandleMessage(&update)
		}
	}
}

func (tc *TelegramClient) SendMessage(c tgbotapi.Chattable) error {
	_, err := tc.bot.Send(c)
	return err
}

func (tc *TelegramClient) Stop() {
	tc.bot.StopReceivingUpdates()
}