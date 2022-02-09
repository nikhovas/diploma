package main

import (
	"control/grpcServer"
	"control/interfaces"
	"control/modules/combined"
	"control/modules/consumers/vk"
	"control/modules/kernel"
	"control/modules/staff/telegram"
	"fmt"
	ctrlProto "github.com/nikhovas/diploma/go/lib/proto/controller"
	"github.com/nikhovas/diploma/go/lib/utils/clients"
	"github.com/nikhovas/diploma/go/lib/utils/consts"
	"github.com/nikhovas/diploma/go/lib/utils/distfs"
	"google.golang.org/grpc"
	"net"

	_ "github.com/lib/pq"
)

func createServerSocket() net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:7777"))
	if err != nil {
		panic(err)
	}

	return lis
}

func main() {
	qwConn, qwClient := clients.CreateQuestionWorkerClient()
	defer qwConn.Close()
	tsbConn, tsbClient := clients.CreateTelegramStaffBotClient()
	defer tsbConn.Close()
	cbConn, cbClient := clients.CreateConsumerBotClient()
	defer cbConn.Close()
	redisDb := clients.CreateRedisClient()
	sqlDb := clients.CreateSqlConn()
	lis := createServerSocket()

	grpcCoreServer := grpc.NewServer()
	distFsRoot := distfs.NewRoot(redisDb, nil)

	ctrlProto.RegisterControllerServer(
		grpcCoreServer,
		&grpcServer.Server{
			Database: sqlDb,
			Combined: &combined.Combined{
				Kernel: &kernel.Kernel{
					DistFsRoot:        distFsRoot,
					QwClient:          qwClient,
					ConsumerBotClient: cbClient,
				},
				Staff: map[string]interfaces.Staff{
					"telegram": &telegram.Telegram{
						Bot: tsbClient,
					},
				},
				Consumers: map[string]interfaces.Consumer{
					"vk": &vk.Vk{
						DistFsMetaServiceName: distFsRoot.CdBots().MetaCdServiceName(consts.VkConsumerBotServiceName),
					},
				},
			},
		},
	)
	err := grpcCoreServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
