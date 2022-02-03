package states

import (
	"context"
	"state_machine_executor/application"
	"state_machine_executor/state_machine/actions"
	"state_machine_executor/state_machine/localStorage"
)

type MessageWait struct {
	MessageToStack bool
	Next           string
}

func (s *MessageWait) Process(ctx context.Context, app *application.Application, ls *localStorage.Storage) (string, bool) {
	if ls.MessageDeque.Len() == 0 {
		return ls.KvStorage.Get("state").(string), true
	}

	msg := ls.MessageDeque.PopFront().(string)
	ls.KvStorage.Set("message", msg)

	if s.MessageToStack {
		actions.MessageToStackFunc(ls, nil, nil)
	}

	return s.Next, false
}
