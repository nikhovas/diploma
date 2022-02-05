// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package consumer_bot

import (
	context "context"
	common "github.com/nikhovas/diploma/go/lib/proto/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VkServerClient is the client API for VkServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VkServerClient interface {
	SendSimpleMessage(ctx context.Context, in *SimpleMessageInformation, opts ...grpc.CallOption) (*common.EmptyResponse, error)
	SendReplyMessage(ctx context.Context, in *ReplyMessageInformation, opts ...grpc.CallOption) (*common.EmptyResponse, error)
}

type vkServerClient struct {
	cc grpc.ClientConnInterface
}

func NewVkServerClient(cc grpc.ClientConnInterface) VkServerClient {
	return &vkServerClient{cc}
}

func (c *vkServerClient) SendSimpleMessage(ctx context.Context, in *SimpleMessageInformation, opts ...grpc.CallOption) (*common.EmptyResponse, error) {
	out := new(common.EmptyResponse)
	err := c.cc.Invoke(ctx, "/consumer_bot.VkServer/SendSimpleMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vkServerClient) SendReplyMessage(ctx context.Context, in *ReplyMessageInformation, opts ...grpc.CallOption) (*common.EmptyResponse, error) {
	out := new(common.EmptyResponse)
	err := c.cc.Invoke(ctx, "/consumer_bot.VkServer/SendReplyMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VkServerServer is the server API for VkServer service.
// All implementations must embed UnimplementedVkServerServer
// for forward compatibility
type VkServerServer interface {
	SendSimpleMessage(context.Context, *SimpleMessageInformation) (*common.EmptyResponse, error)
	SendReplyMessage(context.Context, *ReplyMessageInformation) (*common.EmptyResponse, error)
	mustEmbedUnimplementedVkServerServer()
}

// UnimplementedVkServerServer must be embedded to have forward compatible implementations.
type UnimplementedVkServerServer struct {
}

func (UnimplementedVkServerServer) SendSimpleMessage(context.Context, *SimpleMessageInformation) (*common.EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSimpleMessage not implemented")
}
func (UnimplementedVkServerServer) SendReplyMessage(context.Context, *ReplyMessageInformation) (*common.EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendReplyMessage not implemented")
}
func (UnimplementedVkServerServer) mustEmbedUnimplementedVkServerServer() {}

// UnsafeVkServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VkServerServer will
// result in compilation errors.
type UnsafeVkServerServer interface {
	mustEmbedUnimplementedVkServerServer()
}

func RegisterVkServerServer(s grpc.ServiceRegistrar, srv VkServerServer) {
	s.RegisterService(&VkServer_ServiceDesc, srv)
}

func _VkServer_SendSimpleMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessageInformation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VkServerServer).SendSimpleMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consumer_bot.VkServer/SendSimpleMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VkServerServer).SendSimpleMessage(ctx, req.(*SimpleMessageInformation))
	}
	return interceptor(ctx, in, info, handler)
}

func _VkServer_SendReplyMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyMessageInformation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VkServerServer).SendReplyMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consumer_bot.VkServer/SendReplyMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VkServerServer).SendReplyMessage(ctx, req.(*ReplyMessageInformation))
	}
	return interceptor(ctx, in, info, handler)
}

// VkServer_ServiceDesc is the grpc.ServiceDesc for VkServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VkServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "consumer_bot.VkServer",
	HandlerType: (*VkServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSimpleMessage",
			Handler:    _VkServer_SendSimpleMessage_Handler,
		},
		{
			MethodName: "SendReplyMessage",
			Handler:    _VkServer_SendReplyMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consumer_bot.proto",
}
