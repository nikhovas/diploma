package combined

import (
	"context"
	pb "control/grpc/control"
)

func (s *Server) GetShopIdByKey(ctx context.Context, key *pb.ShopKey) (int, error) {
	shopId, err := s.kernel.GetShopIdByKey(ctx, key.GetCommon())
	if err != nil {
		return 0, err
	}

	return shopId, nil
}

func (s *Server) CheckContextWithShopId(ctx context.Context, messageInfo *pb.MessageInformation, shopId int, role string) (int, bool, error) {
	switch op := messageInfo.Info.(type) {
	case *pb.MessageInformation_Telegram:
		return s.telegram.CheckContextWithShopId(ctx, op.Telegram, shopId, role)
	default:
		panic("adsf")
	}
}
