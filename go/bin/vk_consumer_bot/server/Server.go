package server

import (
	"context"
	"github.com/nikhovas/diploma/go/lib/proto/common"
	consumerBot "github.com/nikhovas/diploma/go/lib/proto/consumer_bot"
	"vk_consumer_bot/bots"
)

type Server struct {
	consumerBot.UnimplementedVkServerServer
	Bot *bots.CombinedBot
}

func (s *Server) SendSimpleMessage(ctx context.Context, information *consumerBot.SimpleMessageInformation) (*common.EmptyResponse, error) {
	s.Bot.SendMessage(int(information.MessageDestination.GroupId), int(information.MessageDestination.UserId), information.Text)
	return &common.EmptyResponse{}, nil
}

func (s *Server) SendReplyMessage(ctx context.Context, information *consumerBot.ReplyMessageInformation) (*common.EmptyResponse, error) {
	s.Bot.SendMessage(int(information.MessageDestination.GroupId), int(information.MessageDestination.UserId), information.Text)
	return &common.EmptyResponse{}, nil
}

func (s *Server) AddBot(ctx context.Context, request *consumerBot.BotsActionRequest) (*common.EmptyResponse, error) {
	tokenPointer := request.Token
	token := ""
	if tokenPointer != nil {
		token = *tokenPointer
	}

	if err := s.Bot.AddBot(int(request.GroupId), token); err != nil {
		return nil, err
	}

	return &common.EmptyResponse{}, nil
}

func (s *Server) RemoveBot(ctx context.Context, request *consumerBot.BotsActionRequest) (*common.EmptyResponse, error) {
	s.Bot.RemoveBot(int(request.GroupId))
	return &common.EmptyResponse{}, nil
}

func NewServer(bot *bots.CombinedBot) *Server {
	s := &Server{Bot: bot}
	return s
}
