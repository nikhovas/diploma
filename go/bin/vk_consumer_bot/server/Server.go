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

func (s *Server) SendSimpleMessage(
	ctx context.Context,
	information *consumerBot.SendSimpleMessageRequest,
) (*common.EmptyResponse, error) {
	_, err := s.Bot.SendMessage(
		int(information.Info.MsgLocation.GroupId),
		int(information.Info.MsgLocation.UserId),
		information.Info.Text,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &common.EmptyResponse{}, nil
}

func (s *Server) SendReplyMessage(
	ctx context.Context,
	information *consumerBot.SendReplyMessageRequest,
) (*common.EmptyResponse, error) {
	replyMsg := int(information.Info.ReplyMessageId)

	_, err := s.Bot.SendMessage(
		int(information.Info.MsgLocation.GroupId),
		int(information.Info.MsgLocation.UserId),
		information.Info.Text,
		&replyMsg,
	)
	if err != nil {
		return nil, err
	}
	return &common.EmptyResponse{}, nil
}

func NewServer(bot *bots.CombinedBot) *Server {
	s := &Server{Bot: bot}
	return s
}
