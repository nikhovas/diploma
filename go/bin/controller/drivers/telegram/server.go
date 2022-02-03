package telegram

import (
	"context"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
)

func (t *Telegram) AddShop(ctx context.Context, commonBotInfo *ctrlProto.CommonBotInfo, telegramBotInfo *ctrlProto.TelegramBotInfo, messageInfo *ctrlProto.TelegramMessageInfo) (*ctrlProto.DefaultResponse, error) {
	if !messageInfo.IsChatPrivate {
		return generateBadChatTypeResponse()
	}

	userId, err := t.AddNewUserOrReturn(ctx, int(messageInfo.UserId), int(messageInfo.ChatId))
	if err != nil {
		return nil, err
	}

	shopId, err := t.kernel.AddShop(ctx, commonBotInfo.Name, userId)
	if err != nil {
		return nil, err
	}

	const telegramAddShopQuery = `
	INSERT INTO telegram_staff_shop_group(shop_id, telegram_group_id)
	VALUES ($1, 0);
	`
	_, err = t.db.ExecContext(ctx, telegramAddShopQuery, shopId)
	if err != nil {
		return nil, err
	}

	return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_Success{}}, nil
}

func (t *Telegram) ListShops(ctx context.Context, messageInfo *ctrlProto.TelegramMessageInfo) (*ctrlProto.ListShopsResponse, error) {
	if !messageInfo.IsChatPrivate {
		return &ctrlProto.ListShopsResponse{Resp: &ctrlProto.ListShopsResponse_BadChatType{BadChatType: &ctrlProto.BadChatType{}}}, nil
	}

	const telegramListShopsQuery = `
	SELECT name
	FROM shop
	INNER JOIN telegram_staff_account ON telegram_staff_account.user_id = shop.creator
	WHERE telegram_staff_account.telegram_user_id = $1;
	`
	rows, _ := t.db.QueryContext(ctx, telegramListShopsQuery, messageInfo.UserId)
	var bots []*ctrlProto.CommonBotInfo
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		bots = append(bots, &ctrlProto.CommonBotInfo{Name: name})
	}

	resp := ctrlProto.ListBotSuccessResponse{Bots: bots}
	return &ctrlProto.ListShopsResponse{Resp: &ctrlProto.ListShopsResponse_Success{Success: &resp}}, nil
}

func (t *Telegram) ChangeBotState(ctx context.Context, messageInfo *ctrlProto.TelegramMessageInfo, toEnabled bool) (*ctrlProto.DefaultResponse, error) {
	if messageInfo.IsChatPrivate {
		return generateBadChatTypeResponse()
	}

	userId, err := t.GetUserByTelegramUserId(ctx, messageInfo.UserId)
	if err != nil {
		return nil, err
	}

	shopId, err := t.GetShopIdByChatId(ctx, messageInfo.ChatId)
	if err != nil {
		return nil, err
	}

	grant, err := t.kernel.CheckForRole(ctx, shopId, userId, "admin")
	if err != nil {
		return nil, err
	} else if !grant {
		return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_NoRoleError{NoRoleError: &ctrlProto.NoRoleResponse{}}}, nil
	}

	err = t.kernel.ChangeBotState(ctx, shopId, toEnabled)
	if err != nil {
		return nil, err
	}

	return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_Success{}}, nil
}

func (t *Telegram) ModifyShop(ctx context.Context, shopId int, newShopInfo *ctrlProto.OptionalBotInfo) (*ctrlProto.DefaultResponse, error) {
	err := t.UpdateTelegramPart(ctx, shopId, newShopInfo)
	if err != nil {
		return nil, err
	}

	err = t.kernel.UpdateShopCommonPart(ctx, shopId, newShopInfo)
	if err != nil {
		return nil, err
	}

	return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_Success{}}, nil
}

func (t *Telegram) DeleteShop(ctx context.Context, shopId int) (*ctrlProto.DefaultResponse, error) {
	err := t.DeleteTelegramPart(ctx, shopId)
	if err != nil {
		return nil, err
	}

	err = t.kernel.DeleteShopCommonPart(ctx, shopId)
	if err != nil {
		return nil, err
	}

	return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_Success{}}, nil
}
