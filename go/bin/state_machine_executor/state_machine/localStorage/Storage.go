package localStorage

import "github.com/gammazero/deque"

type Storage struct {
	KvStorage    KvStorage
	MessageDeque deque.Deque
}
