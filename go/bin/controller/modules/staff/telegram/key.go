package telegram

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (t *Telegram) GetShopIdByGroupId(ctx context.Context, r sq.BaseRunner, groupId int64) (*int, error) {
	rows, err := sq.Select("shop_id").From("telegram_staff_shop_group").
		Where(sq.Eq{"telegram_group_id": groupId}).RunWith(r).PlaceholderFormat(sq.Dollar).QueryContext(ctx)
	if err != nil {
		return nil, nil
	} else if !rows.Next() {
		return nil, nil
	}

	var res int
	_ = rows.Scan(&res)
	return &res, nil
}

func (t *Telegram) GetShopIdByShopKey(ctx context.Context, r sq.BaseRunner, key *ctrlProto.ShopKey) (*int, error) {
	telegramKey := key.GetTelegramStaff()
	if telegramKey == nil {
		return nil, nil
	}

	switch op := telegramKey.TelegramKey.(type) {
	case *ctrlProto.TelegramStaffShopKey_GroupId:
		return t.GetShopIdByGroupId(ctx, r, op.GroupId)
	default:
		return nil, nil
	}
}
