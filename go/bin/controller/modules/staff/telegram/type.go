package telegram

import (
	tsb "github.com/nikhovas/diploma/go/lib/proto/staff_bot"
)

type Telegram struct {
	Bot tsb.TelegramStaffBotClient
}
