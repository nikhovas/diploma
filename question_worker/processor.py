from typing import List

from neural import Model
from storage import StorageInterface


class Processor:
    def __init__(self, storage: StorageInterface, model: Model):
        self.storage = storage
        self.model = model

    async def get_answer(self, question: str, base_path: str) -> (str, float):
        min_q = None
        min_dist = float('inf')
        question_text = self.model.new_text(question)

        async for qa in self.storage.get_all(base_path):
            dist = self.model.get_distance(qa, question_text)
            if dist < min_dist:
                min_q = qa
                min_dist = dist

        if min_q is None:
            return '', min_dist
        else:
            return min_q.answer, min_dist

    def check_duplicate(self, question: str, previous_questions: List[str]) -> float:
        min_dist = float('inf')
        q_text = self.model.new_text(question)

        for pq in previous_questions:
            dist = self.model.get_distance(q_text, self.model.new_text(pq))
            min_dist = min(dist, min_dist)

        return min_dist

    async def add(self, question: str, answer: str, base_path: str):
        qa = self.model.new_qa(question, answer)
        await self.storage.set(qa, base_path)
