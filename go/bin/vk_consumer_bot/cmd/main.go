package main

import (
	pb "github.com/nikhovas/diploma/go/lib/proto/consumer_bot"
	"github.com/nikhovas/diploma/go/lib/utils/clients"
	"github.com/nikhovas/diploma/go/lib/utils/consts"
	"github.com/nikhovas/diploma/go/lib/utils/distfs"
	"github.com/nikhovas/diploma/go/lib/utils/env"
	"github.com/nikhovas/diploma/go/lib/utils/log"
	apiServer "github.com/nikhovas/diploma/go/lib/vk/api_server"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
	"vk_consumer_bot/application"
	"vk_consumer_bot/bots"
	"vk_consumer_bot/server"
)

var amqpChannel *amqp.Channel
var amqpQueue amqp.Queue

func setUpAmqp() {
	amqpUrl := env.GetAmqpUrl()
	queueName := "action_events"

	var ampqConn *amqp.Connection
	var err error

	for i := 0; i < 50; i++ {
		ampqConn, err = amqp.Dial(amqpUrl)
		if err != nil {
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}

	if err != nil {
		panic(err)
	}

	amqpChannel, err = ampqConn.Channel()
	if err != nil {
		panic(err)
		return
	}
	amqpQueue, err = amqpChannel.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		panic(err)
		return
	}
}

func runGrpcServer(bot *bots.CombinedBot) {
	lis, _ := net.Listen("tcp", env.GetVkConsumerBotGrpcHost())

	grpcServer := grpc.NewServer()
	pb.RegisterVkServerServer(grpcServer, server.NewServer(bot))
	err := grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func main() {
	const vkApiHost = "http://api.vk.com"
	const vkApiVersion = "5.92"

	ctrlConn, ctrlClient := clients.CreateControllerClient()
	defer ctrlConn.Close()

	setUpAmqp()

	consulClient := clients.CreateConsulClient()

	vkApiServer := apiServer.VkApiServer{
		Host:    vkApiHost,
		Version: vkApiVersion,
	}

	rdb := clients.CreateRedisClient()
	defer rdb.Close()

	vkServiceDistFs := distfs.NewRoot(rdb, consulClient).CdBots().MetaCdServiceName(consts.VkConsumerBotServiceName)

	app := application.Application{
		Consul:       consulClient,
		Redis:        rdb,
		CtrlClient:   ctrlClient,
		VkDistFsBase: vkServiceDistFs,
	}

	var bot bots.CombinedBot
	bot.Init(&app, &vkApiServer, callback)

	var wg sync.WaitGroup
	wg.Add(3)

	globalContext := log.NewRootContext()

	go bot.Run(&wg)
	time.Sleep(3 * time.Second)
	go runGrpcServer(&bot)

	go redisReadWorker(log.NewContext(globalContext), &app, &bot)

	wg.Wait()
}
