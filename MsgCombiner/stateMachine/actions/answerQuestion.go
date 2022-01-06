package actions

import (
	"MsgCombiner/stateMachine/localStorage"
	"fmt"
)

type AnswerQuestion struct {
	GenericAction
}

func NewAnswerQuestion(genericAction *GenericAction) ActionInterface {
	return &AnswerQuestion{GenericAction: *genericAction}
}

func (a *AnswerQuestion) Run(storage *localStorage.Storage) {
	message := storage.KvStorage.Get("message").(string)
	previousQuestions := storage.KvStorage.Get("previousQuestions").([]string)
	previousQuestions = append(previousQuestions, message)
	storage.KvStorage.Set("previousQuestions", previousQuestions)
	fmt.Println("This will be answer for ", message)
}
