package states

import (
	"MsgCombiner/stateMachine/actions"
	"MsgCombiner/stateMachine/localStorage"
)

type MessageWait struct {
	MessageToStack bool
	Next           string
}

func (s *MessageWait) Process(storage *localStorage.Storage) (string, bool) {
	if storage.MessageDeque.Len() == 0 {
		return storage.KvStorage.Get("state").(string), true
	}

	msg := storage.MessageDeque.PopFront().(string)
	storage.KvStorage.Set("message", msg)

	if s.MessageToStack {
		actions.MessageToStackFunc(storage, nil, nil)
	}

	return s.Next, false
}
