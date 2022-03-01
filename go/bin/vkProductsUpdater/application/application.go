package application

import (
	"context"
	"github.com/nikhovas/diploma/go/lib/proto/vk_products_updater"
	"github.com/nikhovas/diploma/go/lib/utils/clients"
	"github.com/nikhovas/diploma/go/lib/utils/distQueue"
	"github.com/nikhovas/diploma/go/lib/utils/distfs"
	"log"
	"reflect"
	"sync"
	"vkProductsUpdater/queueWorker"
	"vkProductsUpdater/server"
)

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type ApplicationConfig struct {
	QueueWorkerConfig queueWorker.QueueWorkerConfig `yaml:"queue_worker"`
	QueueConfig       distQueue.DistQueueConfig     `yaml:"dist_queue"`
	Server            server.ServerConfig           `yaml:"server"`
}

type Application struct {
	Server      *server.Server
	QueueWorker *queueWorker.QueueWorker
}

func NewApplication(config ApplicationConfig) (Application, error) {
	app := Application{}

	redis := clients.CreateRedisClient()
	consul := clients.CreateConsulClient()
	distFsRoot := distfs.NewRoot(redis, consul)

	dq, err := distQueue.NewDistQueueFromCfg(
		config.QueueConfig,
		reflect.Indirect(reflect.ValueOf(&vk_products_updater.UpdateProductsRequest{})).Type(),
	)
	if err != nil {
		return app, err
	}

	qw, err := queueWorker.NewQueueWorker(config.QueueWorkerConfig, dq, distFsRoot)
	if err != nil {
		return app, err
	}

	s, err := server.NewServer(config.Server, dq)
	if err != nil {
		return app, err
	}

	app = Application{
		Server:      s,
		QueueWorker: qw,
	}

	return app, nil
}

func (a *Application) Run(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	wg.Add(2)
	go func() {
		logError(a.QueueWorker.Run(ctx, wg))
	}()
	go func() {
		logError(a.Server.Run(ctx, wg))
	}()

	return nil
}
