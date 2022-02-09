package kernel

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (kernel *Kernel) CheckForRole(ctx context.Context, r sq.BaseRunner, shopId int, userId int, role string) (bool, error) {
	roles := sq.Or{sq.Eq{"role": "creator"}}
	if role != "creator" {
		roles = append(roles, sq.Eq{"role": "admin"})
	}
	if role != "admin" {
		roles = append(roles, sq.Eq{"role": role})
	}
	rows, err := sq.Select("1").From("shop_staff_roles").Where(sq.And{
		sq.Eq{"shop_id": shopId},
		sq.Eq{"user_id": userId},
		roles,
	}).PlaceholderFormat(sq.Dollar).RunWith(r).QueryContext(ctx)
	defer rows.Close()
	if err != nil {
		return false, err
	}

	exists := rows.Next()
	return exists, nil
}

func (kernel *Kernel) ModifyRole(ctx context.Context, r sq.BaseRunner, shopId int, userId int, role string, add bool) error {
	var err error
	if add {
		_, err = sq.Insert("shop_staff_roles").Columns("shop_id", "user_id", "role").
			Values(shopId, userId, role).PlaceholderFormat(sq.Dollar).RunWith(r).ExecContext(ctx)
	} else {
		_, err = sq.Delete("shop_staff_roles").Where(sq.Eq{"shop_id": shopId, "user_id": userId, "role": role}).
			PlaceholderFormat(sq.Dollar).RunWith(r).ExecContext(ctx)
	}
	if err != nil {
		return err
	}

	return nil
}
