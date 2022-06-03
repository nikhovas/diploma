from telegram_manager_bot.telegram import dp, decorators, stub, QUESTION_MARKER
from proto.controller import controller_pb2
from aiogram import types


def shop_key_from_chat(chat_id):
    return controller_pb2.ShopKey(telegramStaff=controller_pb2.TelegramStaffShopKey(groupId=chat_id))


@decorators.action('add_question', '\n', 2)
@decorators.group_only
async def add_question(uuid, message: types.Message, cmi: controller_pb2.TelegramMessageInfo, question: str, answer: str):
    response = await stub.AddQuestionAnswer(controller_pb2.AddQuestionAnswerRequest(
        uuid='',
        messageInfo=controller_pb2.MessageInformation(telegram=cmi),
        key=shop_key_from_chat(cmi.chatId),
        question=question,
        answer=answer,
        checkQuestionInSet=False,
    ))

    await message.reply('Ваш вопрос добавлен')


@decorators.action('list_shops', ' ', 0)
@decorators.private_only
async def list_shops(uuid, message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    response = await stub.ListShops(controller_pb2.ListShopsRequest(
        uuid=uuid,
        messageInfo=controller_pb2.MessageInformation(telegram=cmi),
    ))

    if response.success is not None:
        result_text = '**Ваши боты**\n'
        for bot in response.success.bots:
            result_text += f'* {bot.name}\n'

        await message.answer(result_text, parse_mode='Markdown')
    else:
        await message.answer('Прозошла ошибка')


@decorators.action('add_shop', ' ', 3)
@decorators.private_only
async def add_shop(uuid, message: types.Message, cmi: controller_pb2.TelegramMessageInfo, name, vk_token, vk_group_id):
    vk_group_id = int(vk_group_id)

    response = await stub.AddShop(controller_pb2.AddShopRequest(
        uuid=uuid,
        messageInfo=controller_pb2.MessageInformation(telegram=cmi),
        bot=controller_pb2.BotInfo(
            commonBotInfo=controller_pb2.CommonBotInfo(name=name, token=vk_token, groupId=vk_group_id),
            platformBotInfo=controller_pb2.PlatformBotInfo(telegram=controller_pb2.TelegramBotInfo(chatId=0)),
        )
    ))

    await message.answer('Бот магазина добавлен')


@decorators.action('answer')
@decorators.group_only
async def answer_message(uuid, message: types.Message, cmi: controller_pb2.TelegramMessageInfo, text: str):
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


@decorators.action('delete_shop', ' ', 1)
@decorators.private_only
async def delete_shop(uuid, message: types.Message, cmi: controller_pb2.TelegramMessageInfo, shop_name):
    # check for role

    # start

    # send answer to client

    await message.answer('Бот магазина запущен')


@decorators.action('modify_shop', ' ')
@decorators.private_only
async def modify_shop(uuid, message: types.Message, cmi: controller_pb2.TelegramMessageInfo, *args):
    # check for role

    # start

    # send answer to client

    await message.answer('Бот магазина запущен')


@decorators.action('register_group', ' ', 1)
@decorators.group_only
async def register_shop_group(uuid, message: types.Message, cmi: controller_pb2.TelegramMessageInfo, shop_name):
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


@decorators.action('start_shop_bot', ' ', 0)
@decorators.group_only
async def start_shop_bot(uuid, message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    response = await stub.ChangeBotState(controller_pb2.ChangeBotStateRequest(
        messageInfo=controller_pb2.MessageInformation(telegram=cmi),
        key=controller_pb2.ShopKey(telegramStaff=controller_pb2.TelegramStaffShopKey(groupId=message.chat.id)),
        toEnabled=True,
    ))
    await message.reply('🕐 Отправлен запрос на включение бота')


@decorators.action('stop_shop_bot', ' ', 0)
@decorators.group_only
async def stop_shop_bot(uuid, message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    response = await stub.ChangeBotState(controller_pb2.ChangeBotStateRequest(
        messageInfo=controller_pb2.MessageInformation(telegram=cmi),
        key=controller_pb2.ShopKey(telegramStaff=controller_pb2.TelegramStaffShopKey(groupId=message.chat.id)),
        toEnabled=False,
    ))
    await message.reply('🕐 Отправлен запрос на остановку бота')


@dp.message_handler(commands=['start', 'help'])
async def send_welcome(uuid, message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
    print(message.from_user)
    print(message.chat)
    await message.reply("Hi!\nI'm EchoBot!\nPowered by aiogram.")


@dp.message_handler()
async def echo(message: types.Message):
    await message.answer(message.text)
