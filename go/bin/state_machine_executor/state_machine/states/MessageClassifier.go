package states

import (
	"context"
	"state_machine_executor/coremodules"
	"state_machine_executor/state_machine/localStorage"
	"strings"
)

type MessageClassifier struct {
	CheckOrderStateName string
	QuestionAnswerStateName string
	WaitInputStateName string
}

const (
	AddressPrefix = "/адрес"
	ObjectPrefix  = "/товар"
	OrderPrefix  = "/заказ"
)

func (s *MessageClassifier) Process(ctx context.Context, cm *coremodules.CoreModules, ls *localStorage.Storage) (string, bool) {
	//if breakState := BreakIfNoMessages(ls, false); breakState != "" {
	//	return breakState, true
	//}

	message := ls.KvStorage.Get("message").(string)

	if strings.HasPrefix(message, AddressPrefix) {
		text := message[len(AddressPrefix):]
		text = strings.Trim(text, " ")
		SetAddress(ctx, cm, ls, text)
		adr := "Адрес " + text + " установлен"
		SendMessageFunc(ctx, cm, ls, adr)
	} else if strings.HasPrefix(message, ObjectPrefix) {
		text := message[len(ObjectPrefix):]
		text = strings.Trim(text, " ")
		SetNewObject(ctx, cm, ls, text)
		msg := "Товар " + text + " добавлен"
		SendMessageFunc(ctx, cm, ls, msg)
	} else if strings.HasPrefix(message, OrderPrefix) {
		text := message[len(ObjectPrefix):]
		text = strings.Trim(text, " ")
		return s.CheckOrderStateName, false
	} else {
		return s.QuestionAnswerStateName, false
	}

	return s.WaitInputStateName, true
}

func SetAddress(ctx context.Context, cm *coremodules.CoreModules, storage *localStorage.Storage, address string) {
	storage.KvStorage.Set("address", address)
}

func SetNewObject(ctx context.Context, cm *coremodules.CoreModules, storage *localStorage.Storage, object string) {
	items := storage.KvStorage.Get("items").([]string)
	items = append(items, object)
	storage.KvStorage.Set("items", items)
}