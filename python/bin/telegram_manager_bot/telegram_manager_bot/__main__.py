import asyncio
from telegram_manager_bot import telegram
from telegram_manager_bot.server import grpc_server


if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    asyncio.ensure_future(telegram.dp.start_polling())
    asyncio.ensure_future(grpc_server(telegram.bot))
    loop.run_forever()
