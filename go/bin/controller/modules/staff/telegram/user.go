package telegram

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (t *Telegram) getUserInternal(ctx context.Context, r sq.BaseRunner, tgUserId *int64, tgChatId *int64) (*int, error) {
	var whereCond interface{}
	if tgUserId != nil && tgChatId != nil {
		whereCond = sq.Or{sq.Eq{"telegram_user_id": *tgUserId}, sq.Eq{"telegram_chat_id": *tgChatId}}
	} else if tgUserId != nil {
		whereCond = sq.Eq{"telegram_user_id": *tgUserId}
	} else if tgChatId != nil {
		whereCond = sq.Eq{"telegram_chat_id": *tgChatId}
	} else {
		return nil, nil
	}

	rows, err := sq.Select("user_id").From("telegram_staff_account").Where(whereCond).
		PlaceholderFormat(sq.Dollar).RunWith(r).QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var res int
	if !rows.Next() {
		return nil, nil
	}

	_ = rows.Scan(&res)
	return &res, nil
}

func (t *Telegram) GetUser(ctx context.Context, r sq.BaseRunner, messageInfo *ctrlProto.MessageInformation) (*int, error) {
	telegramMessageInfo := messageInfo.GetTelegram()
	if telegramMessageInfo == nil {
		return nil, nil
	}

	return t.getUserInternal(ctx, r, &telegramMessageInfo.UserId, &telegramMessageInfo.ChatId)
}

func (t *Telegram) AddUser(ctx context.Context, r sq.BaseRunner, userId int, messageInfo *ctrlProto.MessageInformation) error {
	telegramMessageInfo := messageInfo.GetTelegram()
	if telegramMessageInfo == nil {
		return nil
	}

	_, err := sq.Insert("telegram_staff_account").
		Columns("user_id", "telegram_user_id", "telegram_chat_id").
		Values(userId, telegramMessageInfo.UserId, telegramMessageInfo.ChatId).RunWith(r).ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (t *Telegram) GetUserByUserKey(ctx context.Context, r sq.BaseRunner, userKey *ctrlProto.UserKey) (*int, error) {
	telegramUserId := userKey.GetTelegramStaffUserId()
	return t.getUserInternal(ctx, r, &telegramUserId, nil)
}
