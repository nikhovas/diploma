package bots

import (
	"context"
	"github.com/nikhovas/diploma/go/lib/utils/distfs/bots"
	"github.com/nikhovas/diploma/go/lib/utils/distfs/bots/meta_service_name"
	"github.com/nikhovas/diploma/go/lib/utils/distvars"
	apiServer "github.com/nikhovas/diploma/go/lib/vk/api_server"
	longPullServer "github.com/nikhovas/diploma/go/lib/vk/long_pull_server"
	"math/big"
)

type DistributedBot struct {
	VkBot
	//app             *application.Application
	BotDistFsBase   *meta_service_name.MetaGroupId
	tokenDistFs     *distvars.RedisString
	messageIdDistFs *distvars.ConsulInt
	newMsgUniqueId  *distvars.RedisCounter
}

func (bot *DistributedBot) GetTokenValue() (string, error) {
	token, err := bot.tokenDistFs.Get(context.Background())
	if err != nil {
		return token, err
	} else if token == "" {
		return "", big.ErrNaN{}
	}
	return token, nil
}

func (bot *DistributedBot) Init(
	groupId int,
	vkApiServer *apiServer.VkApiServer,
	vkServiceDistFs *bots.MetaServiceName,
) error {
	bot.GroupId = groupId
	var err error

	bot.BotDistFsBase = vkServiceDistFs.MetaCdGroupId(groupId)
	bot.tokenDistFs = bot.BotDistFsBase.CdCommon().CdToken()
	bot.messageIdDistFs = bot.BotDistFsBase.CdCommon().CdMessageId()
	bot.newMsgUniqueId = bot.BotDistFsBase.CdCommon().CdNewMsgUniqueId()

	token, err := bot.GetTokenValue()
	if err != nil {
		return err
	}

	return bot.VkBot.Init(token, groupId, vkApiServer)
}

func (bot *DistributedBot) Authorize() error {
	err := bot.VkBot.Authorize()
	if err != nil {
		return err
	}

	bot.CurrentTs, _, err = bot.messageIdDistFs.SwapIfGreater(context.Background(), bot.CurrentTs)
	if err != nil {
		return err
	}

	return nil
}

func (bot *DistributedBot) GetUpdates() (updTs int, updates []longPullServer.UpdateObject, err error) {
	var startedFromTs int
	startedFromTs, _, err = bot.messageIdDistFs.Get(context.Background())
	if err != nil {
		return 0, []longPullServer.UpdateObject{}, err
	}

	if bot.VkBot.CurrentTs < startedFromTs {
		bot.VkBot.CurrentTs = startedFromTs
	}

	updates, err = bot.VkBot.GetUpdates()
	realStartFromTs := bot.CurrentTs - len(updates)

	_, shouldReadFrom, err := bot.messageIdDistFs.SwapIfGreater(context.Background(), bot.CurrentTs)
	if bot.CurrentTs <= shouldReadFrom {
		return 0, []longPullServer.UpdateObject{}, nil
	} else {
		return realStartFromTs, updates[(realStartFromTs - shouldReadFrom):], nil
	}
}

func (bot *DistributedBot) SendMessage(userId int, text string, replyTo *int) (int, error) {
	//sendingMsgId := atomic.AddUint64(&vkBot.currentMessageId, 1)
	incr, err := bot.newMsgUniqueId.Incr(context.Background())
	if err != nil {
		return 0, err
	}
	return bot.VkBot.VkApiServer.MMessagesSend(bot.VkBot.AccessToken, bot.VkBot.GroupId, userId, text, uint64(incr), replyTo)
}
