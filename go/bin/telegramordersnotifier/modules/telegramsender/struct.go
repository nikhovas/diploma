package telegramsender

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nikhovas/diploma/go/lib/utils/foundation/modules/norun"
	"log"
)

type TelegramSender struct {
	norun.NoRun
	bot *tgbotapi.BotAPI
}

func FromConfig(config Config) *TelegramSender {
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic(err)
	}

	return &TelegramSender{
		bot: bot,
	}
}

func (ts *TelegramSender) SendMessage(message string, chatId int) {
	msg := tgbotapi.NewMessage(int64(chatId), message)
	_, err := ts.bot.Send(msg)
	if err != nil {
		fmt.Println(err)
	}
}
