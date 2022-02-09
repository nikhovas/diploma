package telegram

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (t *Telegram) getTelegramGroupIdByShopId(ctx context.Context, r sq.BaseRunner, shopId int) (*int, error) {
	rows, err := sq.Select("telegram_group_id").From("telegram_staff_shop_group").
		Where(sq.Eq{"shop_id": shopId}).PlaceholderFormat(sq.Dollar).RunWith(r).QueryContext(ctx)
	if err != nil {
		return nil, err
	} else if !rows.Next() {
		return nil, nil
	}

	var groupId int
	_ = rows.Scan(&groupId)

	return &groupId, nil
}
