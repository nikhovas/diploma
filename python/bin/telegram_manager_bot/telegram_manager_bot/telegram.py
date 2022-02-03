import logging
import grpc

from aiogram import Bot, Dispatcher, types

from proto.controller import controller_pb2, controller_pb2_grpc


API_TOKEN = '1946618862:AAHgQl-2ZQnkYogkYyzPwZYqq50DBZJIZhw'
QUESTION_MARKER = 'Вопрос: '

# Configure logging
logging.basicConfig(level=logging.INFO)

# Initialize bot and dispatcher
bot = Bot(token=API_TOKEN)
dp = Dispatcher(bot)


channel = grpc.aio.insecure_channel('localhost:7777')
stub = controller_pb2_grpc.ControllerStub(channel)


class BotError(Exception):
    def __init__(self, grpc_resp):
        self.grpc_resp = grpc_resp


def action_decorator(f):
    async def internal(message: types.Message):
        cmi = controller_pb2.TelegramMessageInfo(
            chatId=message.chat.id,
            isChatPrivate=(message.chat.type == 'private'),
            userId=message.from_user.id,
        )
        try:
            await f(message, cmi)
        except BotError as err:
            await message.answer('asdf')

    return internal


def check_resp_for_error(resp):
    pass


def group_only(f):
    async def internal(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
        if cmi.isChatPrivate is True:
            raise BotError(controller_pb2.DefaultResponse(
                noRoleError=controller_pb2.BadChatType(shouldBePrivate=False)),
            )
            # await message.reply('Команда работает только в группе магазина')
            # return
        await f(message, cmi)

    return internal


def private_only(f):
    async def internal(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
        if cmi.isChatPrivate is False:
            raise BotError(controller_pb2.DefaultResponse(
                noRoleError=controller_pb2.BadChatType(shouldBePrivate=True)),
            )
            # await message.reply('Команда работает только в личных сообщениях')
            # return
        await f(message, cmi)

    return internal


@dp.message_handler(commands=['start', 'help'])
async def send_welcome(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    print(message.from_user)
    print(message.chat)
    await message.reply("Hi!\nI'm EchoBot!\nPowered by aiogram.")


@dp.message_handler(commands=['answer'])
@action_decorator
@group_only
async def answer_message(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    answer = message.text[len('/answer'):]
    answer = answer.strip()
    reply_msg = message.reply_to_message
    if reply_msg is None:
        await message.reply('Нет прикрепленного сообщения')
        return

    question: str = reply_msg.text
    if not question.startswith(QUESTION_MARKER):
        await message.reply('Нету прикрепленного сообщения с вопросом')
        return

    question = question[len(QUESTION_MARKER):]
    question = question.strip()

    response = await stub.AddQuestionAnswer(
        controller_pb2.AddQuestionAnswerRequest(question=question, answer=answer),
    )
    print("Greeter client received: " + response.message)
    # check for role

    # add to database

    # send answer to client

    print(question)
    print(answer)

    await message.reply('Ваш ответ зарегестрирован')


@dp.message_handler(commands=['addquestion'])
@action_decorator
@group_only
async def add_question(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    lines = message.text.split('\n')
    question = lines[1]
    answer = lines[2]

    # check for role

    # add to database

    # send answer to client

    print(question)
    print(answer)

    await message.reply('Ваш вопрос добавлен')


@dp.message_handler(commands=['stopshopbot'])
@action_decorator
async def stop_shop_bot(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    response = await stub.ChangeBotState(controller_pb2.ChangeBotStateRequest(
        messageInfo=controller_pb2.MessageInformation(telegram=cmi),
        toEnabled=False,
    ))
    await message.reply('Бот магазина остановлен')


@dp.message_handler(commands=['startshopbot'])
@action_decorator
async def start_shop_bot(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    response = await stub.ChangeBotState(controller_pb2.ChangeBotStateRequest(
        messageInfo=controller_pb2.MessageInformation(telegram=cmi),
        toEnabled=True,
    ))
    await message.reply('Бот магазина запущен')


@dp.message_handler(commands=['addshop'])
@action_decorator
@private_only
async def add_shop(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    try:
        tokens = message.text.split(' ')
        telegram_group_id = int(tokens[1])
        name = tokens[2]
        vk_token = tokens[3]
        vk_group_id = int(tokens[4])
    except Exception:
        await message.reply('Ошибка')
        return

    response = await stub.AddShop(controller_pb2.AddShopRequest(
        messageInfo=controller_pb2.MessageInformation(telegram=cmi),
        bot=controller_pb2.BotInfo(name=name, token=vk_token, groupId=vk_group_id),
        platformBotInfo=controller_pb2.PlatformBotInfo(platform=controller_pb2.TelegramBotInfo(chatId=cmi.chatId))
    ))
    # response = await stub.AddQuestionAnswer(
    #     controller_pb2.AddQuestionAnswerRequest(question=question, answer=answer),
    # )

    # start

    # send answer to client

    await message.reply('Бот магазина запущен')


@dp.message_handler(commands=['deleteshop'])
@action_decorator
@private_only
async def delete_shop(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    # check for role

    # start

    # send answer to client

    await message.reply('Бот магазина запущен')


@dp.message_handler(commands=['modifyshop'])
@action_decorator
@private_only
async def modify_shop(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    # check for role

    # start

    # send answer to client

    await message.reply('Бот магазина запущен')


@dp.message_handler(commands=['register_group'])
@action_decorator
@group_only
async def register_shop_group(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    try:
        tokens = message.text.split(' ')
        shop_name = tokens[1]
    except Exception:
        await message.reply('Ошибка')
        return

    response = await stub.ModifyShop(controller_pb2.ModifyShopRequest(
        messageInfo=controller_pb2.MessageInformation(telegram=cmi),
        key=controller_pb2.ShopKey(name=shop_name),
        bot=controller_pb2.OptionalBotInfo(
            platformBotInfo=controller_pb2.OptionalPlatformBotInfo(
                telegram=controller_pb2.OptionalTelegramBotInfo(
                    chatId=cmi.chatId,
                )
            )
        ),
    ))

    await message.reply('Ваш вопрос добавлен')


@dp.message_handler()
async def echo(message: types.Message):
    await message.answer(message.text)