package states

import (
	"context"
	"fmt"
	"github.com/nikhovas/diploma/go/lib/proto/common"
	"github.com/nikhovas/diploma/go/lib/proto/consumer_messages"
	"state_machine_executor/coremodules"
	"state_machine_executor/state_machine/localStorage"
	"strconv"
)

type CheckOrder struct {
	ConfirmOrderStateName string
	WaitInputStateName string
}

func (s *CheckOrder) Process(ctx context.Context, cm *coremodules.CoreModules, ls *localStorage.Storage) (string, bool) {
	addressInterface := ls.KvStorage.Get("address")
	address := ""
	if addressInterface != nil {
		address = addressInterface.(string)
	}

	itemsInterface := ls.KvStorage.Get("items")
	var items []string
	if itemsInterface != nil {
		items = itemsInterface.([]string)
	}


	botService := ls.KvStorage.Get("botService").(string)
	groupId, _ := strconv.Atoi(ls.KvStorage.Get("groupId").(string))
	userId, _ := strconv.Atoi(ls.KvStorage.Get("userId").(string))

	if address == "" {
		answerText := "Вы не ввели адрес доставки"
		SendMessageFunc(ctx, cm, ls, answerText)
		return s.WaitInputStateName, true
	} else if len(items) == 0 {
		answerText := "Ваш заказ пустой"
		SendMessageFunc(ctx, cm, ls, answerText)
		return s.WaitInputStateName, true
	}

	text := `Адрес: %s
Товары:`

	text = fmt.Sprintf(text, address)

	for i, item := range items {
		text += fmt.Sprintf("\n%d: %s", i + 1, item)
	}

	err := cm.ConsumerMessageSender.SendMessage(&consumer_messages.MessageToConsumer{
		Uuid: "",
		MsgLocation: &common.MsgLocation{
			Service: botService,
			GroupId: int64(groupId),
			UserId:  int64(userId),
		},
		Text:  text,
		Reply: nil,
	})
	if err != nil {
		fmt.Println(err)
	}

	return s.ConfirmOrderStateName, false
}
