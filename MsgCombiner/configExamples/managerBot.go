package configExamples

import (
	"MsgCombiner/configs/external"
	v0 "MsgCombiner/configs/v0"
)

var ManagerBotV0 = external.External{
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
							Name:      "managerBot",
							Arguments: map[string]string{},
							Return:    map[string]string{},
						},
					},
					Next: "initial",
				},
			},
		},
		Data: map[string]v0.DataElement{
			"previousQuestions": {
				Memory: "long",
				Type:   "stringList",
			},
		},
	},
}

var ManagerBotInternal = ManagerBotV0.V0.ToInternal()
