package combined

import (
	"context"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListShops(ctx context.Context, request *ctrlProto.ListShopsRequest) (*ctrlProto.ListShopsResponse, error) {
	if t := request.MessageInfo.GetTelegram(); t != nil {
		return s.telegram.ListShops(ctx, t)
	}
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

func (s *Server) AddShop(ctx context.Context, req *ctrlProto.AddShopRequest) (*ctrlProto.DefaultResponse, error) {
	switch op := req.MessageInfo.Info.(type) {
	case *ctrlProto.MessageInformation_Telegram:
		return s.telegram.AddShop(ctx, req.Bot.CommonBotInfo, req.Bot.PlatformBotInfo.GetTelegram(), op.Telegram)
	default:
		return nil, status.Errorf(codes.Unimplemented, "method AddShop not implemented")
	}
}

func (s *Server) ModifyShop(ctx context.Context, req *ctrlProto.ModifyShopRequest) (*ctrlProto.DefaultResponse, error) {
	shopId, err := s.GetShopIdByKey(ctx, req.Key)
	if err != nil {
		return nil, err
	}

	_, grant, err := s.CheckContextWithShopId(ctx, req.MessageInfo, int(shopId.ShopId), "admin")
	if err != nil {
		return nil, err
	} else if !grant {
		return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_NoRoleError{NoRoleError: &ctrlProto.NoRoleResponse{}}}, nil
	}

	switch req.MessageInfo.Info.(type) {
	case *ctrlProto.MessageInformation_Telegram:
		return s.telegram.ModifyShop(ctx, int(shopId.ShopId), req.GetBot())
	default:
		return nil, status.Errorf(codes.Unimplemented, "method TelegramModifyShop not implemented")
	}
}

func (s *Server) DeleteShop(ctx context.Context, req *ctrlProto.DeleteShopRequest) (*ctrlProto.DefaultResponse, error) {
	shopId, err := s.GetShopIdByKey(ctx, req.Key)
	if err != nil {
		return nil, err
	}

	_, grant, err := s.CheckContextWithShopId(ctx, req.MessageInfo, int(shopId.ShopId), "creator")
	if err != nil {
		return nil, err
	} else if !grant {
		return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_NoRoleError{NoRoleError: &ctrlProto.NoRoleResponse{}}}, nil
	}

	switch req.MessageInfo.Info.(type) {
	case *ctrlProto.MessageInformation_Telegram:
		return s.telegram.DeleteShop(ctx, int(shopId.ShopId))
	default:
		return nil, status.Errorf(codes.Unimplemented, "method TelegramDeleteShop not implemented")
	}
}

func (s *Server) AddQuestionAnswer(ctx context.Context, req *ctrlProto.AddQuestionAnswerRequest) (*ctrlProto.DefaultResponse, error) {
	shopId, err := s.GetShopIdByKey(ctx, req.Key)
	if err != nil {
		return nil, err
	}

	_, grant, err := s.CheckContextWithShopId(ctx, req.MessageInfo, int(shopId.ShopId), "support")
	if err != nil {
		return nil, err
	} else if !grant {
		return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_NoRoleError{NoRoleError: &ctrlProto.NoRoleResponse{}}}, nil
	}

	err = s.kernel.AddQuestionAnswer(ctx, int(shopId.ShopId), req.Question, req.Answer)
	if err != nil {
		return nil, err
	}

	err = s.kernel.AddQuestionAnswer(ctx, int(shopId.ShopId), req.Question, req.Answer)
	if err != nil {
		return nil, err
	}

	return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_Success{Success: &ctrlProto.EmptyOkResponse{}}}, nil
}

func (s *Server) AddQuestion(ctx context.Context, req *ctrlProto.AddQuestionRequest) (*ctrlProto.DefaultResponse, error) {
	shopId, err := s.GetShopIdByKey(ctx, req.Key)
	if err != nil {
		return nil, err
	}

	err = s.kernel.AddQuestion(ctx, int(shopId.ShopId), req.Question)
	if err != nil {
		return nil, err
	}

	return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_Success{Success: &ctrlProto.EmptyOkResponse{}}}, nil
}

func (s *Server) ChangeBotState(ctx context.Context, req *ctrlProto.ChangeBotStateRequest) (*ctrlProto.DefaultResponse, error) {
	switch op := req.MessageInfo.Info.(type) {
	case *ctrlProto.MessageInformation_Telegram:
		return s.telegram.ChangeBotState(ctx, op.Telegram, req.ToEnabled)
	default:
		return nil, status.Errorf(codes.Unimplemented, "method TelegramModifyShop not implemented")
	}
}

func (s *Server) ChangeUserRole(context.Context, *ctrlProto.ChangeUserRoleRequest) (*ctrlProto.DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TelegramChangeUserRole not implemented")
}

func (s *Server) NotifyBotStatusChange(ctx context.Context, req *ctrlProto.NotifyBotStatusChangeRequest) (*ctrlProto.DefaultResponse, error) {
	shopId, err := s.GetShopIdByKey(ctx, req.Key)
	if err != nil {
		return nil, err
	}

	err = s.kernel.NotifyBotStatusChange(ctx, int(shopId.ShopId), req.ToEnabled)
	if err != nil {
		return nil, err
	}

	return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_Success{Success: &ctrlProto.EmptyOkResponse{}}}, nil
}
