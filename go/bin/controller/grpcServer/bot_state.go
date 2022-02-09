package grpcServer

import (
	"context"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (s *Server) ChangeBotState(ctx context.Context, req *ctrlProto.ChangeBotStateRequest) (*ctrlProto.DefaultResponse, error) {
	conn := s.Database

	shopId, resp, err := s.checkRequestForAccessShop(ctx, conn, req.Key, req.MessageInfo, "admin")
	if resp != nil || err != nil {
		return resp, err
	}

	err = s.Combined.ChangeBotState(ctx, s.Database, shopId, req.ToEnabled)
	if err != nil {
		return nil, err
	}

	return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_Success{}}, nil
}

func (s *Server) NotifyBotStatusChange(ctx context.Context, req *ctrlProto.NotifyBotStatusChangeRequest) (*ctrlProto.DefaultResponse, error) {
	conn := s.Database

	shopId, err := s.Combined.GetShopIdByShopKey(ctx, conn, req.Key)
	if err != nil {
		return nil, err
	} else if shopId != nil {
		// TODO: error
	}

	err = s.Combined.NotifyBotStatusChange(ctx, conn, *shopId, req.ToEnabled)
	if err != nil {
		return nil, err
	}

	return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_Success{Success: &ctrlProto.EmptyOkResponse{}}}, nil
}
