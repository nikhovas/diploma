package grpcserver

import (
	"context"
	"github.com/nikhovas/diploma/go/lib/proto/common"
	consumerBot "github.com/nikhovas/diploma/go/lib/proto/consumer_bot"
	"vk_consumer_bot/modules/combinedbot"
)

type InternalServer struct {
	consumerBot.UnimplementedVkServerServer
	Bot *combinedbot.CombinedBot
}

func (is *InternalServer) SendSimpleMessage(
	ctx context.Context,
	information *consumerBot.SendSimpleMessageRequest,
) (*common.EmptyResponse, error) {
	_, err := is.Bot.SendMessage(
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

func (is *InternalServer) SendReplyMessage(
	ctx context.Context,
	information *consumerBot.SendReplyMessageRequest,
) (*common.EmptyResponse, error) {
	replyMsg := int(information.Info.ReplyMessageId)

	_, err := is.Bot.SendMessage(
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

func NewInternalServer(bot *combinedbot.CombinedBot) *InternalServer {
	s := &InternalServer{Bot: bot}
	return s
}
