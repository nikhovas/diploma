package application

import (
	"context"
	consulApi "github.com/hashicorp/consul/api"
	ctrl "github.com/nikhovas/diploma/go/lib/proto/controller"
	"github.com/nikhovas/diploma/go/lib/utils/clients"
	"github.com/nikhovas/diploma/go/lib/utils/consts"
	"github.com/nikhovas/diploma/go/lib/utils/distfs"
	"github.com/nikhovas/diploma/go/lib/utils/distfs/bots"
	apiServer "github.com/nikhovas/diploma/go/lib/vk/api_server"
	"google.golang.org/grpc"
	"sync"
	"time"
	"vk_consumer_bot/modules/combinedbot"
	"vk_consumer_bot/modules/eventsqueuewriter"
	"vk_consumer_bot/modules/grpcserver"
	"vk_consumer_bot/modules/messageobserver"
	"vk_consumer_bot/modules/readmessagestosend"
	"vk_consumer_bot/modules/workingbotsupdater"
)
import "github.com/go-redis/redis/v8"

type Application struct {
	Consul       *consulApi.Client
	Redis        *redis.Client
	CtrlClient   ctrl.ControllerClient
	VkDistFsBase *bots.MetaServiceName

	messageObserver *messageobserver.MessageObserver
	combinedBot *combinedbot.CombinedBot
	grpcServer *grpcserver.GrpcServer
	workingBotsUpdater *workingbotsupdater.WorkingBotsUpdater
	readMessagesToSend *readmessagestosend.ReadMessagesToSend
	eventsQueueWriter *eventsqueuewriter.EventsQueueWriter

	ctrlConn *grpc.ClientConn
}

func FromConfig(config Config) *Application {
	const vkApiHost = "http://api.vk.com"
	const vkApiVersion = "5.92"

	ctrlConn, ctrlClient := clients.CreateControllerClient()
	consulClient := clients.CreateConsulClient()

	vkApiServer := apiServer.VkApiServer{
		Host:    vkApiHost,
		Version: vkApiVersion,
	}

	rdb := clients.CreateRedisClient()

	vkServiceDistFs := distfs.NewRoot(rdb, consulClient).CdBots().MetaCdServiceName(consts.VkConsumerBotServiceName)

	eventsQueueWriter := eventsqueuewriter.FromConfig(config.EventsQueueWriter)

	messageObserver := messageobserver.FromConfig(config.MessageObserver, eventsQueueWriter, vkServiceDistFs, ctrlClient)
	combinedBot := combinedbot.FromConfig(&vkApiServer, messageObserver, vkServiceDistFs, ctrlClient)

	grpcServer := grpcserver.FromConfig(config.GrpcServer, combinedBot)
	workingBotsUpdater := workingbotsupdater.FromConfig(config.WorkingBotsUpdater, vkServiceDistFs, combinedBot)
	readMessagesToSend := readmessagestosend.FromConfig(config.ReadMessagesToSend, combinedBot)

	//return app

	return &Application{
		Consul:       consulClient,
		Redis:        rdb,
		CtrlClient:   ctrlClient,
		VkDistFsBase: vkServiceDistFs,

		messageObserver: messageObserver,
		combinedBot: combinedBot,
		grpcServer: grpcServer,
		workingBotsUpdater: workingBotsUpdater,
		readMessagesToSend: readMessagesToSend,
		eventsQueueWriter: eventsQueueWriter,

		ctrlConn: ctrlConn,
	}
}

func (app *Application) Run(ctx context.Context, wg *sync.WaitGroup) {
	app.eventsQueueWriter.Run(ctx, wg)
	app.messageObserver.Run(ctx, wg)

	app.combinedBot.Run(ctx, wg)
	time.Sleep(3 * time.Second)

	app.grpcServer.Run(ctx, wg)
	app.workingBotsUpdater.Run(ctx, wg)
	app.readMessagesToSend.Run(ctx, wg)
}

func (app *Application) Stop() {
}