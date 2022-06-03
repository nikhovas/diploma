package kernel

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
	"strings"
)

func (kernel *Kernel) GetShopItems(ctx context.Context, r sq.BaseRunner, key *ctrlProto.ShopKey) ([]string, error) {
	shopId, err := kernel.GetShopIdByShopKey(ctx, r, key)
	if err != nil {
		return nil, err
	}

	data, err := kernel.DistFsRoot.CdCommon().MetaCdShopId(*shopId).CdShopItems().Get(ctx)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(data, "\n")
	return parts, nil
}

func (kernel *Kernel) SetShopItems(ctx context.Context, r sq.BaseRunner, key *ctrlProto.ShopKey, data []string) error {
	shopId, err := kernel.GetShopIdByShopKey(ctx, r, key)
	if err != nil {
		return err
	}

	err = kernel.DistFsRoot.CdCommon().MetaCdShopId(*shopId).CdShopItems().Set(ctx, strings.Join(data, "\n"))
	if err != nil {
		return err
	}

	return nil
}
