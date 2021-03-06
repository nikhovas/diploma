package apiServer

import "strconv"

type mMessagesSendResponse struct {
	response int `json:"response"`
}

func (vkApiServer *VkApiServer) MMessagesSend(
	accessToken string,
	groupId int,
	peerId int,
	message string,
	uniqueId uint64,
	replyMessageId *int,
) (int, error) {
	var resp mMessagesSendResponse

	query := map[string]string{
		"access_token": accessToken,
		"message":      message,
		"peer_id":      strconv.Itoa(peerId),
		"group_id":     strconv.Itoa(groupId),
		"random_id":    strconv.Itoa(int(uniqueId)),
	}
	if replyMessageId != nil {
		query["reply_to"] = strconv.Itoa(*replyMessageId)
	}

	err := vkApiServer.SendMethodRequest("messages", "send", query, &resp)
	return resp.response, err
}
