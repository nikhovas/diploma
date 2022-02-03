import asyncio
import logging

from grpc import aio

from question_worker.server import QuestionWorkerServer
from proto.question_worker import question_worker_pb2_grpc
from neural import Model
from question_worker.processor import Processor
from question_worker.storage.redis import Redis


async def serve():
    storage = Redis(host='127.0.0.1', password='redis_password')
    model = Model()
    processor = Processor(storage, model)

    server = aio.server()
    question_worker_pb2_grpc.add_QuestionWorkerServicer_to_server(QuestionWorkerServer(processor), server)
    listen_addr = '[::]:50051'
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()
    await server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    asyncio.run(serve())
