package actions

var ActionBuilders = map[string]func(genericAction *GenericAction) ActionInterface{
	"sendMessage":    NewSendMessage,
	"answerQuestion": NewAnswerQuestion,
	"messageToStack": NewMessageToStack,
	//"classifyMessage": NewMessageClassifier,
}
