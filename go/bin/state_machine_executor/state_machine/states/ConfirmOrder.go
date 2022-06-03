package states

import (
	"context"
	cp "github.com/nikhovas/diploma/go/lib/proto/common"
	"log"
	"state_machine_executor/coremodules"
	"state_machine_executor/state_machine/localStorage"
	"state_machine_executor/wordTools"
	"strconv"
)

type ConfirmOrder struct {
	SuccessOrderState string
	ChangeOrderState string
	RetryConfirmState string
}

func (s *ConfirmOrder) Process(ctx context.Context, cm *coremodules.CoreModules, ls *localStorage.Storage) (string, bool) {
	if breakState := BreakIfNoMessages(ls, false); breakState != "" {
		return breakState, true
	}

	userId, _ := strconv.Atoi(ls.KvStorage.Get("userId").(string))
	message := ls.KvStorage.Get("message").(string)
	result, detected := wordTools.BoolDetect(message)

	if !detected {
		return s.RetryConfirmState, false
	} else if result {
		SendMessageFunc(ctx, cm, ls, "Ваш заказ отправлен операторам")

		address := ls.KvStorage.Get("address").(string)
		items := ls.KvStorage.Get("items").([]string)

		shopId, _ := strconv.Atoi(ls.KvStorage.Get("shopId").(string))

		order := cp.Order{
			Id:         0,
			ShopId: int64(shopId),
			SelfPickup: false,
			Address:    address,
			Item:       items,
			UserId:     "https://vk.com/id" + strconv.Itoa(userId),
		}

		err := cm.OrdersQueueWriter.SendMessage(&order)
		if err != nil {
			log.Println(err)
		}

		ls.KvStorage.Set("items", []string{})
		ls.KvStorage.Set("address", "")

		return s.SuccessOrderState, false
	} else {
		SendMessageFunc(ctx, cm, ls, "Хорошо, измените ваш заказ")
		return s.ChangeOrderState, false
	}
}
