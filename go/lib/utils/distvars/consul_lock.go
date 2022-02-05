package distvars

import consulApi "github.com/hashicorp/consul/api"

type ConsulLock struct {
	Storage *consulApi.Client
	Path    string
	lock    *consulApi.Lock
}

func (dv *ConsulLock) TryLock() (bool, error) {
	var err error
	dv.lock, err = dv.Storage.LockOpts(&consulApi.LockOptions{
		Key:         dv.Path,
		LockTryOnce: true,
	})
	if err != nil {
		return false, err
	}

	ch, err := dv.lock.Lock(nil)
	if err != nil {
		return false, err
	} else if ch == nil {
		return false, nil
	}

	return true, nil
}

func (dv *ConsulLock) Unlock() error {
	return dv.lock.Unlock()
}
