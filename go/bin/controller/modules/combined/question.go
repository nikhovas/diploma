package combined

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/nikhovas/diploma/go/lib/proto/common"
)

func (combined *Combined) AddQuestionAnswer(ctx context.Context, r sq.BaseRunner, shopId int, question string, answer string) error {
	return combined.Kernel.AddQuestionAnswer(ctx, r, shopId, question, answer)
}

func (combined *Combined) AddQuestion(ctx context.Context, r sq.BaseRunner, shopId int, info *common.WaitingQuesionInformation) error {
	err := combined.Kernel.AddQuestion(ctx, r, shopId, info)
	if err != nil {
		return err
	}

	for _, s := range combined.Staff {
		err := s.AddQuestion(ctx, r, shopId, info.Question)
		if err != nil {
			return err
		}
	}

	return nil
}
