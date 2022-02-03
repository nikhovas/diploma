package apiServer

import "strconv"

func (vkApiServer *VkApiServer) MMarketGet(accessToken string, groupId int, peerId int, message string) (int, error) {
	var resp mMessagesSendResponse
	err := vkApiServer.SendMethodRequest(
		"market",
		"get",
		map[string]string{
			"access_token": accessToken,
			"message":      message,
			"peer_id":      strconv.Itoa(peerId),
			"group_id":     strconv.Itoa(groupId),
		},
		&resp,
	)
	return resp.response, err
}
