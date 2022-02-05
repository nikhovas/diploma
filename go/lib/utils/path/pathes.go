package path

import "fmt"

const BotCommonDataService = "common"

func GetConsumerBotCommonDataPath(service string) string {
	return fmt.Sprintf("bots/%s/common", service)
}

func GetBotCommonDataPath(service string, groupId interface{}) string {
	return fmt.Sprintf("bots/%s/%v/common", service, groupId)
}

func GetBotUserDataPath(service string, groupId interface{}, userId interface{}) string {
	return fmt.Sprintf("bots/%s/%v/consumers/%v", service, groupId, userId)
}

func GetBotActionsConsulPath(service string, groupId interface{}, userId interface{}) string {
	return fmt.Sprintf("%s/actions", GetBotUserDataPath(service, groupId, userId))
}

func GetBotMessagesConsulPath(service string, groupId interface{}, userId interface{}) string {
	return fmt.Sprintf("%s/messages", GetBotUserDataPath(service, groupId, userId))
}

func GetBotMessageConsulPath(service string, groupId interface{}, userId interface{}, messageId interface{}) string {
	return fmt.Sprintf("%s/%v", GetBotMessagesConsulPath(service, groupId, userId), messageId)
}

//func GetBotTokenRedisPath(service string, groupId interface{}) string {
//	return fmt.Sprintf("%s/token", GetBotCommonDataPath(service, groupId))
//}

func GetBotWaitingQuestionsRedisPath(groupId interface{}) string {
	return fmt.Sprintf("%s/waiting-questions", GetBotCommonDataPath(BotCommonDataService, groupId))
}

func GetBotQuestionAnswersRedisPath(groupId interface{}) string {
	return fmt.Sprintf("%s/question-answers", GetBotCommonDataPath(BotCommonDataService, groupId))
}

//func GetBotMessageIdConsulPath(service string, groupId interface{}) string {
//	return fmt.Sprintf("%s/message-id", GetBotCommonDataPath(service, groupId))
//}
