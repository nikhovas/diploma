package grpcserver

import (
	"context"
	"fmt"
	pb "github.com/nikhovas/diploma/go/lib/proto/consumer_bot"
	"google.golang.org/grpc"
	"net"
	"sync"
	"vk_consumer_bot/modules/combinedbot"
)

type GrpcServer struct {
	address string
	bot *combinedbot.CombinedBot
	server *grpc.Server
}

func FromConfig(config Config, bot *combinedbot.CombinedBot) *GrpcServer {
	return &GrpcServer{
		address: fmt.Sprintf("%s:%d", config.Host, config.Port),
		bot: bot,
		server: grpc.NewServer(),
	}
}

func (gs *GrpcServer) Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go gs.RunBlock(ctx, wg)
}

func (gs *GrpcServer) RunBlock(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	lis, _ := net.Listen("tcp", gs.address)
	pb.RegisterVkServerServer(gs.server, NewInternalServer(gs.bot))
	if err := gs.server.Serve(lis); err != nil {
		panic(err)
	}
}

func (gs *GrpcServer) Stop() {
	gs.server.Stop()
}
