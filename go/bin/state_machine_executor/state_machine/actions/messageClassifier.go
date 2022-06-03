package actions

//import (
//	"context"
//	"fmt"
//	"state_machine_executor/application"
//	"state_machine_executor/state_machine/localStorage"
//	"state_machine_executor/wordTools"
//	"strings"
//)
//
//type MessageClassifier struct {
//	GenericAction
//}
//
//func NewMessageClassifier(genericAction *GenericAction) ActionInterface {
//	return &MessageClassifier{GenericAction: *genericAction}
//}
//
//const (
//	AddressPrefix = "/адрес"
//	ObjectPrefix  = "/товар"
//	OrderPrefix  = "/заказ"
//)
//
//func (a *MessageClassifier) Run(ctx context.Context, app *application.Application, storage *localStorage.Storage) {
//	currentState := storage.KvStorage.Get("currentState").(string)
//	newState := ""
//
//	if currentState == "waitNewCmd" {
//		newState = FromWaitNewCmd(ctx, app, storage)
//	} else if currentState == "shouldConfirmOrder" {
//		newState = FromShouldConfirmOrder(ctx, app, storage)
//	} else {
//
//	}
//
//	storage.KvStorage.Set("currentState", newState)
//}
//
//func FromStart(ctx context.Context, app *application.Application, storage *localStorage.Storage) string {
//	message := `Это будет приветствие`
//
//}
//
//func ToStart(ctx context.Context, app *application.Application, storage *localStorage.Storage) string {
//
//}
//
//func FromWaitNewCmd(ctx context.Context, app *application.Application, storage *localStorage.Storage) string {
//	message := storage.KvStorage.Get("message").(string)
//
//	if strings.HasPrefix(message, AddressPrefix) {
//		text := message[len(AddressPrefix):]
//		text = strings.Trim(text, " ")
//		SetAddress(ctx, app, storage, text)
//	} else if strings.HasPrefix(message, ObjectPrefix) {
//		text := message[len(ObjectPrefix):]
//		text = strings.Trim(text, " ")
//		SetNewObject(ctx, app, storage, text)
//	} else if strings.HasPrefix(message, OrderPrefix) {
//		text := message[len(ObjectPrefix):]
//		text = strings.Trim(text, " ")
//		return OrderStart(ctx, app, storage)
//	} else {
//		AnswerQuestionInternalFunc(ctx, app, storage)
//	}
//
//	return ToWaitNewCmd(ctx, app, storage)
//}
//
//func ToWaitNewCmd(ctx context.Context, app *application.Application, storage *localStorage.Storage) string {
//	return "waitNewCmd"
//}
//
//func SetAddress(ctx context.Context, app *application.Application, storage *localStorage.Storage, address string) {
//	storage.KvStorage.Set("address", address)
//}
//
//func SetNewObject(ctx context.Context, app *application.Application, storage *localStorage.Storage, object string) {
//	items := storage.KvStorage.Get("items").([]string)
//	items = append(items, object)
//	storage.KvStorage.Set("items", object)
//}
//
//func OrderStart(ctx context.Context, app *application.Application, storage *localStorage.Storage) string {
//	if !CheckOrderCorrectness(ctx, app, storage) {
//		return ToWaitNewCmd(ctx, app, storage)
//	}
//
//	address := storage.KvStorage.Get("address").(string)
//	items := storage.KvStorage.Get("items").([]string)
//
//	text := `Подтвердите ваш заказ:
//Адрес: %s
//Товары:`
//
//	text = fmt.Sprintf(text, address)
//
//	for i, item := range items {
//		text += fmt.Sprintf("\n%d: %s", i, item)
//	}
//
//	return ToShouldConfirmOrder(ctx, app, storage)
//}
//
//func CheckOrderCorrectness(ctx context.Context, app *application.Application, storage *localStorage.Storage) bool {
//	address := storage.KvStorage.Get("address").(string)
//	items := storage.KvStorage.Get("items").([]string)
//
//	if address == "" {
//		answerText := "Вы не ввели адрес доставки"
//		SendMessageFunc(ctx, app, storage, map[string]string{"text": answerText}, Returns{})
//		return false
//	} else if len(items) == 0 {
//		answerText := "Ваш заказ пустой"
//		SendMessageFunc(ctx, app, storage, map[string]string{"text": answerText}, Returns{})
//		return false
//	}
//
//	return true
//}
//
//func ToShouldConfirmOrder(ctx context.Context, app *application.Application, storage *localStorage.Storage) string {
//	answer := "Подтвердите корректность заказа"
//	SendMessageFunc(ctx, app, storage, map[string]string{"text": answer}, Returns{})
//	return "shouldConfirmOrder"
//}
//
//func FromShouldConfirmOrder(ctx context.Context, app *application.Application, storage *localStorage.Storage) string {
//	message := storage.KvStorage.Get("message").(string)
//	result, detected := wordTools.BoolDetect(message)
//
//	if !detected {
//		return ToShouldConfirmOrder(ctx, app, storage)
//	} else if result {
//		answer := "Ваш заказ отправлен операторам"
//		SendMessageFunc(ctx, app, storage, map[string]string{"text": answer}, Returns{})
//	} else {
//		answer := "Хорошо, измените ваш заказ"
//		SendMessageFunc(ctx, app, storage, map[string]string{"text": answer}, Returns{})
//		return ToWaitNewCmd(ctx, app, storage)
//	}
//}
//
////func SetNewObject(ctx context.Context, app *application.Application, storage *localStorage.Storage) {
////
////}
//
////func sendQuestionToStaff(
////	ctx context.Context,
////	app *application.Application,
////	question string,
////	shopId int,
////	service string,
////	groupId int,
////	userId int,
////	messageId int,
////) {
////	_, err := app.ControlClient.AddQuestion(ctx, &ctrl.AddQuestionRequest{
////		Key: &ctrl.ShopKey{
////			Key: &ctrl.ShopKey_Common{
////				Common: &ctrl.CommonShopKey{
////					CommonKey: &ctrl.CommonShopKey_Id{
////						Id: int64(shopId),
////					},
////				},
////			},
////		},
////		Info: &common.WaitingQuesionInformation{
////			MsgLocation: &common.MsgLocation{
////				Service: service,
////				GroupId: int64(groupId),
////				UserId:  int64(userId),
////			},
////			Question:      question,
////			QuestionMsgId: int64(messageId),
////		},
////	})
////	if err != nil {
////		fmt.Println(err)
////		return
////	}
////}
