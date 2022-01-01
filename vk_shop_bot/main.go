package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/nikhovas/diploma/proto/data/userActions"
	pb "github.com/nikhovas/diploma/proto/servers/VkServer"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
	"vk_shop_bot/bots"
	"vk_shop_bot/server"
	"vk_shop_bot/vkApi/VkApiServer"
	"vk_shop_bot/vkApi/VkLongPullServer"
)

func GetDataBasePath(groupId int, userId int) string {
	return fmt.Sprintf("messages/vk-shop-bot/%d/%d", groupId, userId)
}

func GetActionsBasePath(groupId int, userId int) string {
	return fmt.Sprintf("%s/actions", GetDataBasePath(groupId, userId))
}

func GetActions(groupId int, userId int) string {
	return fmt.Sprintf("%s/messages", GetDataBasePath(groupId, userId))
}

func callback(groupId int, update VkLongPullServer.UpdateObject) {
	ro := update.Object
	switch v := ro.(type) {
	case *VkLongPullServer.NewMessageObject:
		userAction := UserActions.UserAction{
			ActionType: "",
			Object:     nil,
		}
		fmt.Println(v.Body)
	default:
		fmt.Println("Unsupported message type")
	}
}

func grpcServer(bot *bots.CombinedBot) {
	lis, _ := net.Listen("tcp", fmt.Sprintf("localhost:5555"))

	grpcServer := grpc.NewServer()
	pb.RegisterVkServerServer(grpcServer, server.NewServer(bot))
	grpcServer.Serve(lis)
}

func main() {
	const vkApiHost = "http://api.vk.com"
	const vkApiVersion = "5.89"

	//const token = "df706c54e7ab475336001dc165d6143bf344211fe84a41a54334d451b583fb8cc247e021d9f2b285d1ed3"
	//const groupId = 209867018

	client, _ := api.NewClient(api.DefaultConfig())

	vkApiServer := VkApiServer.VkApiServer{
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
