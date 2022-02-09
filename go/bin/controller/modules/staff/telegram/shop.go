package telegram

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (t *Telegram) AddShop(ctx context.Context, r sq.BaseRunner, shopId int, botInfo *ctrlProto.BotInfo) error {
	if botInfo.GetPlatformBotInfo() == nil {
		return nil
	}

	telegramBotInfo := botInfo.GetPlatformBotInfo().GetTelegram()
	if telegramBotInfo == nil {
		return nil
	}

	_, err := sq.Insert("telegram_staff_shop_group").Columns("shop_id", "telegram_group_id").
		Values(shopId, telegramBotInfo.ChatId).RunWith(r).ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (t *Telegram) UpdateShop(ctx context.Context, r sq.BaseRunner, shopId int, botInfo *ctrlProto.OptionalBotInfo) error {
	if botInfo.GetPlatformBotInfo() == nil {
		return nil
	}

	telegramBotInfo := botInfo.GetPlatformBotInfo().GetTelegram()
	if telegramBotInfo == nil {
		return nil
	}

	shouldUpdate := false
	builder := sq.Update("telegram_staff_shop_group").Where(sq.Eq{"shop_id": shopId})
	if telegramBotInfo.ChatId != nil {
		builder = builder.Set("telegram_group_id", *telegramBotInfo.ChatId)
		shouldUpdate = true
	}

	if !shouldUpdate {
		return nil
	}

	_, err := builder.RunWith(r).ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (t *Telegram) DeleteShop(ctx context.Context, r sq.BaseRunner, shopId int) error {
	_, err := sq.Delete("telegram_staff_shop_group").Where(sq.Eq{"shop_id": shopId}).RunWith(r).
		ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
