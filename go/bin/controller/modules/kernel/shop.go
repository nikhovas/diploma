package kernel

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (kernel *Kernel) ListShops(ctx context.Context, r sq.BaseRunner, userId int) ([]*ctrlProto.CommonBotInfo, error) {
	var res []*ctrlProto.CommonBotInfo
	rows, err := sq.Select("id", "name").From("shop").Where(sq.Eq{"creator": userId}).
		PlaceholderFormat(sq.Dollar).RunWith(r).QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		var name string
		_ = rows.Scan(&id, &name)
		res = append(res, &ctrlProto.CommonBotInfo{
			Name:    name,
			Token:   "",
			GroupId: 0,
		})
	}

	return res, nil
}

func (kernel *Kernel) AddShop(ctx context.Context, r sq.BaseRunner, userId int, botInfo *ctrlProto.BotInfo) (int, error) {
	rows, err := sq.Insert("shop").Columns("name", "creator").Values(botInfo.CommonBotInfo.Name, userId).
		Suffix("returning id").PlaceholderFormat(sq.Dollar).RunWith(r).QueryContext(ctx)
	if err != nil {
		return 0, err
	} else if !rows.Next() {
		panic("error")
	}

	var shopId int
	_ = rows.Scan(&shopId)
	fmt.Print(shopId)

	err = kernel.ModifyRole(ctx, r, shopId, userId, "creator", true)
	if err != nil {
		return 0, err
	}

	return shopId, nil
}

func (kernel *Kernel) UpdateShop(ctx context.Context, r sq.BaseRunner, shopId int, info *ctrlProto.OptionalBotInfo) error {
	commonInfo := info.GetCommonBotInfo()
	if commonInfo == nil {
		return nil
	}

	shouldUpdate := false
	builder := sq.Update("shop").Where(sq.Eq{"shop_id": shopId})
	if commonInfo.Name != nil {
		builder = builder.Set("name", *commonInfo.Name)
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

func (kernel *Kernel) DeleteShop(ctx context.Context, r sq.BaseRunner, shopId int) error {
	whereCond := sq.Eq{"shop_id": shopId}

	_, err := sq.Delete("shop").Where(whereCond).PlaceholderFormat(sq.Dollar).RunWith(r).ExecContext(ctx)
	if err != nil {
		return err
	}

	_, err = sq.Delete("shop_staff_roles").PlaceholderFormat(sq.Dollar).RunWith(r).ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}
