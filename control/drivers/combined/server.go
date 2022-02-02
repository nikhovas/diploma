package combined

import (
	"context"
	pb "control/grpc/control"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListShops(ctx context.Context, request *pb.ListShopsRequest) (*pb.ListShopsResponse, error) {
	if t := request.MessageInfo.GetTelegram(); t != nil {
		return s.telegram.ListShops(ctx, t)
	}
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}

func (s *Server) AddShop(ctx context.Context, req *pb.AddShopRequest) (*pb.DefaultResponse, error) {
	switch op := req.MessageInfo.Info.(type) {
	case *pb.MessageInformation_Telegram:
		return s.telegram.AddShop(ctx, req.Bot.CommonBotInfo, req.Bot.PlatformBotInfo.GetTelegram(), op.Telegram)
	default:
		return nil, status.Errorf(codes.Unimplemented, "method AddShop not implemented")
	}
}

func (s *Server) ModifyShop(ctx context.Context, req *pb.ModifyShopRequest) (*pb.DefaultResponse, error) {
	shopId, err := s.GetShopIdByKey(ctx, req.Key)
	if err != nil {
		return nil, err
	}

	_, grant, err := s.CheckContextWithShopId(ctx, req.MessageInfo, shopId, "admin")
	if err != nil {
		return nil, err
	} else if !grant {
		return &pb.DefaultResponse{Resp: &pb.DefaultResponse_NoRoleError{NoRoleError: &pb.NoRoleResponse{}}}, nil
	}

	switch req.MessageInfo.Info.(type) {
	case *pb.MessageInformation_Telegram:
		return s.telegram.ModifyShop(ctx, shopId, req.GetBot())
	default:
		return nil, status.Errorf(codes.Unimplemented, "method TelegramModifyShop not implemented")
	}
}

func (s *Server) DeleteShop(ctx context.Context, req *pb.DeleteShopRequest) (*pb.DefaultResponse, error) {
	shopId, err := s.GetShopIdByKey(ctx, req.Key)
	if err != nil {
		return nil, err
	}

	_, grant, err := s.CheckContextWithShopId(ctx, req.MessageInfo, shopId, "creator")
	if err != nil {
		return nil, err
	} else if !grant {
		return &pb.DefaultResponse{Resp: &pb.DefaultResponse_NoRoleError{NoRoleError: &pb.NoRoleResponse{}}}, nil
	}

	switch req.MessageInfo.Info.(type) {
	case *pb.MessageInformation_Telegram:
		return s.telegram.DeleteShop(ctx, shopId)
	default:
		return nil, status.Errorf(codes.Unimplemented, "method TelegramDeleteShop not implemented")
	}
}

func (s *Server) AddQuestionAnswer(ctx context.Context, req *pb.AddQuestionAnswerRequest) (*pb.DefaultResponse, error) {
	shopId, err := s.GetShopIdByKey(ctx, req.Key)
	if err != nil {
		return nil, err
	}

	_, grant, err := s.CheckContextWithShopId(ctx, req.MessageInfo, shopId, "support")
	if err != nil {
		return nil, err
	} else if !grant {
		return &pb.DefaultResponse{Resp: &pb.DefaultResponse_NoRoleError{NoRoleError: &pb.NoRoleResponse{}}}, nil
	}

	err = s.kernel.AddQuestionAnswer(ctx, shopId, req.Question, req.Answer)
	if err != nil {
		return nil, err
	}

	err = s.kernel.AddQuestionAnswer(ctx, shopId, req.Question, req.Answer)
	if err != nil {
		return nil, err
	}

	return &pb.DefaultResponse{Resp: &pb.DefaultResponse_Success{Success: &pb.EmptyOkResponse{}}}, nil
}

func (s *Server) AddQuestion(ctx context.Context, req *pb.AddQuestionRequest) (*pb.DefaultResponse, error) {
	shopId, err := s.GetShopIdByKey(ctx, req.Key)
	if err != nil {
		return nil, err
	}

	err = s.kernel.AddQuestion(ctx, shopId, req.Question)
	if err != nil {
		return nil, err
	}

	return &pb.DefaultResponse{Resp: &pb.DefaultResponse_Success{Success: &pb.EmptyOkResponse{}}}, nil
}

func (s *Server) ChangeBotState(ctx context.Context, req *pb.ChangeBotStateRequest) (*pb.DefaultResponse, error) {
	switch op := req.MessageInfo.Info.(type) {
	case *pb.MessageInformation_Telegram:
		return s.telegram.ChangeBotState(ctx, op.Telegram, req.ToEnabled)
	default:
		return nil, status.Errorf(codes.Unimplemented, "method TelegramModifyShop not implemented")
	}
}

func (s *Server) ChangeUserRole(context.Context, *pb.ChangeUserRoleRequest) (*pb.DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TelegramChangeUserRole not implemented")
}

func (s *Server) NotifyBotStatusChange(ctx context.Context, req *pb.NotifyBotStatusChangeRequest) (*pb.DefaultResponse, error) {
	shopId, err := s.GetShopIdByKey(ctx, req.Key)
	if err != nil {
		return nil, err
	}

	err = s.kernel.NotifyBotStatusChange(ctx, shopId, req.ToEnabled)
	if err != nil {
		return nil, err
	}

	return &pb.DefaultResponse{Resp: &pb.DefaultResponse_Success{Success: &pb.EmptyOkResponse{}}}, nil
}
