package utils

import "fmt"

func GetDataBasePath(botService string, groupId string, userId string) string {
	return fmt.Sprintf("messages/%s/%s/%s", botService, groupId, userId)
}

type MessageInfo struct {
	Text string
	Id   int
}
