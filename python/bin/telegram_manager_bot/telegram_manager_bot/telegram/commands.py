from telegram_manager_bot.telegram import dp, decorators, stub, QUESTION_MARKER
from proto.controller import controller_pb2
from aiogram import types


def shop_key_from_chat(chat_id):
    return controller_pb2.ShopKey(telegramStaff=controller_pb2.TelegramStaffShopKey(groupId=chat_id))


@decorators.action('add_question', '\n', 2)
@decorators.group_only
async def add_question(message: types.Message, cmi: controller_pb2.TelegramMessageInfo, question: str, answer: str):
    stub.AddQuestionAnswer(controller_pb2.AddQuestionAnswerRequest(
        uuid='',
        messageInfo=controller_pb2.MessageInformation(telegram=cmi),
        key=shop_key_from_chat(cmi.chatId),
        question=question,
        answer=answer,
        checkQuestionInSet=False,
    ))

    await message.reply('Ваш вопрос добавлен')


@decorators.action('add_shop', ' ', 3)
@decorators.private_only
async def add_shop(message: types.Message, cmi: controller_pb2.TelegramMessageInfo, name, vk_token, vk_group_id):
    vk_group_id = int(vk_group_id)

    response = await stub.AddShop(controller_pb2.AddShopRequest(
        messageInfo=controller_pb2.MessageInformation(telegram=cmi),
        bot=controller_pb2.BotInfo(name=name, token=vk_token, groupId=vk_group_id),
        platformBotInfo=controller_pb2.PlatformBotInfo(platform=controller_pb2.TelegramBotInfo(chatId=0))
    ))

    await message.reply('Бот магазина добавлен')


@decorators.action('answer')
@decorators.group_only
async def answer_message(message: types.Message, cmi: controller_pb2.TelegramMessageInfo, text: str):
    text = text.strip()
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
        controller_pb2.AddQuestionAnswerRequest(
            messageInfo=controller_pb2.MessageInformation(telegram=cmi),
            key=shop_key_from_chat(cmi.chatId),
            question=question,
            answer=text,
        ),
    )

    await message.reply('Ваш ответ зарегестрирован')


@dp.message_handler(commands=['deleteshop'])
@decorators.action('delete_shop', ' ', 1)
@decorators.private_only
async def delete_shop(message: types.Message, cmi: controller_pb2.TelegramMessageInfo, shop_name):
    # check for role

    # start

    # send answer to client

    await message.reply('Бот магазина запущен')


@dp.message_handler(commands=['modifyshop'])
@decorators.action_decorator
@decorators.private_only
async def modify_shop(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    # check for role

    # start

    # send answer to client

    await message.reply('Бот магазина запущен')


@dp.message_handler(commands=['register_group'])
@decorators.action_decorator
@decorators.group_only
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


@dp.message_handler(commands=['startshopbot'])
@decorators.action_decorator
async def start_shop_bot(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    response = await stub.ChangeBotState(controller_pb2.ChangeBotStateRequest(
        messageInfo=controller_pb2.MessageInformation(telegram=cmi),
        toEnabled=True,
    ))
    await message.reply('🕐 Отправлен запрос на включение бота')


@dp.message_handler(commands=['stopshopbot'])
@decorators.action('stop_shop_bot')
async def stop_shop_bot(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    response = await stub.ChangeBotState(controller_pb2.ChangeBotStateRequest(
        messageInfo=controller_pb2.MessageInformation(telegram=cmi),
        toEnabled=False,
    ))
    await message.reply('🕐 Отправлен запрос на остановку бота')


@dp.message_handler(commands=['start', 'help'])
async def send_welcome(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    print(message.from_user)
    print(message.chat)
    await message.reply("Hi!\nI'm EchoBot!\nPowered by aiogram.")


@dp.message_handler()
async def echo(message: types.Message):
    await message.answer(message.text)
