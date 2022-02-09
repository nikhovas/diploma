package vk

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (vk *Vk) AddShop(ctx context.Context, r sq.BaseRunner, shopId int, botInfo *ctrlProto.BotInfo) error {
	_, err := sq.Insert("vk_client_group").Columns("shop_id", "vk_group_id", "access_token").
		Values(shopId, botInfo.GetCommonBotInfo().GroupId, botInfo.GetCommonBotInfo().Token).
		PlaceholderFormat(sq.Dollar).RunWith(r).ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (vk *Vk) UpdateShop(ctx context.Context, r sq.BaseRunner, shopId int, info *ctrlProto.OptionalBotInfo) error {
	commonInfo := info.GetCommonBotInfo()
	if commonInfo == nil {
		return nil
	}

	shouldUpdate := false
	builder := sq.Update("vk_client_group").Where(sq.Eq{"shop_id": shopId})
	if commonInfo.Token != nil {
		builder = builder.Set("token", *commonInfo.Token)
		shouldUpdate = true
	}
	if commonInfo.GroupId != nil {
		builder = builder.Set("vk_client_group", *commonInfo.GroupId)
		shouldUpdate = true
	}

	if !shouldUpdate {
		return nil
	}

	_, err := builder.PlaceholderFormat(sq.Dollar).RunWith(r).ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (vk *Vk) DeleteShop(ctx context.Context, r sq.BaseRunner, shopId int) error {
	_, err := sq.Delete("vk_client_group").Where(sq.Eq{"shop_id": shopId}).PlaceholderFormat(sq.Dollar).
		RunWith(r).ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
