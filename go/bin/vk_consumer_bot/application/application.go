package application

import (
	consulApi "github.com/hashicorp/consul/api"
	ctrl "github.com/nikhovas/diploma/go/lib/proto/controller"
	"github.com/nikhovas/diploma/go/lib/utils/distfs/bots"
)
import "github.com/go-redis/redis/v8"

type Application struct {
	Consul       *consulApi.Client
	Redis        *redis.Client
	CtrlClient   ctrl.ControllerClient
	VkDistFsBase *bots.MetaServiceName
}
