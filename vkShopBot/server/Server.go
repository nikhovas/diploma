package server

import (
	"context"
	"vkShopBot/bots"

	pb "github.com/nikhovas/diploma/proto/servers/VkServer"
)

type Server struct {
	pb.UnimplementedVkServerServer
	Bot *bots.CombinedBot
}

func (s *Server) SendMessage(ctx context.Context, request *pb.SendMessageRequest) (*pb.EmptyResponse, error) {
	s.Bot.SendMessage(int(request.GroupId), int(request.UserId), request.Text)
	return &pb.EmptyResponse{}, nil
}

func (s *Server) AddBot(ctx context.Context, request *pb.BotsActionRequest) (*pb.EmptyResponse, error) {
	tokenPointer := request.Token
	token := ""
	if tokenPointer != nil {
		token = *tokenPointer
	}

	if err := s.Bot.AddBot(int(request.GroupId), token); err != nil {
		return &pb.EmptyResponse{}, err
	}

	return &pb.EmptyResponse{}, nil
}

func (s *Server) RemoveBot(ctx context.Context, request *pb.BotsActionRequest) (*pb.EmptyResponse, error) {
	s.Bot.RemoveBot(int(request.GroupId))
	return &pb.EmptyResponse{}, nil
}

func NewServer(bot *bots.CombinedBot) *Server {
	s := &Server{Bot: bot}
	return s
}
