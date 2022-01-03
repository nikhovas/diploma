package messageProcess

import (
	"MsgCombiner/utils"
	"bytes"
	"github.com/golang/protobuf/jsonpb"
	"github.com/hashicorp/consul/api"
	UserActions "github.com/nikhovas/diploma/proto/data/userActions"
	"strconv"
	"time"
)

func (aep *ActionEventProcessor) GetAllUserActions(messagesKey string) (newestTs int, userActionList []*UserActions.UserAction, err error) {
	startReadFromKey := utils.GetStartReadFromKey(vkShopBot, aep.ActionEvent.BotId, aep.ActionEvent.UserId)

	startFrom := 0 // TODO: implement
	var startReadFromKv *api.KVPair
	startReadFromKv, _, err = aep.Application.ConsulClient.KV().Get(startReadFromKey, nil)
	if startReadFromKv != nil {
		startFrom, _ = strconv.Atoi(string(startReadFromKv.Value))
	}

	newestTs = 0

	listLength := 0

	var list api.KVPairs

	for {
		if list, _, err = aep.Application.ConsulClient.KV().List(messagesKey, nil); err != nil {
			return
		}

		newListLength := len(list)
		if newListLength == listLength {
			break
		}

		listLength = newListLength

		for _, kv := range list {
			//key := kv.Key

			var userAction UserActions.UserAction
			if err = jsonpb.Unmarshal(bytes.NewReader(kv.Value), &userAction); err != nil {
				return
			}
			shortKey := "2345" // TODO: here
			var ts int
			if ts, err = strconv.Atoi(shortKey); err != nil {
				return
			}

			if ts < startFrom {
				_, _ = aep.Application.ConsulClient.KV().Delete(kv.Key, nil)
				listLength--
			} else {
				userActionList = append(userActionList, &userAction)
				if ts > newestTs {
					newestTs = ts
				}
			}
		}

		time.Sleep(5 * time.Second)
	}

	_, err = aep.Application.ConsulClient.KV().Put(
		&api.KVPair{Key: startReadFromKey, Value: []byte(strconv.Itoa(startFrom))},
		nil,
	)
	if err != nil {
		return
	}

	return
}
