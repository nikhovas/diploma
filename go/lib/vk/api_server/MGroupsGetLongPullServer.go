package apiServer

import "strconv"

type VkGetLongPullServerResponse struct {
	Response struct {
		Key    string `json:"key"`
		Server string `json:"server"`
		Ts     string `json:"ts"`
	} `json:"response"`
}

func (vkApiServer *VkApiServer) MGroupsGetLongPollServer(accessToken string, groupId int) (VkGetLongPullServerResponse, error) {
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
