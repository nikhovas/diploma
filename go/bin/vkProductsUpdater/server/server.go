package server

import (
	"context"
	"fmt"
	"github.com/nikhovas/diploma/go/lib/proto/common"
	"github.com/nikhovas/diploma/go/lib/proto/vk_products_updater"
	"github.com/nikhovas/diploma/go/lib/utils/distQueue"
	"google.golang.org/grpc"
	"net"
	"sync"
)

type ServerInternal struct {
	vk_products_updater.UnimplementedVkProductsUpdaterServer
	DistQueue *distQueue.DistQueue
}

func (s *ServerInternal) UpdateProducts(
	ctx context.Context,
	req *vk_products_updater.UpdateProductsRequest,
) (*common.EmptyResponse, error) {
	err := s.DistQueue.Push(req, 0)
	if err != nil {
		return nil, err
	}
	return &common.EmptyResponse{}, nil
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Server struct {
	serverInternal ServerInternal
	grpcServer     *grpc.Server
	config         ServerConfig
}

func (s *Server) Run(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	lis, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", s.config.Host, s.config.Port))
	err := s.grpcServer.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}

func NewServer(config ServerConfig, dq *distQueue.DistQueue) (*Server, error) {
	result := Server{
		serverInternal: ServerInternal{
			DistQueue: dq,
		},
		grpcServer: grpc.NewServer(),
		config:     config,
	}

	vk_products_updater.RegisterVkProductsUpdaterServer(result.grpcServer, &result.serverInternal)
	return &result, nil
}
