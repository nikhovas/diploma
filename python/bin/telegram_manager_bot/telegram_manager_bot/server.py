import logging

from aiogram import Bot
from proto.staff_bot import staff_bot_pb2, staff_bot_pb2_grpc
from proto.common import common_pb2
from grpc import aio


class TelegramStaffBotServer(staff_bot_pb2_grpc.TelegramStaffBotServicer):
    def __init__(self, bot: Bot):
        self.bot = bot

    async def SendNewQuestion(self, request: staff_bot_pb2.NewQuestionRequest, context):
        text = '❓ Вопрос\n' + request.question
        await self.bot.send_message(request.groupId, text)
        return common_pb2.EmptyResponse()

    async def NotifyBotStatusTelegramChange(self, request: staff_bot_pb2.NotifyBotStatusChangeTelegramRequest, context):
        if request.enabled:
            text = '🟩 Бот включен'
        else:
            text = '🟥 Бот выключен'
        await self.bot.send_message(request.groupId, text)
        return common_pb2.EmptyResponse()


async def grpc_server(bot: Bot):
    server = aio.server()
    staff_bot_pb2_grpc.add_TelegramStaffBotServicer_to_server(TelegramStaffBotServer(bot), server)
    listen_addr = '[::]:50060'
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()
    await server.wait_for_termination()
