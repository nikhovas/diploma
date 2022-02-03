package configExamples

import (
	"state_machine_executor/configs/external"
	"state_machine_executor/configs/v0"
)

var QuestionsOnlyV0 = external.External{
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
					Next:           "fromWaitInputActions",
				},
			},
			"fromWaitInputActions": {
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
				Default: "hello!",
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
		},
	},
}

var QuestionsOnlyInternal = QuestionsOnlyV0.V0.ToInternal()
