package utils

import "fmt"

func GetDataBasePath(botService string, groupId string, userId string) string {
	return fmt.Sprintf("messages/%s/%s/%s", botService, groupId, userId)
}

func GetMessagesKey(botService string, groupId string, userId string) string {
	return fmt.Sprintf("%s/messages", GetDataBasePath(botService, groupId, userId))
}

func GetMessagePath(botService string, groupId string, userId string, messageTs int) string {
	return fmt.Sprintf("%s/%d", GetMessagesKey(botService, groupId, userId), messageTs)
}

func GetStartReadFromKey(botService string, groupId string, userId string) string {
	return fmt.Sprintf("%s/start-read-from", GetDataBasePath(botService, groupId, userId))
}

func GetLockKey(botService string, groupId string, userId string) string {
	return fmt.Sprintf("%s/lock", GetDataBasePath(botService, groupId, userId))
}
