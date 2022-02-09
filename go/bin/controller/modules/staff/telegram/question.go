package telegram

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	tsb "github.com/nikhovas/diploma/go/lib/proto/staff_bot"
)

func (t *Telegram) AddQuestion(ctx context.Context, r sq.BaseRunner, shopId int, question string) error {
	groupId, err := t.getTelegramGroupIdByShopId(ctx, r, shopId)
	if err != nil {
		return err
	}

	_, err = t.Bot.SendNewQuestion(
		ctx,
		&tsb.NewQuestionRequest{
			Question: question,
			GroupId:  int64(*groupId),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
