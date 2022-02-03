package bots

import (
	apiServer "github.com/nikhovas/diploma/go/lib/vk/api_server"
	longPullServer "github.com/nikhovas/diploma/go/lib/vk/long_pull_server"
	"strconv"
)

type VkBot struct {
	AccessToken string
	GroupId     int

	CurrentTs int

	VkLongPullServer *longPullServer.VkLongPullServer
	VkApiServer      *apiServer.VkApiServer
}

func (vkBot *VkBot) Init(token string, groupId int, vkApiServer *apiServer.VkApiServer) error {
	vkBot.AccessToken = token
	vkBot.GroupId = groupId
	vkBot.VkApiServer = vkApiServer
	return nil
}

func (vkBot *VkBot) Authorize() (err error) {
	resp, err := vkBot.VkApiServer.MGroupsGetLongPollServer(vkBot.AccessToken, vkBot.GroupId)
	if err != nil {
		return
	}

	var currentTs int
	currentTs, err = strconv.Atoi(resp.Response.Ts)
	if err != nil {
		return
	}
	vkBot.CurrentTs = currentTs

	vkBot.VkLongPullServer = &longPullServer.VkLongPullServer{
		BaseUrl: resp.Response.Server,
		Key:     resp.Response.Key,
	}
	return
}

func (vkBot *VkBot) GetUpdates() (objects []longPullServer.UpdateObject, err error) {
	var resp longPullServer.ACheckResponse
	resp, err = vkBot.VkLongPullServer.ACheck(vkBot.CurrentTs, 25)
	objects = resp.Updates

	if err != nil {
		return
	}

	vkBot.CurrentTs, err = strconv.Atoi(resp.Ts)
	return
}

func (vkBot *VkBot) SendMessage(userId int, text string) (int, error) {
	return vkBot.VkApiServer.MMessagesSend(vkBot.AccessToken, vkBot.GroupId, userId, text)
}
