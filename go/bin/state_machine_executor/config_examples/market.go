package configExamples

import (
	"state_machine_executor/configs/external"
	v0 "state_machine_executor/configs/v0"
	"state_machine_executor/state_machine/states"
)

var MarketV0 = external.External{
	V0: &v0.Model{
		StartState: "initial",
		States: map[string]v0.State{
			"initial": {
				MessageWaitStateData: &v0.MessageWaitState{
					MessageToStack: true,
					Next:           "fromInitialActions",
				},
			},
			"fromInitialActions": {
				ActionStateData: &v0.ActionState{
					Actions: []v0.Action{
						{
							Name:      "sendMessage",
							Arguments: map[string]string{"text": "greetingText"},
							Return:    map[string]string{},
						},
					},
					Next: "waitInputState",
				},
			},
			"waitInputState": {
				MessageWaitStateData: &v0.MessageWaitState{
					MessageToStack: false,
					Next:           "messageClassifier",
				},
			},
			"messageClassifier": {
				MessageClassifierStateData: &states.MessageClassifier{
					CheckOrderStateName: "checkOrder",
					QuestionAnswerStateName: "questionAnswer",
					WaitInputStateName: "waitInputState",
				},
			},
			"confirmOrder": {
				ConfirmOrderStateData: &states.ConfirmOrder{
					SuccessOrderState: "initial",
					ChangeOrderState: "waitInputState",
					RetryConfirmState: "beforeConfirmOrder",
				},
			},
			"beforeConfirmOrder": {
				ActionStateData: &v0.ActionState{
					Actions: []v0.Action{
						{
							Name:      "sendMessage",
							Arguments: map[string]string{"text": "checkOrderMessage"},
							Return:    map[string]string{},
						},
					},
					Next:    "confirmOrder",
				},
			},
			"checkOrder": {
				CheckOrderStateData: &states.CheckOrder{
					ConfirmOrderStateName: "beforeConfirmOrder",
					WaitInputStateName:    "waitInputState",
				},
			},
			"questionAnswer": {
				ActionStateData: &v0.ActionState{
					Actions: []v0.Action{
						{
							Name:      "answerQuestion",
							Arguments: map[string]string{},
							Return:    map[string]string{},
						},
					},
					Next: "waitInputState",
				},
			},
		},
		Data: map[string]v0.DataElement{
			"greetingText": {
				Memory:  "const",
				Type:    "string",
				Default: "Привет! Я бот магазина. Мне можно задавать вопросы, а также добавлять позиции товаров с помощью команды /товар, добавлять адрес при помощи команды /адрес, оформлять заказ при помощи команды /заказ.",
			},
			"checkOrderMessage": {
				Memory:  "const",
				Type:    "string",
				Default: "Подтвердите корректность заказа",
			},
			"questionAnswerText": {
				Memory: "short",
				Type:   "string",
			},
			"questionAnswerProbability": {
				Memory: "short",
				Type:   "float",
			},
			"previousQuestions": {
				Memory: "long",
				Type:   "stringList",
			},
			"selfPickupAddresses": {
				Memory: "long",
				Type:   "stringList",
			},
			"address": {
				Memory: "long",
				Type:   "string",
			},
			"items": {
				Memory: "long",
				Type:   "stringList",
			},
			"metaField": {
				Memory: "short",
				Type:   "string",
			},
		},
	},
}

var MarketInternal = MarketV0.V0.ToInternal()
