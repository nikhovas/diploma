package combined

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (combined *Combined) GetShopIdByShopKey(ctx context.Context, r sq.BaseRunner, key *ctrlProto.ShopKey) (*int, error) {
	switch key.Key.(type) {
	case *ctrlProto.ShopKey_Common:
		return combined.Kernel.GetShopIdByShopKey(ctx, r, key)
	case *ctrlProto.ShopKey_TelegramStaff:
		return combined.Staff["telegram"].GetShopIdByShopKey(ctx, r, key)
	case *ctrlProto.ShopKey_VkConsumer:
		return combined.Consumers["vk"].GetShopIdByShopKey(ctx, r, key)
	default:
		return nil, nil
	}
}
