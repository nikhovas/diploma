package kernel

import (
	"bytes"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/jsonpb"
	"github.com/nikhovas/diploma/go/lib/proto/common"
	"github.com/nikhovas/diploma/go/lib/proto/consumer_bot"
	qw "github.com/nikhovas/diploma/go/lib/proto/question_worker"
)

func (kernel *Kernel) AddQuestionAnswer(ctx context.Context, r sq.BaseRunner, shopId int, question string, answer string) error {
	shopIdDistDir := kernel.DistFsRoot.CdCommon().MetaCdShopId(shopId)
	questionSources := shopIdDistDir.CdQuestionSources()

	data, err := questionSources.Get(ctx, question)
	if err != nil {
		return err
	}

	var protoData common.WaitingQuesionInformation
	err = jsonpb.Unmarshal(bytes.NewReader([]byte(data)), &protoData)
	if err != nil {
		return err
	}

	_, err = kernel.QwClient.AddQuestion(
		ctx,
		&qw.AddQuestionRequest{
			Uuid:     "",
			Question: question,
			Answer:   answer,
			BasePath: shopIdDistDir.CdQa().Path,
		},
	)
	if err != nil {
		return err
	}

	err = questionSources.Delete(ctx, question)
	if err != nil {
		return err
	}

	answerTextWithoutReply := fmt.Sprintf("Вы задавали вопрос.\n%kernel\nОтвет.\n%kernel", question, answer)

	_, err = kernel.ConsumerBotClient.SendReplyMessage(
		ctx,
		&consumer_bot.SendReplyMessageRequest{
			Info: &consumer_bot.ReplyMessageInformation{
				MsgLocation:          protoData.MsgLocation,
				Text:                 answer,
				ReplyMessageId:       uint64(protoData.QuestionMsgId),
				ReplyUnsupportedText: answerTextWithoutReply,
			},
			Uuid: "",
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func (kernel *Kernel) AddQuestion(ctx context.Context, r sq.BaseRunner, shopId int, info *common.WaitingQuesionInformation) error {
	shopIdDistDir := kernel.DistFsRoot.CdCommon().MetaCdShopId(shopId)

	m := jsonpb.Marshaler{}
	questionSourceStr, _ := m.MarshalToString(info)
	err := shopIdDistDir.CdQuestionSources().Set(ctx, info.Question, questionSourceStr)
	if err != nil {
		return err
	}

	return nil
}
