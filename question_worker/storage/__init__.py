from abc import ABC, abstractmethod
from typing import AsyncIterator

from service_types import Qa


class StorageInterface(ABC):
    @abstractmethod
    async def get(self, key: str, path: str) -> Qa:
        pass

    @abstractmethod
    async def set(self, qa: Qa, path: str):
        pass

    @abstractmethod
    async def get_all(self, path: str) -> AsyncIterator[Qa]:
        pass
