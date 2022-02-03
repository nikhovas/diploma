package combined

import (
	"context"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (s *Server) GetShopIdByKey(ctx context.Context, key *ctrlProto.ShopKey) (*ctrlProto.GetShopIdByKeyResponse, error) {
	shopId, err := s.kernel.GetShopIdByKey(ctx, key.GetCommon())
	if err != nil {
		return nil, err
	}

	return &ctrlProto.GetShopIdByKeyResponse{ShopId: int64(shopId)}, nil
}

func (s *Server) CheckContextWithShopId(ctx context.Context, messageInfo *ctrlProto.MessageInformation, shopId int, role string) (int, bool, error) {
	switch op := messageInfo.Info.(type) {
	case *ctrlProto.MessageInformation_Telegram:
		return s.telegram.CheckContextWithShopId(ctx, op.Telegram, shopId, role)
	default:
		panic("adsf")
	}
}
