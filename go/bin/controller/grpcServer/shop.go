package grpcServer

import (
	"context"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func generateDefaultOkResp() *ctrlProto.DefaultResponse {
	return &ctrlProto.DefaultResponse{
		Resp: &ctrlProto.DefaultResponse_Success{
			Success: &ctrlProto.EmptyOkResponse{},
		},
	}
}

func (s *Server) ListShops(ctx context.Context, request *ctrlProto.ListShopsRequest) (*ctrlProto.ListShopsResponse, error) {
	conn := s.Database

	userId, err := s.Combined.GetUser(ctx, conn, request.MessageInfo)
	if err != nil {
		return nil, err
	} else if userId == nil {
		// TODO: not found
	}

	shops, err := s.Combined.ListShops(ctx, conn, *userId)
	if err != nil {
		return nil, err
	}

	return &ctrlProto.ListShopsResponse{
		Resp: &ctrlProto.ListShopsResponse_Success{
			Success: &ctrlProto.ListBotSuccessResponse{
				Bots: shops,
			},
		},
	}, nil
}

func (s *Server) AddShop(ctx context.Context, req *ctrlProto.AddShopRequest) (*ctrlProto.DefaultResponse, error) {
	conn := s.Database

	userId, err := s.Combined.GetUserOrCreate(ctx, conn, req.MessageInfo)
	if err != nil {
		return nil, err
	}

	err = s.Combined.AddShop(ctx, conn, userId, req.Bot)
	if err != nil {
		return nil, err
	} else {
		return generateDefaultOkResp(), nil
	}
}

func (s *Server) ModifyShop(ctx context.Context, req *ctrlProto.ModifyShopRequest) (*ctrlProto.DefaultResponse, error) {
	conn := s.Database

	shopId, resp, err := s.checkRequestForAccessShop(ctx, conn, req.Key, req.MessageInfo, "admin")
	if resp != nil || err != nil {
		return resp, err
	}

	err = s.Combined.UpdateShop(ctx, conn, shopId, req.Bot)
	if err != nil {
		return nil, err
	}

	return generateDefaultOkResp(), nil
}

func (s *Server) DeleteShop(ctx context.Context, req *ctrlProto.DeleteShopRequest) (*ctrlProto.DefaultResponse, error) {
	conn := s.Database

	shopId, resp, err := s.checkRequestForAccessShop(ctx, conn, req.Key, req.MessageInfo, "creator")
	if resp != nil || err != nil {
		return resp, err
	}

	err = s.Combined.DeleteShop(ctx, conn, shopId)
	if err != nil {
		return nil, err
	}

	return generateDefaultOkResp(), nil
}
