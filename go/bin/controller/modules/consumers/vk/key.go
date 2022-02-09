package vk

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
	"log"
)

func (vk *Vk) GetShopIdByGroupId(ctx context.Context, r sq.BaseRunner, groupId int64) (*int, error) {
	rows, err := sq.Select("shop_id").From("vk_client_group").Where(sq.Eq{"vk_group_id": groupId}).
		PlaceholderFormat(sq.Dollar).RunWith(r).QueryContext(ctx)
	if err != nil {
		return nil, err
	} else if !rows.Next() {
		return nil, nil
	}

	var shopId int
	_ = rows.Scan(&shopId)
	return &shopId, nil
}

func (vk *Vk) GetShopIdByShopKey(ctx context.Context, r sq.BaseRunner, key *ctrlProto.ShopKey) (*int, error) {
	vkConsumerShopKey := key.GetVkConsumer()
	if vkConsumerShopKey == nil {
		return nil, nil
	}

	switch op := vkConsumerShopKey.Key.(type) {
	case *ctrlProto.VkConsumerShopKey_GroupId:
		return vk.GetShopIdByGroupId(ctx, r, op.GroupId)
	default:
		log.Fatal("adf")
	}
	return nil, nil
}
