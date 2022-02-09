package grpcServer

import (
	"context"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (s *Server) GetShopIdByKey(ctx context.Context, key *ctrlProto.ShopKey) (*ctrlProto.GetShopIdByKeyResponse, error) {
	shopId, err := s.Combined.GetShopIdByShopKey(ctx, s.Database, key)
	if err != nil {
		return nil, err
	} else if shopId != nil {
		// TODO: error
	}

	return &ctrlProto.GetShopIdByKeyResponse{ShopId: int64(*shopId)}, nil
}
