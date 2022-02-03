# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto.common import common_pb2 as common__pb2
from proto.consumer_bot import consumer_bot_pb2 as consumer__bot__pb2


class VkServerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SendSimpleMessage = channel.unary_unary(
                '/consumer_bot.VkServer/SendSimpleMessage',
                request_serializer=consumer__bot__pb2.SimpleMessageInformation.SerializeToString,
                response_deserializer=common__pb2.EmptyResponse.FromString,
                )
        self.SendReplyMessage = channel.unary_unary(
                '/consumer_bot.VkServer/SendReplyMessage',
                request_serializer=consumer__bot__pb2.ReplyMessageInformation.SerializeToString,
                response_deserializer=common__pb2.EmptyResponse.FromString,
                )
        self.AddBot = channel.unary_unary(
                '/consumer_bot.VkServer/AddBot',
                request_serializer=consumer__bot__pb2.BotsActionRequest.SerializeToString,
                response_deserializer=common__pb2.EmptyResponse.FromString,
                )
        self.RemoveBot = channel.unary_unary(
                '/consumer_bot.VkServer/RemoveBot',
                request_serializer=consumer__bot__pb2.BotsActionRequest.SerializeToString,
                response_deserializer=common__pb2.EmptyResponse.FromString,
                )


class VkServerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SendSimpleMessage(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendReplyMessage(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddBot(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RemoveBot(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_VkServerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'SendSimpleMessage': grpc.unary_unary_rpc_method_handler(
                    servicer.SendSimpleMessage,
                    request_deserializer=consumer__bot__pb2.SimpleMessageInformation.FromString,
                    response_serializer=common__pb2.EmptyResponse.SerializeToString,
            ),
            'SendReplyMessage': grpc.unary_unary_rpc_method_handler(
                    servicer.SendReplyMessage,
                    request_deserializer=consumer__bot__pb2.ReplyMessageInformation.FromString,
                    response_serializer=common__pb2.EmptyResponse.SerializeToString,
            ),
            'AddBot': grpc.unary_unary_rpc_method_handler(
                    servicer.AddBot,
                    request_deserializer=consumer__bot__pb2.BotsActionRequest.FromString,
                    response_serializer=common__pb2.EmptyResponse.SerializeToString,
            ),
            'RemoveBot': grpc.unary_unary_rpc_method_handler(
                    servicer.RemoveBot,
                    request_deserializer=consumer__bot__pb2.BotsActionRequest.FromString,
                    response_serializer=common__pb2.EmptyResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'consumer_bot.VkServer', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class VkServer(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SendSimpleMessage(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/consumer_bot.VkServer/SendSimpleMessage',
            consumer__bot__pb2.SimpleMessageInformation.SerializeToString,
            common__pb2.EmptyResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SendReplyMessage(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/consumer_bot.VkServer/SendReplyMessage',
            consumer__bot__pb2.ReplyMessageInformation.SerializeToString,
            common__pb2.EmptyResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddBot(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/consumer_bot.VkServer/AddBot',
            consumer__bot__pb2.BotsActionRequest.SerializeToString,
            common__pb2.EmptyResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RemoveBot(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/consumer_bot.VkServer/RemoveBot',
            consumer__bot__pb2.BotsActionRequest.SerializeToString,
            common__pb2.EmptyResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
