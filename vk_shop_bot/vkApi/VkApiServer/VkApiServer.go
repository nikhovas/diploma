package vkApi

import (
	"strconv"
	"vk_shop_bot/utils"
)

type VkGetLongPullServerResponse struct {
	Response struct {
		Key    string `json:"key"`
		Server string `json:"server"`
		Ts     string `json:"ts"`
	} `json:"response"`
}

type VkApiServer struct {
	Host    string
	Version string
}

func (vkApiServer *VkApiServer) GetMethodUrl(methodGroup string, methodName string) string {
	return vkApiServer.Host + "/method/" + methodGroup + "." + methodName
}

func (vkApiServer *VkApiServer) SendMethodRequest(methodGroup string, methodName string, query map[string]string, response interface{}) error {
	query["v"] = vkApiServer.Version
	return utils.SendGetRequest(
		vkApiServer.GetMethodUrl("groups", "getLongPollServer"), query, &response,
	)
}

func (vkApiServer *VkApiServer) GroupsGetLongPollServer(accessToken string, groupId int) (VkGetLongPullServerResponse, error) {
	var resp VkGetLongPullServerResponse
	err := vkApiServer.SendMethodRequest(
		"groups",
		"getLongPollServer",
		map[string]string{
			"access_token": accessToken,
			"group_id":     strconv.Itoa(groupId),
		},
		&resp,
	)
	return resp, err
}
