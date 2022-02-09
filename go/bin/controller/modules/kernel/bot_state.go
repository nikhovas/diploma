package kernel

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (kernel *Kernel) ChangeBotState(ctx context.Context, r sq.BaseRunner, shopId int, toEnabled bool) error {
	return nil
}

func (kernel *Kernel) NotifyBotStatusChange(ctx context.Context, r sq.BaseRunner, shopId int, toEnabled bool) error {
	return nil
}
