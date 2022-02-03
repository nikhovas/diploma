package kernel

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
	qw "github.com/nikhovas/diploma/go/lib/proto/question_worker"
	tsb "github.com/nikhovas/diploma/go/lib/proto/staff_bot"
)

func (s *Kernel) AddShop(ctx context.Context, name string, creator int) (*ctrlProto.DefaultResponse, error) {
	const telegramAddShopQuery = `
	INSERT INTO shop(name, creator)
	VALUES ($1, $2)
	RETURNING id;
	`
	rows, err := s.db.QueryContext(ctx, telegramAddShopQuery, name, creator)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, &NoSuchItemError{}
	}

	var shopId int
	_ = rows.Scan(&shopId)
	fmt.Print(shopId)

	err = s.AddRole(ctx, shopId, creator, "creator")
	if err != nil {
		return nil, err
	}

	return &ctrlProto.DefaultResponse{Resp: &ctrlProto.DefaultResponse_Success{}}, nil
}

func (s *Kernel) AddQuestionAnswer(ctx context.Context, shopId int, question string, answer string) error {
	_, err := s.qwClient.AddQuestion(ctx, &qw.AddQuestionRequest{Question: question, Answer: answer})
	if err != nil {
		return err
	}

	return nil
}

func (s *Kernel) GetTelegramGrouctrlProtoyId(ctx context.Context, shopId int) (int, error) {
	queue, args, _ := sq.Select("telegram_group_id").From("telegram_staff_shop_group").
		Where(sq.Eq{"shop_id": shopId}).PlaceholderFormat(sq.Dollar).ToSql()
	rows, err := s.db.QueryContext(ctx, queue, args...)
	if err != nil {
		return 0, err
	} else if !rows.Next() {
		return 0, &NoSuchItemError{}
	}

	var groupId int
	_ = rows.Scan(&groupId)

	return groupId, nil
}

func (s *Kernel) ChangeBotState(ctx context.Context, shopId int, toEnabled bool) error {
	queue, args, _ := sq.Select("vk_group_id", "access_token").From("vk_client_group").
		Where(sq.Eq{"shop_id": shopId}).PlaceholderFormat(sq.Dollar).ToSql()

	rows, err := s.db.QueryContext(ctx, queue, args...)
	if err != nil {
		return err
	}
	if !rows.Next() {
		panic("No created element")
	}

	var vkGroupId int
	var token string
	_ = rows.Scan(&vkGroupId, &token)

	if toEnabled {
		s.rdb.LPush(ctx, "/bots/vk/enabled", vkGroupId)
		path := fmt.Sprintf("bots/data/vk/%d/token", vkGroupId)
		s.rdb.Set(ctx, path, token, 0)
	} else {
		s.rdb.LRem(ctx, "/bots/enabled", 0, vkGroupId)
	}

	return nil
}

func (s *Kernel) AddQuestion(ctx context.Context, shopId int, question string) error {
	path := fmt.Sprintf("bots/data/internal/%d/questions-list", shopId)
	s.rdb.LPush(ctx, path, question)

	groupId, err := s.GetTelegramGrouctrlProtoyId(ctx, shopId)
	if err != nil {
		return err
	}

	_, err = s.telegramStaffBot.SendNewQuestion(ctx, &tsb.NewQuestionRequest{Question: question, GroupId: int64(groupId)})
	if err != nil {
		return err
	}

	return nil
}

func (s *Kernel) NotifyBotStatusChange(ctx context.Context, shopId int, toEnabled bool) error {
	groupId, err := s.GetTelegramGrouctrlProtoyId(ctx, shopId)
	if err != nil {
		return err
	}

	_, err = s.telegramStaffBot.NotifyBotStatusTelegramChange(ctx, &tsb.NotifyBotStatusChangeTelegramRequest{GroupId: int64(groupId), Enabled: toEnabled})
	if err != nil {
		return err
	}

	return nil
}
