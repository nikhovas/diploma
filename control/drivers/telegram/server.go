package telegram

import (
	"context"
	pb "control/grpc/control"
)

func (t *Telegram) AddShop(ctx context.Context, commonBotInfo *pb.CommonBotInfo, telegramBotInfo *pb.TelegramBotInfo, messageInfo *pb.TelegramMessageInfo) (*pb.DefaultResponse, error) {
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

	return &pb.DefaultResponse{Resp: &pb.DefaultResponse_Success{}}, nil
}

func (t *Telegram) ListShops(ctx context.Context, messageInfo *pb.TelegramMessageInfo) (*pb.ListShopsResponse, error) {
	if !messageInfo.IsChatPrivate {
		return &pb.ListShopsResponse{Resp: &pb.ListShopsResponse_BadChatType{BadChatType: &pb.BadChatType{}}}, nil
	}

	const telegramListShopsQuery = `
	SELECT name
	FROM shop
	INNER JOIN telegram_staff_account ON telegram_staff_account.user_id = shop.creator
	WHERE telegram_staff_account.telegram_user_id = $1;
	`
	rows, _ := t.db.QueryContext(ctx, telegramListShopsQuery, messageInfo.UserId)
	var bots []*pb.CommonBotInfo
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		bots = append(bots, &pb.CommonBotInfo{Name: name})
	}

	resp := pb.ListBotSuccessResponse{Bots: bots}
	return &pb.ListShopsResponse{Resp: &pb.ListShopsResponse_Success{Success: &resp}}, nil
}

func (t *Telegram) ChangeBotState(ctx context.Context, messageInfo *pb.TelegramMessageInfo, toEnabled bool) (*pb.DefaultResponse, error) {
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
		return &pb.DefaultResponse{Resp: &pb.DefaultResponse_NoRoleError{NoRoleError: &pb.NoRoleResponse{}}}, nil
	}

	err = t.kernel.ChangeBotState(ctx, shopId, toEnabled)
	if err != nil {
		return nil, err
	}

	return &pb.DefaultResponse{Resp: &pb.DefaultResponse_Success{}}, nil
}

func (t *Telegram) ModifyShop(ctx context.Context, shopId int, newShopInfo *pb.OptionalBotInfo) (*pb.DefaultResponse, error) {
	err := t.UpdateTelegramPart(ctx, shopId, newShopInfo)
	if err != nil {
		return nil, err
	}

	err = t.kernel.UpdateShopCommonPart(ctx, shopId, newShopInfo)
	if err != nil {
		return nil, err
	}

	return &pb.DefaultResponse{Resp: &pb.DefaultResponse_Success{}}, nil
}

func (t *Telegram) DeleteShop(ctx context.Context, shopId int) (*pb.DefaultResponse, error) {
	err := t.DeleteTelegramPart(ctx, shopId)
	if err != nil {
		return nil, err
	}

	err = t.kernel.DeleteShopCommonPart(ctx, shopId)
	if err != nil {
		return nil, err
	}

	return &pb.DefaultResponse{Resp: &pb.DefaultResponse_Success{}}, nil
}
