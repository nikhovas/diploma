package vk

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"strconv"
)

func (vk *Vk) ChangeBotState(ctx context.Context, r sq.BaseRunner, shopId int, toEnabled bool) error {
	rows, err := sq.Select("vk_group_id", "access_token").From("vk_client_group").
		Where(sq.Eq{"shop_id": shopId}).PlaceholderFormat(sq.Dollar).RunWith(r).QueryContext(ctx)
	if err != nil {
		return err
	}
	if !rows.Next() {
		panic("No created element")
	}

	var vkGroupId int
	var token string
	_ = rows.Scan(&vkGroupId, &token)

	enabledBotsVar := vk.DistFsMetaServiceName.CdCommon().CdEnabledBots()
	tokenBotVar := vk.DistFsMetaServiceName.MetaCdGroupId(vkGroupId).CdCommon().CdToken()

	if toEnabled {
		if err := tokenBotVar.Set(ctx, token); err != nil {
			return err
		}
		if err := enabledBotsVar.Add(ctx, strconv.Itoa(vkGroupId)); err != nil {
			return err
		}
	} else {
		if err := enabledBotsVar.Delete(ctx, strconv.Itoa(vkGroupId)); err != nil {
			return err
		}
		if err := tokenBotVar.Set(ctx, token); err != nil {
			return err
		}
	}

	return nil
}
