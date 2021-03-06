# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto.controller import controller_pb2 as controller__pb2


class ControllerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ListShops = channel.unary_unary(
                '/controller.Controller/ListShops',
                request_serializer=controller__pb2.ListShopsRequest.SerializeToString,
                response_deserializer=controller__pb2.ListShopsResponse.FromString,
                )
        self.AddShop = channel.unary_unary(
                '/controller.Controller/AddShop',
                request_serializer=controller__pb2.AddShopRequest.SerializeToString,
                response_deserializer=controller__pb2.DefaultResponse.FromString,
                )
        self.ModifyShop = channel.unary_unary(
                '/controller.Controller/ModifyShop',
                request_serializer=controller__pb2.ModifyShopRequest.SerializeToString,
                response_deserializer=controller__pb2.DefaultResponse.FromString,
                )
        self.DeleteShop = channel.unary_unary(
                '/controller.Controller/DeleteShop',
                request_serializer=controller__pb2.DeleteShopRequest.SerializeToString,
                response_deserializer=controller__pb2.DefaultResponse.FromString,
                )
        self.AddQuestionAnswer = channel.unary_unary(
                '/controller.Controller/AddQuestionAnswer',
                request_serializer=controller__pb2.AddQuestionAnswerRequest.SerializeToString,
                response_deserializer=controller__pb2.DefaultResponse.FromString,
                )
        self.AddQuestion = channel.unary_unary(
                '/controller.Controller/AddQuestion',
                request_serializer=controller__pb2.AddQuestionRequest.SerializeToString,
                response_deserializer=controller__pb2.DefaultResponse.FromString,
                )
        self.ChangeBotState = channel.unary_unary(
                '/controller.Controller/ChangeBotState',
                request_serializer=controller__pb2.ChangeBotStateRequest.SerializeToString,
                response_deserializer=controller__pb2.DefaultResponse.FromString,
                )
        self.ChangeUserRole = channel.unary_unary(
                '/controller.Controller/ChangeUserRole',
                request_serializer=controller__pb2.ChangeUserRoleRequest.SerializeToString,
                response_deserializer=controller__pb2.DefaultResponse.FromString,
                )
        self.NotifyBotStatusChange = channel.unary_unary(
                '/controller.Controller/NotifyBotStatusChange',
                request_serializer=controller__pb2.NotifyBotStatusChangeRequest.SerializeToString,
                response_deserializer=controller__pb2.DefaultResponse.FromString,
                )
        self.GetShopIdByKey = channel.unary_unary(
                '/controller.Controller/GetShopIdByKey',
                request_serializer=controller__pb2.ShopKey.SerializeToString,
                response_deserializer=controller__pb2.GetShopIdByKeyResponse.FromString,
                )


class ControllerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ListShops(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddShop(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ModifyShop(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteShop(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddQuestionAnswer(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AddQuestion(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ChangeBotState(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ChangeUserRole(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def NotifyBotStatusChange(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetShopIdByKey(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ControllerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ListShops': grpc.unary_unary_rpc_method_handler(
                    servicer.ListShops,
                    request_deserializer=controller__pb2.ListShopsRequest.FromString,
                    response_serializer=controller__pb2.ListShopsResponse.SerializeToString,
            ),
            'AddShop': grpc.unary_unary_rpc_method_handler(
                    servicer.AddShop,
                    request_deserializer=controller__pb2.AddShopRequest.FromString,
                    response_serializer=controller__pb2.DefaultResponse.SerializeToString,
            ),
            'ModifyShop': grpc.unary_unary_rpc_method_handler(
                    servicer.ModifyShop,
                    request_deserializer=controller__pb2.ModifyShopRequest.FromString,
                    response_serializer=controller__pb2.DefaultResponse.SerializeToString,
            ),
            'DeleteShop': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteShop,
                    request_deserializer=controller__pb2.DeleteShopRequest.FromString,
                    response_serializer=controller__pb2.DefaultResponse.SerializeToString,
            ),
            'AddQuestionAnswer': grpc.unary_unary_rpc_method_handler(
                    servicer.AddQuestionAnswer,
                    request_deserializer=controller__pb2.AddQuestionAnswerRequest.FromString,
                    response_serializer=controller__pb2.DefaultResponse.SerializeToString,
            ),
            'AddQuestion': grpc.unary_unary_rpc_method_handler(
                    servicer.AddQuestion,
                    request_deserializer=controller__pb2.AddQuestionRequest.FromString,
                    response_serializer=controller__pb2.DefaultResponse.SerializeToString,
            ),
            'ChangeBotState': grpc.unary_unary_rpc_method_handler(
                    servicer.ChangeBotState,
                    request_deserializer=controller__pb2.ChangeBotStateRequest.FromString,
                    response_serializer=controller__pb2.DefaultResponse.SerializeToString,
            ),
            'ChangeUserRole': grpc.unary_unary_rpc_method_handler(
                    servicer.ChangeUserRole,
                    request_deserializer=controller__pb2.ChangeUserRoleRequest.FromString,
                    response_serializer=controller__pb2.DefaultResponse.SerializeToString,
            ),
            'NotifyBotStatusChange': grpc.unary_unary_rpc_method_handler(
                    servicer.NotifyBotStatusChange,
                    request_deserializer=controller__pb2.NotifyBotStatusChangeRequest.FromString,
                    response_serializer=controller__pb2.DefaultResponse.SerializeToString,
            ),
            'GetShopIdByKey': grpc.unary_unary_rpc_method_handler(
                    servicer.GetShopIdByKey,
                    request_deserializer=controller__pb2.ShopKey.FromString,
                    response_serializer=controller__pb2.GetShopIdByKeyResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'controller.Controller', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Controller(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ListShops(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/controller.Controller/ListShops',
            controller__pb2.ListShopsRequest.SerializeToString,
            controller__pb2.ListShopsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddShop(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/controller.Controller/AddShop',
            controller__pb2.AddShopRequest.SerializeToString,
            controller__pb2.DefaultResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ModifyShop(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/controller.Controller/ModifyShop',
            controller__pb2.ModifyShopRequest.SerializeToString,
            controller__pb2.DefaultResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteShop(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/controller.Controller/DeleteShop',
            controller__pb2.DeleteShopRequest.SerializeToString,
            controller__pb2.DefaultResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddQuestionAnswer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/controller.Controller/AddQuestionAnswer',
            controller__pb2.AddQuestionAnswerRequest.SerializeToString,
            controller__pb2.DefaultResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AddQuestion(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/controller.Controller/AddQuestion',
            controller__pb2.AddQuestionRequest.SerializeToString,
            controller__pb2.DefaultResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ChangeBotState(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/controller.Controller/ChangeBotState',
            controller__pb2.ChangeBotStateRequest.SerializeToString,
            controller__pb2.DefaultResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ChangeUserRole(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/controller.Controller/ChangeUserRole',
            controller__pb2.ChangeUserRoleRequest.SerializeToString,
            controller__pb2.DefaultResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def NotifyBotStatusChange(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/controller.Controller/NotifyBotStatusChange',
            controller__pb2.NotifyBotStatusChangeRequest.SerializeToString,
            controller__pb2.DefaultResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetShopIdByKey(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/controller.Controller/GetShopIdByKey',
            controller__pb2.ShopKey.SerializeToString,
            controller__pb2.GetShopIdByKeyResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
