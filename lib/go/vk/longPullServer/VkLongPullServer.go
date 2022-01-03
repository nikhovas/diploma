package longPullServer

import "vk/utils"

type VkLongPullServer struct {
	BaseUrl string
	Key     string
}

func (vkLongPullServer *VkLongPullServer) SendActionRequest(action string, query map[string]string, response interface{}) error {
	query["key"] = vkLongPullServer.Key
	query["act"] = action
	return utils.SendGetRequest(
		vkLongPullServer.BaseUrl, query, &response,
	)
}
