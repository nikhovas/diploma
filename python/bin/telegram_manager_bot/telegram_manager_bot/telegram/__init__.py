import logging
import grpc

from aiogram import Bot, Dispatcher

from proto.controller import controller_pb2_grpc


API_TOKEN = '1946618862:AAHgQl-2ZQnkYogkYyzPwZYqq50DBZJIZhw'
QUESTION_MARKER = '❓ Вопрос\n'

# Configure logging
logging.basicConfig(level=logging.INFO)

# Initialize bot and dispatcher
bot = Bot(token=API_TOKEN)
dp = Dispatcher(bot)


channel = grpc.aio.insecure_channel('localhost:7777')
stub = controller_pb2_grpc.ControllerStub(channel)
