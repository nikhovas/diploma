package VkLongPullServer

import (
	"encoding/json"
	"strconv"
)

type ResponseObject interface{}

type UnknownObjectType struct{}

type AllowMessageObject struct {
	UserId int    `json:"user_id"`
	Key    string `json:"key"`
}

type NewMessageObject struct {
	Id        int           `json:"id"`
	Date      int           `json:"date"`
	Out       int           `json:"out"`
	UserId    int           `json:"user_id"`
	ReadState int           `json:"read_state"`
	Title     string        `json:"title"`
	Body      string        `json:"body"`
	OwnerIds  []interface{} `json:"owner_ids"`
}

type UpdateObject struct {
	Object  interface{} `json:"object"`
	GroupId int         `json:"group_id"`
	EventId string      `json:"event_id"`
}

type ACheckResponse struct {
	Ts      string         `json:"ts"`
	Updates []UpdateObject `json:"updates"`
}

type RawUpdateObject struct {
	Type    string      `json:"type"`
	Object  interface{} `json:"object"`
	GroupId int         `json:"group_id"`
	EventId string      `json:"event_id"`
}

func (rawUpdateObject *RawUpdateObject) ToUpdateObject() (UpdateObject, bool) {
	updateObject := UpdateObject{
		Object:  nil,
		GroupId: rawUpdateObject.GroupId,
		EventId: rawUpdateObject.EventId,
	}

	switch rawUpdateObject.Type {
	case "message_new":
		obj := NewMessageObject{}
		data, _ := json.Marshal(rawUpdateObject.Object)
		json.Unmarshal(data, &obj)
		updateObject.Object = &obj
	default:
		updateObject.Object = UnknownObjectType{}
		return updateObject, false
	}

	return updateObject, true
}

type ACheckRawResponse struct {
	Ts      string            `json:"ts"`
	Updates []RawUpdateObject `json:"updates"`
}

func (vkLongPullServer *VkLongPullServer) ACheck(ts int, wait int) (ACheckResponse, error) {
	var resp ACheckRawResponse
	err := vkLongPullServer.SendActionRequest(
		"a_check",
		map[string]string{
			"ts":   strconv.Itoa(ts),
			"wait": strconv.Itoa(wait),
		},
		&resp,
	)

	final := ACheckResponse{
		Ts:      resp.Ts,
		Updates: []UpdateObject{},
	}

	for _, rawUpdateObject := range resp.Updates {
		updateObject, ok := rawUpdateObject.ToUpdateObject()
		if ok {
			final.Updates = append(final.Updates, updateObject)
		}
	}

	return final, err
}
