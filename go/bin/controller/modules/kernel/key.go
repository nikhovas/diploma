package kernel

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (kernel *Kernel) GetShopIdByName(ctx context.Context, r sq.BaseRunner, name string) (*int, error) {
	rows, err := sq.Select("id").From("shop").Where(sq.Eq{"name": name}).RunWith(r).
		PlaceholderFormat(sq.Dollar).QueryContext(ctx)
	if err != nil {
		return nil, err
	} else if !rows.Next() {
		return nil, nil
	}

	var shopId int
	_ = rows.Scan(&shopId)
	return &shopId, nil
}

func (kernel *Kernel) GetShopIdByShopKey(ctx context.Context, r sq.BaseRunner, key *ctrlProto.ShopKey) (*int, error) {
	commonShopKey := key.GetCommon()
	if commonShopKey == nil {
		return nil, nil
	}

	switch op := commonShopKey.CommonKey.(type) {
	case *ctrlProto.CommonShopKey_Id:
		res := int(op.Id)
		return &res, nil
	case *ctrlProto.CommonShopKey_Name:
		return kernel.GetShopIdByName(ctx, r, op.Name)
	default:
		return nil, nil
	}
}
