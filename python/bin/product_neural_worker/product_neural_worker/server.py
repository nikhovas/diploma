from proto.question_worker import question_worker_pb2, question_worker_pb2_grpc
from processor import Processor


class QuestionWorkerServer(question_worker_pb2_grpc.QuestionWorkerServicer):
    def __init__(self, processor: Processor):
        self.processor = processor

    async def GetQuestionAnswer(self, request: question_worker_pb2.GetQuestionAnswerRequest, context):
        answer, answer_distance = await self.processor.get_answer(request.question, request.basePath)
        repeated_distance = self.processor.check_duplicate(request.question, request.previousQuestions)
        return question_worker_pb2.GetQuestionAnswerResponse(
            answer=answer,
            repeatedDistance=repeated_distance,
            answerDistance=answer_distance,
        )

    async def AddQuestion(self, request: question_worker_pb2.AddQuestionRequest, context):
        await self.processor.add(request.question, request.answer, request.basePath)
        return question_worker_pb2.AddQuestionResponse()
