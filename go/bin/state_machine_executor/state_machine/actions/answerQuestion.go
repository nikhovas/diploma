package actions

import (
	"context"
	"fmt"
	"github.com/nikhovas/diploma/go/lib/proto/common"
	ctrl "github.com/nikhovas/diploma/go/lib/proto/controller"
	qw "github.com/nikhovas/diploma/go/lib/proto/question_worker"
	"github.com/nikhovas/diploma/go/lib/utils/distfs"
	"state_machine_executor/coremodules"
	"state_machine_executor/state_machine/localStorage"
	"strconv"
)

type AnswerQuestion struct {
	GenericAction
}

func NewAnswerQuestion(genericAction *GenericAction) ActionInterface {
	return &AnswerQuestion{GenericAction: *genericAction}
}

func (a *AnswerQuestion) Run(ctx context.Context, cm *coremodules.CoreModules, storage *localStorage.Storage) {
	AnswerQuestionInternalFunc(ctx, cm, storage)
}

func AnswerQuestionInternalFunc(ctx context.Context, cm *coremodules.CoreModules, storage *localStorage.Storage) {
	message := storage.KvStorage.Get("message").(string)
	previousQuestions := storage.KvStorage.Get("previousQuestions").([]string)
	previousQuestions = append(previousQuestions, message)
	storage.KvStorage.Set("previousQuestions", previousQuestions)

	shopId, _ := strconv.Atoi(storage.KvStorage.Get("shopId").(string))

	path := distfs.NewRoot(cm.RedisClient, cm.ConsulClient).CdCommon().MetaCdShopId(shopId).CdQa().Path

	answer, err := cm.QwClient.GetQuestionAnswer(ctx, &qw.GetQuestionAnswerRequest{
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
		sendQuestionToStaff(ctx, cm, message, shopId, botService, groupId, userId, messageId)

		toUserText := "На ваш вопрос на данный момент нет ответа. Вопрос отправлен оператору."
		storage.KvStorage.Set("metaField", toUserText)
		SendMessageFunc(ctx, cm, storage, map[string]string{"text": "metaField"}, Returns{})
	} else {
		storage.KvStorage.Set("metaField", questionAnswer)
		SendMessageFunc(ctx, cm, storage, map[string]string{"text": "metaField"}, Returns{})
	}
}

func sendQuestionToStaff(
	ctx context.Context,
	cm *coremodules.CoreModules,
	question string,
	shopId int,
	service string,
	groupId int,
	userId int,
	messageId int,
) {
	_, err := cm.ControlClient.AddQuestion(ctx, &ctrl.AddQuestionRequest{
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
