package combined

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (combined *Combined) ListShops(ctx context.Context, r sq.BaseRunner, userId int) ([]*ctrlProto.CommonBotInfo, error) {
	return combined.Kernel.ListShops(ctx, r, userId)
}

func (combined *Combined) AddShop(ctx context.Context, r sq.BaseRunner, userId int, botInfo *ctrlProto.BotInfo) error {
	shopId, err := combined.Kernel.AddShop(ctx, r, userId, botInfo)
	if err != nil {
		return err
	}

	for _, s := range combined.Staff {
		err := s.AddShop(ctx, r, shopId, botInfo)
		if err != nil {
			return err
		}
	}

	for _, c := range combined.Consumers {
		err := c.AddShop(ctx, r, shopId, botInfo)
		if err != nil {
			return err
		}
	}

	return nil
}

func (combined *Combined) UpdateShop(ctx context.Context, r sq.BaseRunner, shopId int, info *ctrlProto.OptionalBotInfo) error {
	err := combined.Kernel.UpdateShop(ctx, r, shopId, info)
	if err != nil {
		return err
	}

	for _, s := range combined.Staff {
		err := s.UpdateShop(ctx, r, shopId, info)
		if err != nil {
			return err
		}
	}

	for _, c := range combined.Consumers {
		err := c.UpdateShop(ctx, r, shopId, info)
		if err != nil {
			return err
		}
	}

	return nil
}

func (combined *Combined) DeleteShop(ctx context.Context, r sq.BaseRunner, shopId int) error {
	for _, s := range combined.Staff {
		err := s.DeleteShop(ctx, r, shopId)
		if err != nil {
			return err
		}
	}

	for _, c := range combined.Consumers {
		err := c.DeleteShop(ctx, r, shopId)
		if err != nil {
			return err
		}
	}

	err := combined.Kernel.DeleteShop(ctx, r, shopId)
	if err != nil {
		return err
	}

	return nil
}
