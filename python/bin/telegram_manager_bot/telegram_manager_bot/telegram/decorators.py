from typing import Optional

from aiogram import types
from proto.controller import controller_pb2
from telegram_manager_bot.telegram import dp


class BotError(Exception):
    def __init__(self, grpc_resp):
        self.grpc_resp = grpc_resp


def action(cmd_name: str, separator: Optional[str] = None, args_count: Optional[int] = None):
    def decorator(f):
        async def internal(message: types.Message):
            cmi = controller_pb2.TelegramMessageInfo(
                chatId=message.chat.id,
                isChatPrivate=(message.chat.type == 'private'),
                userId=message.from_user.id,
            )

            text = message.text
            if separator is not None:
                tokens = text.split(separator)

                # check for tokens in first line
                first_line = tokens[0]
                space_pos = first_line.find(' ')
                if space_pos != -1:
                    tokens[0] = first_line[space_pos + 1:]
                else:
                    tokens = tokens[1:]

                if args_count is not None and args_count != len(tokens):
                    await message.answer(f'У команды должно быть {args_count} аргументов')
                    return
            else:
                tokens = [message.text[len(cmd_name) + 1:]]

            try:
                await f(message, cmi, *tokens)
            except BotError as err:
                await message.answer('asdf')

        dp.register_message_handler(internal, commands=[cmd_name])
        return internal

    return decorator


def check_resp_for_error(resp):
    pass


def group_only(f):
    async def internal(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
        if cmi.isChatPrivate is True:
            raise BotError(controller_pb2.DefaultResponse(
                noRoleError=controller_pb2.BadChatType(shouldBePrivate=False)),
            )
        await f(message, cmi)

    return internal


def private_only(f):
    async def internal(message: types.Message, cmi: controller_pb2.TelegramMessageInfo):
        if cmi.isChatPrivate is False:
            raise BotError(controller_pb2.DefaultResponse(
                noRoleError=controller_pb2.BadChatType(shouldBePrivate=True)),
            )
        await f(message, cmi)

    return internal
