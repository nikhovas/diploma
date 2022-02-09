package grpcServer

import (
	"context"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (s *Server) ChangeUserRole(
	ctx context.Context,
	req *ctrlProto.ChangeUserRoleRequest,
) (*ctrlProto.DefaultResponse, error) {
	conn := s.Database

	shopId, resp, err := s.checkRequestForAccessShop(ctx, conn, req.Key, req.MessageInfo, "admin")
	if resp != nil || err != nil {
		return resp, err
	}

	changingUserId, err := s.Combined.GetUserByUserKey(ctx, conn, req.ModifyUserId)
	if err != nil {
		return nil, err
	} else if changingUserId == nil {
		// TODO: not found
	}

	err = s.Combined.ModifyRole(ctx, conn, shopId, *changingUserId, req.Role, req.IsAddAction)
	if err != nil {
		return nil, err
	}

	return generateDefaultOkResp(), nil
}
