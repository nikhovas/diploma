package telegram

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	tsb "github.com/nikhovas/diploma/go/lib/proto/staff_bot"
)

func (t *Telegram) NotifyBotStatusChange(ctx context.Context, r sq.BaseRunner, shopId int, toEnabled bool) error {
	groupId, err := t.getTelegramGroupIdByShopId(ctx, r, shopId)
	if err != nil {
		return err
	}

	_, err = t.Bot.NotifyBotStatusTelegramChange(
		ctx,
		&tsb.NotifyBotStatusChangeTelegramRequest{GroupId: int64(*groupId), Enabled: toEnabled},
	)
	if err != nil {
		return err
	}

	return nil
}
