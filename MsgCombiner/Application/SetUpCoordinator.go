package Application

import "github.com/hashicorp/consul/api"

func (a *Application) SetUpCoordinator() (err error) {
	a.ConsulClient, err = api.NewClient(api.DefaultConfig())
	return
}
