from typing import AsyncIterator

import aioredis

from question_worker.service_types import Qa
from question_worker.storage import StorageInterface


class Redis(StorageInterface):
    def __init__(self, host: str, password: str):
        # hosts = [host]
        # sentinel = Sentinel([(h, 5000) for h in hosts], socket_timeout=0.1)
        # self.master = sentinel.master_for('redis883')
        # self.slave = sentinel.slave_for('redis883')
        self.master = aioredis.StrictRedis(
            host='localhost',
            port=6379,
            password='',
        )
        self.slave = self.master

    async def get(self, key: str, path: str) -> Qa:
        val = await self.slave.get(path + '/' + key).decode('utf-8')
        return Qa.from_json(val)

    async def set(self, qa: Qa, path: str):
        json_data = qa.to_json()
        await self.master.set(path + '/' + qa.question, json_data)

    async def get_all(self, path: str) -> AsyncIterator[Qa]:
        async for key in self.slave.scan_iter(path + '/*'):
            val = await self.slave.get(key)
            val = val.decode('utf-8')
            yield Qa.from_json(val)
