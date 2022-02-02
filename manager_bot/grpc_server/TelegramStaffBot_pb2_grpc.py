# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import grpc_server.TelegramStaffBot_pb2 as TelegramStaffBot__pb2


class TelegramStaffBotStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.SendNewQuestion = channel.unary_unary(
                '/TelegramStaffBot/SendNewQuestion',
                request_serializer=TelegramStaffBot__pb2.NewQuestionRequest.SerializeToString,
                response_deserializer=TelegramStaffBot__pb2.EmptyResponse.FromString,
                )
        self.NotifyBotStatusTelegramChange = channel.unary_unary(
                '/TelegramStaffBot/NotifyBotStatusTelegramChange',
                request_serializer=TelegramStaffBot__pb2.NotifyBotStatusChangeTelegramRequest.SerializeToString,
                response_deserializer=TelegramStaffBot__pb2.EmptyResponse.FromString,
                )


class TelegramStaffBotServicer(object):
    """Missing associated documentation comment in .proto file."""

    def SendNewQuestion(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def NotifyBotStatusTelegramChange(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TelegramStaffBotServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'SendNewQuestion': grpc.unary_unary_rpc_method_handler(
                    servicer.SendNewQuestion,
                    request_deserializer=TelegramStaffBot__pb2.NewQuestionRequest.FromString,
                    response_serializer=TelegramStaffBot__pb2.EmptyResponse.SerializeToString,
            ),
            'NotifyBotStatusTelegramChange': grpc.unary_unary_rpc_method_handler(
                    servicer.NotifyBotStatusTelegramChange,
                    request_deserializer=TelegramStaffBot__pb2.NotifyBotStatusChangeTelegramRequest.FromString,
                    response_serializer=TelegramStaffBot__pb2.EmptyResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'TelegramStaffBot', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class TelegramStaffBot(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def SendNewQuestion(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/TelegramStaffBot/SendNewQuestion',
            TelegramStaffBot__pb2.NewQuestionRequest.SerializeToString,
            TelegramStaffBot__pb2.EmptyResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def NotifyBotStatusTelegramChange(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/TelegramStaffBot/NotifyBotStatusTelegramChange',
            TelegramStaffBot__pb2.NotifyBotStatusChangeTelegramRequest.SerializeToString,
            TelegramStaffBot__pb2.EmptyResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)