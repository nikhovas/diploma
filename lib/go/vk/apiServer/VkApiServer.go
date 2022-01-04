package apiServer

import "github.com/nikhovas/diploma/lib/go/vk/utils"

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
		vkApiServer.GetMethodUrl(methodGroup, methodName), query, &response,
	)
}
