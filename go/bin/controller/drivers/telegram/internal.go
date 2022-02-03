package telegram

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/hashicorp/go-multierror"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (t *Telegram) UpdateTelegramPart(ctx context.Context, shopId int, info *ctrlProto.OptionalBotInfo) error {
	platformShopInfo := info.GetPlatformBotInfo()
	if platformShopInfo == nil {
		return nil
	}

	telegramShopInfo := platformShopInfo.GetTelegram()
	if telegramShopInfo == nil {
		return nil
	}

	shouldUpdate := false
	builder := sq.Update("telegram_staff_shop_group").Where(sq.Eq{"shop_id": shopId})
	if telegramShopInfo.ChatId != nil {
		builder = builder.Set("telegram_group_id", *telegramShopInfo.ChatId)
		shouldUpdate = true
	}

	if shouldUpdate {
		queue, _, _ := builder.ToSql()
		_, err := t.db.ExecContext(ctx, queue)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Telegram) DeleteTelegramPart(ctx context.Context, shopId int) error {
	queue, _, _ := sq.Delete("telegram_staff_shop_group").Where(sq.Eq{"shop_id": shopId}).ToSql()
	_, err := t.db.ExecContext(ctx, queue)
	if err != nil {
		return err
	}

	return nil
}

const checkTelegramUserExists = `
SELECT user_id
FROM telegram_staff_account
WHERE telegram_user_id = $1;
`

const addUserToTelegramTable = `
INSERT INTO telegram_staff_account(user_id, telegram_user_id, telegram_chat_id)
VALUES ($1, $2, $3)
RETURNING *;
`

func (t *Telegram) AddNewUserOrReturn(ctx context.Context, telegramUserId int, telegramChatId int) (int, error) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	rows, err := tx.QueryContext(ctx, checkTelegramUserExists, telegramUserId)
	if err != nil {
		return 0, multierror.Append(err, tx.Rollback())
	}

	if rows.Next() {
		var res int
		_ = rows.Scan(&res)
		_ = tx.Commit()
		return res, nil
	}

	userId, err := t.kernel.AddNewUserOrReturn(ctx)
	if err != nil {
		return 0, err
	}

	if _, err = tx.ExecContext(ctx, addUserToTelegramTable, userId, telegramUserId, telegramChatId); err != nil {
		return 0, multierror.Append(err, tx.Rollback())
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return userId, nil
}

func (t *Telegram) GetUserByTelegramUserId(ctx context.Context, telegramUserId int64) (int, error) {
	rows, err := t.db.QueryContext(ctx, checkTelegramUserExists, telegramUserId)
	if err != nil {
		return 0, err
	}

	var res int

	if !rows.Next() {
		return 0, nil // no id
	}

	_ = rows.Scan(&res)
	return res, nil
}

const checkTelegramUserExistsByChat = `
SELECT user_id
FROM telegram_staff_account
WHERE telegram_chat_id = ?;
`

func (t *Telegram) GetUserByTelegramChatId(ctx context.Context, telegramChatId int) (int, error) {
	rows, err := t.db.QueryContext(ctx, checkTelegramUserExistsByChat, telegramChatId)
	if err != nil {
		return 0, err
	}

	var res int

	if !rows.Next() {
		return 0, nil // no id
	}

	_ = rows.Scan(&res)
	return res, nil
}

func (t *Telegram) GetShopIdByChatId(ctx context.Context, telegramChatId int64) (int, error) {
	const queue = `
	SELECT shop_id
	FROM telegram_staff_shop_group
	WHERE telegram_group_id = $1;
	`

	rows, err := t.db.QueryContext(ctx, queue, telegramChatId)
	if err != nil {
		return 0, err
	}

	var res int

	if !rows.Next() {
		return 0, nil // no id
	}

	_ = rows.Scan(&res)
	return res, nil
}

func (t *Telegram) CheckContextWithShopId(ctx context.Context, messageInfo *ctrlProto.TelegramMessageInfo, shopId int, role string) (int, bool, error) {
	userId, err := t.GetUserByTelegramUserId(ctx, messageInfo.UserId)
	if err != nil {
		return 0, false, err
	}

	if role != "" {
		grant, err := t.kernel.CheckForRole(ctx, shopId, userId, "admin")
		if err != nil {
			return userId, false, err
		} else if !grant {
			return userId, false, nil
		}
	}

	return userId, true, nil
}
