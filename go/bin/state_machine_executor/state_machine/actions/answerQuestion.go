package actions

import (
	"context"
	"fmt"
	qw "github.com/nikhovas/diploma/go/lib/proto/question_worker"
	"state_machine_executor/application"
	"state_machine_executor/state_machine/localStorage"
)

type AnswerQuestion struct {
	GenericAction
}

func NewAnswerQuestion(genericAction *GenericAction) ActionInterface {
	return &AnswerQuestion{GenericAction: *genericAction}
}

func (a *AnswerQuestion) Run(ctx context.Context, application *application.Application, storage *localStorage.Storage) {
	message := storage.KvStorage.Get("message").(string)
	previousQuestions := storage.KvStorage.Get("previousQuestions").([]string)
	previousQuestions = append(previousQuestions, message)
	storage.KvStorage.Set("previousQuestions", previousQuestions)

	shopId := storage.KvStorage.Get("groupId").(int)

	answer, err := application.QwClient.GetQuestionAnswer(ctx, &qw.GetQuestionAnswerRequest{
		Question:          message,
		PreviousQuestions: previousQuestions,
		BasePath:          fmt.Sprintf("/bots/data/internal/%d/qa", shopId),
	})
	if err != nil {
		return
	}

	questionAnswer := answer.Answer

	fmt.Println("This will be answer for ", message, questionAnswer)
}
