import grpc_server.QuestionWorker_pb2_grpc as QuestionWorker_pb2_grpc
import grpc_server.QuestionWorker_pb2 as QuestionWorker_pb2
from processor import Processor


class QuestionWorkerServer(QuestionWorker_pb2_grpc.QuestionWorkerServicer):
    def __init__(self, processor: Processor):
        self.processor = processor

    async def GetQuestionAnswer(self, request: QuestionWorker_pb2.GetQuestionAnswerRequest, context):
        answer, answer_distance = await self.processor.get_answer(request.question, request.basePath)
        repeated_distance = self.processor.check_duplicate(request.question, request.previousQuestions)
        return QuestionWorker_pb2.GetQuestionAnswerResponse(
            answer=answer,
            repeatedDistance=repeated_distance,
            answerDistance=answer_distance,
        )

    async def AddQuestion(self, request, context):
        await self.processor.add(request.question, request.answer, request.basePath)
        return QuestionWorker_pb2.AddQuestionResponse()
