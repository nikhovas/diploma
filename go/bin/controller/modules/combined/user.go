package combined

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (combined *Combined) GetUserByUserKey(ctx context.Context, r sq.BaseRunner, userKey *ctrlProto.UserKey) (*int, error) {
	switch userKey.Key.(type) {
	case *ctrlProto.UserKey_CommonUserId:
		return combined.Kernel.GetUserByUserKey(ctx, r, userKey)
	case *ctrlProto.UserKey_TelegramStaffUserId:
		return combined.Staff["telegram"].GetUserByUserKey(ctx, r, userKey)
	default:
		return nil, nil
	}
}

func (combined *Combined) GetUser(ctx context.Context, r sq.BaseRunner, messageInfo *ctrlProto.MessageInformation) (*int, error) {
	for _, s := range combined.Staff {
		res, err := s.GetUser(ctx, r, messageInfo)
		if err != nil {
			return nil, err
		} else if res != nil {
			return res, nil
		}
	}

	return nil, nil
}

func (combined *Combined) AddUser(ctx context.Context, r sq.BaseRunner, messageInfo *ctrlProto.MessageInformation) (int, error) {
	userId, err := combined.Kernel.AddUser(ctx, r)
	if err != nil {
		return 0, err
	}

	for _, s := range combined.Staff {
		err := s.AddUser(ctx, r, userId, messageInfo)
		if err != nil {
			return 0, err
		}
	}

	return userId, nil
}

func (combined *Combined) GetUserOrCreate(ctx context.Context, r sq.BaseRunner, messageInfo *ctrlProto.MessageInformation) (int, error) {
	userId, err := combined.GetUser(ctx, r, messageInfo)
	if err != nil {
		return 0, err
	} else if userId != nil {
		return *userId, nil
	}

	return combined.AddUser(ctx, r, messageInfo)
}
