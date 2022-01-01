package vkApi

import (
	"strconv"
	"vk_shop_bot/vkApi/VkApiServer"
	"vk_shop_bot/vkApi/VkLongPullServer"
)

type VkBot struct {
	AccessToken string
	GroupId     int

	CurrentTs int

	VkLongPullServer *VkLongPullServer.VkLongPullServer
	VkApiServer      *VkApiServer.VkApiServer
}

func (vkBot *VkBot) Init(token string, groupId int, vkApiServer *VkApiServer.VkApiServer) error {
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

	vkBot.VkLongPullServer = &VkLongPullServer.VkLongPullServer{
		BaseUrl: resp.Response.Server,
		Key:     resp.Response.Key,
	}
	return
}

func (vkBot *VkBot) GetUpdates() (objects []VkLongPullServer.UpdateObject, err error) {
	var resp VkLongPullServer.ACheckResponse
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
