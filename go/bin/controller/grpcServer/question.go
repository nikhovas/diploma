package grpcServer

import (
	"context"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (s *Server) AddQuestionAnswer(ctx context.Context, req *ctrlProto.AddQuestionAnswerRequest) (*ctrlProto.DefaultResponse, error) {
	conn := s.Database

	shopId, resp, err := s.checkRequestForAccessShop(ctx, conn, req.Key, req.MessageInfo, "support")
	if resp != nil || err != nil {
		return resp, err
	}

	err = s.Combined.AddQuestionAnswer(ctx, conn, shopId, req.Question, req.Answer)
	if err != nil {
		return nil, err
	}

	return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_Success{Success: &ctrlProto.EmptyOkResponse{}}}, nil
}

func (s *Server) AddQuestion(ctx context.Context, req *ctrlProto.AddQuestionRequest) (*ctrlProto.DefaultResponse, error) {
	conn := s.Database

	shopId, err := s.Combined.GetShopIdByShopKey(ctx, conn, req.Key)
	if err != nil {
		return nil, err
	} else if shopId == nil {
		// TODO: not found
	}

	err = s.Combined.AddQuestion(ctx, conn, *shopId, req.GetInfo())
	if err != nil {
		return nil, err
	}

	return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_Success{Success: &ctrlProto.EmptyOkResponse{}}}, nil
}
