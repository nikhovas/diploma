package combined

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (combined *Combined) CheckForRole(ctx context.Context, r sq.BaseRunner, shopId int, userId int, role string) (bool, error) {
	return combined.Kernel.CheckForRole(ctx, r, shopId, userId, role)
}

func (combined *Combined) ModifyRole(ctx context.Context, r sq.BaseRunner, shopId int, userId int, role string, add bool) error {
	return combined.Kernel.ModifyRole(ctx, r, shopId, userId, role, add)
}
