package combined

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (combined *Combined) ChangeBotState(ctx context.Context, r sq.BaseRunner, shopId int, toEnabled bool) error {
	for _, c := range combined.Consumers {
		err := c.ChangeBotState(ctx, r, shopId, toEnabled)
		if err != nil {
			return err
		}
	}

	return nil
}

func (combined *Combined) NotifyBotStatusChange(ctx context.Context, r sq.BaseRunner, shopId int, toEnabled bool) error {
	for _, s := range combined.Staff {
		err := s.NotifyBotStatusChange(ctx, r, shopId, toEnabled)
		if err != nil {
			return err
		}
	}

	return nil
}
