package interfaces

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

type KeyManager interface {
	GetShopIdByShopKey(ctx context.Context, r sq.BaseRunner, key *ctrlProto.ShopKey) (*int, error)
}

type ShopManager interface {
	AddShop(ctx context.Context, r sq.BaseRunner, shopId int, botInfo *ctrlProto.BotInfo) error
	UpdateShop(ctx context.Context, r sq.BaseRunner, shopId int, info *ctrlProto.OptionalBotInfo) error
	DeleteShop(ctx context.Context, r sq.BaseRunner, shopId int) error
}

type UserManager interface {
	GetUser(ctx context.Context, r sq.BaseRunner, messageInfo *ctrlProto.MessageInformation) (*int, error)
	AddUser(ctx context.Context, r sq.BaseRunner, userId int, messageInfo *ctrlProto.MessageInformation) error
	GetUserByUserKey(ctx context.Context, r sq.BaseRunner, userKey *ctrlProto.UserKey) (*int, error)
}

type Staff interface {
	KeyManager
	ShopManager
	UserManager
	AddQuestion(ctx context.Context, r sq.BaseRunner, shopId int, question string) error
	NotifyBotStatusChange(ctx context.Context, r sq.BaseRunner, shopId int, toEnabled bool) error
}

type Consumer interface {
	KeyManager
	ShopManager
	ChangeBotState(ctx context.Context, r sq.BaseRunner, shopId int, toEnabled bool) error
}
