package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/hashicorp/consul/api"
	"github.com/nikhovas/diploma/lib/go/vk/apiServer"
	"github.com/nikhovas/diploma/lib/go/vk/longPullServer"
	actions "github.com/nikhovas/diploma/proto/data/actionEvent"
	UserActions "github.com/nikhovas/diploma/proto/data/userActions"
	pb "github.com/nikhovas/diploma/proto/servers/VkServer"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
	"vk_shop_bot/bots"
	"vk_shop_bot/server"
)

var coordinator *api.KV
var amqpChannel *amqp.Channel
var amqpQueue amqp.Queue

func GetDataBasePath(groupId int, userId int) string {
	return fmt.Sprintf("messages/vk-shop-bot/%d/%d", groupId, userId)
}

func GetActionsBasePath(groupId int, userId int) string {
	return fmt.Sprintf("%s/actions", GetDataBasePath(groupId, userId))
}

func GetMessagesKey(groupId int, userId int) string {
	return fmt.Sprintf("%s/messages", GetDataBasePath(groupId, userId))
}

func GetMessagePath(groupId int, userId int, messageTs int) string {
	return fmt.Sprintf("%s/%d", GetMessagesKey(groupId, userId), messageTs)
}

func callback(groupId int, ts int, update longPullServer.UpdateObject) {
	ro := update.Object

	var userId int
	ae := &actions.ActionEvent{
		UserId: "",
		BotId:  strconv.Itoa(groupId),
		Time:   uint64(ts),
	}
	userAction := UserActions.UserAction{
		Time:   uint64(ts),
		Object: nil,
	}

	switch v := ro.(type) {
	case *longPullServer.NewMessageObject:
		userAction.Object = &UserActions.UserAction_NewMessage{
			NewMessage: &UserActions.NewMessage{Text: v.Body},
		}

		userId = v.UserId

	default:
		fmt.Println("Unsupported message type")
	}

	ae.UserId = strconv.Itoa(userId)
	ae.ServiceName = "vk-shop-bot"

	m := jsonpb.Marshaler{}
	userActionString, _ := m.MarshalToString(&userAction)
	_, _ = coordinator.Put(&api.KVPair{Key: GetMessagePath(groupId, userId, ts), Value: []byte(userActionString)}, nil)

	aeString, _ := m.MarshalToString(ae)
	err := amqpChannel.Publish("", amqpQueue.Name, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(aeString),
		})
	if err != nil {
		log.Println(err)
	}
}

func grpcServer(bot *bots.CombinedBot) {
	lis, _ := net.Listen("tcp", fmt.Sprintf("localhost:5555"))

	grpcServer := grpc.NewServer()
	pb.RegisterVkServerServer(grpcServer, server.NewServer(bot))
	grpcServer.Serve(lis)
}

func setUpAmqp() {
	const amqpUrl = "amqp://guest:guest@localhost:5672/"
	const queueName = "action_events"

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
	} else {
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

func main() {
	const vkApiHost = "http://api.vk.com"
	const vkApiVersion = "5.89"

	//const token = "df706c54e7ab475336001dc165d6143bf344211fe84a41a54334d451b583fb8cc247e021d9f2b285d1ed3"
	//const groupId = 209867018

	setUpAmqp()

	client, _ := api.NewClient(api.DefaultConfig())
	coordinator = client.KV()

	vkApiServer := apiServer.VkApiServer{
		Host:    vkApiHost,
		Version: vkApiVersion,
	}

	var bot bots.CombinedBot
	bot.Init(client.KV(), &vkApiServer, callback)

	var wg sync.WaitGroup
	wg.Add(2)
	go bot.Run(&wg)
	//time.Sleep(3 * time.Second)
	//bot.AddBot(groupId)
	time.Sleep(3 * time.Second)
	go grpcServer(&bot)

	wg.Wait()
}
