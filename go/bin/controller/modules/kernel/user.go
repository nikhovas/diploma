package kernel

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (kernel *Kernel) AddUser(ctx context.Context, r sq.BaseRunner) (int, error) {
	rows, err := sq.Insert("staff").Suffix("DEFAULT VALUES RETURNING id").PlaceholderFormat(sq.Dollar).RunWith(r).
		QueryContext(ctx)
	if err != nil {
		return 0, err
	} else if !rows.Next() {
		panic("No created element")
	}

	var userId int
	_ = rows.Scan(&userId)

	return userId, nil
}

func (kernel *Kernel) GetUserByUserKey(ctx context.Context, r sq.BaseRunner, userKey *ctrlProto.UserKey) (*int, error) {
	common := int(userKey.GetCommonUserId())
	return &common, nil
}
