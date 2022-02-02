import asyncio
import logging

from aiogram import Bot
from cmd import telegram
from grpc_server import TelegramStaffBot_pb2_grpc, TelegramStaffBot_pb2
from grpc import aio


class TelegramStaffBotServer(TelegramStaffBot_pb2_grpc.TelegramStaffBotServicer):
    def __init__(self, bot: Bot):
        self.bot = bot

    async def SendNewQuestion(self, request: TelegramStaffBot_pb2.NewQuestionRequest, context):
        text = '❓ Вопрос\n' + request.question
        await self.bot.send_message(request.groupId, text)
        return TelegramStaffBot_pb2.EmptyResponse()

    async def NotifyBotStatusTelegramChange(self, request: TelegramStaffBot_pb2.NotifyBotStatusChangeTelegramRequest, context):
        text = '!!! Бот '
        if request.enabled:
            text += 'Включен'
        else:
            text += 'Выключен'
        await self.bot.send_message(request.groupId, text)
        return TelegramStaffBot_pb2.EmptyResponse()


async def grpc_server(bot: Bot):
    server = aio.server()
    TelegramStaffBot_pb2_grpc.add_TelegramStaffBotServicer_to_server(TelegramStaffBotServer(bot), server)
    listen_addr = '[::]:50060'
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()
    await server.wait_for_termination()
