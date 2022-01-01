package VkApiServer

import "strconv"

type mMessagesSendResponse struct {
	response int `json:"response"`
}

func (vkApiServer *VkApiServer) MMessagesSend(accessToken string, groupId int, peerId int, message string) (int, error) {
	var resp mMessagesSendResponse
	err := vkApiServer.SendMethodRequest(
		"messages",
		"send",
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
