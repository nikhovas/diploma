package grpcServer

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (s *Server) checkRequestWithShopId(
	ctx context.Context,
	r sq.BaseRunner,
	messageInfo *ctrlProto.MessageInformation,
	shopId int,
	role string,
) (bool, error) {
	userId, err := s.Combined.GetUser(ctx, r, messageInfo)
	if err != nil {
		return false, err
	}

	roleExists, err := s.Combined.CheckForRole(ctx, r, shopId, *userId, role)
	if err != nil {
		return false, err
	}

	return roleExists, nil
}

func (s *Server) checkRequestForAccessShop(
	ctx context.Context,
	r sq.BaseRunner,
	key *ctrlProto.ShopKey,
	msgInfo *ctrlProto.MessageInformation,
	role string,
) (int, *ctrlProto.DefaultResponse, error) {
	shopId, err := s.Combined.GetShopIdByShopKey(ctx, r, key)
	if err != nil {
		return 0, nil, err
	} else if shopId == nil {
		// TODO: no such shop response
		return 0, &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_NoRoleError{NoRoleError: &ctrlProto.NoRoleResponse{}}}, nil
	}

	userId, err := s.Combined.GetUser(ctx, r, msgInfo)
	if err != nil {
		return 0, nil, err
	} else if userId == nil {
		// TODO: no such shop response
		return 0, &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_NoRoleError{NoRoleError: &ctrlProto.NoRoleResponse{}}}, nil
	}

	grant, err := s.Combined.CheckForRole(ctx, r, *shopId, *userId, role)
	if err != nil {
		return 0, nil, err
	} else if !grant {
		return 0, &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_NoRoleError{NoRoleError: &ctrlProto.NoRoleResponse{}}}, nil
	}

	return *shopId, nil, nil
}
