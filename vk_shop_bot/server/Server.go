package server

import (
	"context"
	"vk_shop_bot/bots"
	"vk_shop_bot/proto"
	grpcDefine "vk_shop_bot/proto"
)

type Server struct {
	grpcDefine.UnimplementedVkServerServer
	Bot *bots.CombinedBot
}

func (s *Server) SendMessage(ctx context.Context, request *proto.SendMessageRequest) (*proto.EmptyResponse, error) {
	s.Bot.SendMessage(int(request.GroupId), int(request.UserId), request.Text)
	return &proto.EmptyResponse{}, nil
}

func (s *Server) AddBot(ctx context.Context, request *proto.BotsActionRequest) (*proto.EmptyResponse, error) {
	tokenPointer := request.Token
	token := ""
	if tokenPointer != nil {
		token = *tokenPointer
	}

	if err := s.Bot.AddBot(int(request.GroupId), token); err != nil {
		return &proto.EmptyResponse{}, err
	}

	return &proto.EmptyResponse{}, nil
}

func (s *Server) RemoveBot(ctx context.Context, request *proto.BotsActionRequest) (*proto.EmptyResponse, error) {
	s.Bot.RemoveBot(int(request.GroupId))
	return &proto.EmptyResponse{}, nil
}

func NewServer(bot *bots.CombinedBot) *Server {
	s := &Server{Bot: bot}
	return s
}
