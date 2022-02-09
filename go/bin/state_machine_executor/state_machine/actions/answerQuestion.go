package actions

import (
	"context"
	"fmt"
	"github.com/nikhovas/diploma/go/lib/proto/common"
	ctrl "github.com/nikhovas/diploma/go/lib/proto/controller"
	qw "github.com/nikhovas/diploma/go/lib/proto/question_worker"
	"github.com/nikhovas/diploma/go/lib/utils/distfs"
	"state_machine_executor/application"
	"state_machine_executor/state_machine/localStorage"
	"strconv"
)

type AnswerQuestion struct {
	GenericAction
}

func NewAnswerQuestion(genericAction *GenericAction) ActionInterface {
	return &AnswerQuestion{GenericAction: *genericAction}
}

func (a *AnswerQuestion) Run(ctx context.Context, app *application.Application, storage *localStorage.Storage) {
	message := storage.KvStorage.Get("message").(string)
	previousQuestions := storage.KvStorage.Get("previousQuestions").([]string)
	previousQuestions = append(previousQuestions, message)
	storage.KvStorage.Set("previousQuestions", previousQuestions)

	shopId, _ := strconv.Atoi(storage.KvStorage.Get("shopId").(string))

	path := distfs.NewRoot(app.RedisClient, app.ConsulClient).CdCommon().MetaCdShopId(shopId).CdQa().Path

	answer, err := app.QwClient.GetQuestionAnswer(ctx, &qw.GetQuestionAnswerRequest{
		Question:          message,
		PreviousQuestions: previousQuestions,
		BasePath:          path,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	questionAnswer := answer.Answer

	if answer.AnswerDistance > 0.7 {
		botService := storage.KvStorage.Get("botService").(string)
		groupId, _ := strconv.Atoi(storage.KvStorage.Get("groupId").(string))
		userId, _ := strconv.Atoi(storage.KvStorage.Get("userId").(string))
		messageId, _ := strconv.Atoi(storage.KvStorage.Get("messageId").(string))
		sendQuestionToStaff(ctx, app, message, shopId, botService, groupId, userId, messageId)

		toUserText := "На ваш вопрос на данный момент нет ответа. Вопрос отправлен оператору."
		SendMessageFunc(ctx, app, storage, map[string]string{"text": toUserText}, Returns{})
	} else {
		SendMessageFunc(ctx, app, storage, map[string]string{"text": questionAnswer}, Returns{})
	}

}

func sendQuestionToStaff(
	ctx context.Context,
	app *application.Application,
	question string,
	shopId int,
	service string,
	groupId int,
	userId int,
	messageId int,
) {
	_, err := app.ControlClient.AddQuestion(ctx, &ctrl.AddQuestionRequest{
		Key: &ctrl.ShopKey{
			Key: &ctrl.ShopKey_Common{
				Common: &ctrl.CommonShopKey{
					CommonKey: &ctrl.CommonShopKey_Id{
						Id: int64(shopId),
					},
				},
			},
		},
		Info: &common.WaitingQuesionInformation{
			MsgLocation: &common.MsgLocation{
				Service: service,
				GroupId: int64(groupId),
				UserId:  int64(userId),
			},
			Question:      question,
			QuestionMsgId: int64(messageId),
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
