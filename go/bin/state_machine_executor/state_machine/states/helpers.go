package states

import (
	"context"
	"fmt"
	"github.com/nikhovas/diploma/go/lib/proto/common"
	"github.com/nikhovas/diploma/go/lib/proto/consumer_messages"
	"state_machine_executor/coremodules"
	"state_machine_executor/state_machine/actions"
	"state_machine_executor/state_machine/localStorage"
	"state_machine_executor/utils"
	"strconv"
)

func SendMessageFunc(
	ctx context.Context,
	cm *coremodules.CoreModules,
	storage *localStorage.Storage,
	message string,
) {
	botService := storage.KvStorage.Get("botService").(string)
	groupId, _ := strconv.Atoi(storage.KvStorage.Get("groupId").(string))
	userId, _ := strconv.Atoi(storage.KvStorage.Get("userId").(string))

	err := cm.ConsumerMessageSender.SendMessage(&consumer_messages.MessageToConsumer{
		Uuid: "",
		MsgLocation: &common.MsgLocation{
			Service: botService,
			GroupId: int64(groupId),
			UserId:  int64(userId),
		},
		Text:  message,
		Reply: nil,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func BreakIfNoMessages(ls *localStorage.Storage, messageToStack bool) string {
	if ls.MessageDeque.Len() == 0 {
		return ls.KvStorage.Get("state").(string)
	}

	data := ls.MessageDeque.PopFront()
	msg := data.(*utils.MessageInfo)
	ls.KvStorage.Set("message", msg.Text)
	ls.KvStorage.Set("messageId", strconv.Itoa(msg.Id))

	if messageToStack {
		actions.MessageToStackFunc(ls, nil, nil)
	}

	return ""
}
