# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: question_worker.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15question_worker.proto\x12\x0fquestion_worker\"g\n\x18GetQuestionAnswerRequest\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x10\n\x08question\x18\x02 \x01(\t\x12\x19\n\x11previousQuestions\x18\x03 \x03(\t\x12\x10\n\x08\x62\x61sePath\x18\x04 \x01(\t\"]\n\x19GetQuestionAnswerResponse\x12\x0e\n\x06\x61nswer\x18\x01 \x01(\t\x12\x18\n\x10repeatedDistance\x18\x02 \x01(\x02\x12\x16\n\x0e\x61nswerDistance\x18\x03 \x01(\x02\"V\n\x12\x41\x64\x64QuestionRequest\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x10\n\x08question\x18\x02 \x01(\t\x12\x0e\n\x06\x61nswer\x18\x03 \x01(\t\x12\x10\n\x08\x62\x61sePath\x18\x04 \x01(\t\"\x15\n\x13\x41\x64\x64QuestionResponse2\xda\x01\n\x0eQuestionWorker\x12l\n\x11GetQuestionAnswer\x12).question_worker.GetQuestionAnswerRequest\x1a*.question_worker.GetQuestionAnswerResponse\"\x00\x12Z\n\x0b\x41\x64\x64Question\x12#.question_worker.AddQuestionRequest\x1a$.question_worker.AddQuestionResponse\"\x00\x42:Z8github.com/nikhovas/diploma/go/lib/proto/question_workerb\x06proto3')



_GETQUESTIONANSWERREQUEST = DESCRIPTOR.message_types_by_name['GetQuestionAnswerRequest']
_GETQUESTIONANSWERRESPONSE = DESCRIPTOR.message_types_by_name['GetQuestionAnswerResponse']
_ADDQUESTIONREQUEST = DESCRIPTOR.message_types_by_name['AddQuestionRequest']
_ADDQUESTIONRESPONSE = DESCRIPTOR.message_types_by_name['AddQuestionResponse']
GetQuestionAnswerRequest = _reflection.GeneratedProtocolMessageType('GetQuestionAnswerRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETQUESTIONANSWERREQUEST,
  '__module__' : 'question_worker_pb2'
  # @@protoc_insertion_point(class_scope:question_worker.GetQuestionAnswerRequest)
  })
_sym_db.RegisterMessage(GetQuestionAnswerRequest)

GetQuestionAnswerResponse = _reflection.GeneratedProtocolMessageType('GetQuestionAnswerResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETQUESTIONANSWERRESPONSE,
  '__module__' : 'question_worker_pb2'
  # @@protoc_insertion_point(class_scope:question_worker.GetQuestionAnswerResponse)
  })
_sym_db.RegisterMessage(GetQuestionAnswerResponse)

AddQuestionRequest = _reflection.GeneratedProtocolMessageType('AddQuestionRequest', (_message.Message,), {
  'DESCRIPTOR' : _ADDQUESTIONREQUEST,
  '__module__' : 'question_worker_pb2'
  # @@protoc_insertion_point(class_scope:question_worker.AddQuestionRequest)
  })
_sym_db.RegisterMessage(AddQuestionRequest)

AddQuestionResponse = _reflection.GeneratedProtocolMessageType('AddQuestionResponse', (_message.Message,), {
  'DESCRIPTOR' : _ADDQUESTIONRESPONSE,
  '__module__' : 'question_worker_pb2'
  # @@protoc_insertion_point(class_scope:question_worker.AddQuestionResponse)
  })
_sym_db.RegisterMessage(AddQuestionResponse)

_QUESTIONWORKER = DESCRIPTOR.services_by_name['QuestionWorker']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z8github.com/nikhovas/diploma/go/lib/proto/question_worker'
  _GETQUESTIONANSWERREQUEST._serialized_start=42
  _GETQUESTIONANSWERREQUEST._serialized_end=145
  _GETQUESTIONANSWERRESPONSE._serialized_start=147
  _GETQUESTIONANSWERRESPONSE._serialized_end=240
  _ADDQUESTIONREQUEST._serialized_start=242
  _ADDQUESTIONREQUEST._serialized_end=328
  _ADDQUESTIONRESPONSE._serialized_start=330
  _ADDQUESTIONRESPONSE._serialized_end=351
  _QUESTIONWORKER._serialized_start=354
  _QUESTIONWORKER._serialized_end=572
# @@protoc_insertion_point(module_scope)
