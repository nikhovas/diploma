package grpcClient

import (
	"github.com/nikhovas/diploma/proto/servers/VkServer"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	Vk VkServer.VkServerClient
}

func (c *GrpcClient) Init() {
	conn, err := grpc.Dial("localhost:5555", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	c.Vk = VkServer.NewVkServerClient(conn)
}

var GrpcClientSingletone GrpcClient
