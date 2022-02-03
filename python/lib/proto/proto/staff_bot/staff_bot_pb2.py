# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: staff_bot.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.common import common_pb2 as common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='staff_bot.proto',
  package='staff_bot',
  syntax='proto3',
  serialized_options=b'Z2github.com/nikhovas/diploma/go/lib/proto/staff_bot',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0fstaff_bot.proto\x12\tstaff_bot\x1a\x0c\x63ommon.proto\"7\n\x12NewQuestionRequest\x12\x0f\n\x07groupId\x18\x01 \x01(\x03\x12\x10\n\x08question\x18\x02 \x01(\t\"H\n$NotifyBotStatusChangeTelegramRequest\x12\x0f\n\x07groupId\x18\x01 \x01(\x03\x12\x0f\n\x07\x65nabled\x18\x02 \x01(\x08\x32\xc8\x01\n\x10TelegramStaffBot\x12I\n\x0fSendNewQuestion\x12\x1d.staff_bot.NewQuestionRequest\x1a\x15.common.EmptyResponse\"\x00\x12i\n\x1dNotifyBotStatusTelegramChange\x12/.staff_bot.NotifyBotStatusChangeTelegramRequest\x1a\x15.common.EmptyResponse\"\x00\x42\x34Z2github.com/nikhovas/diploma/go/lib/proto/staff_botb\x06proto3'
  ,
  dependencies=[common__pb2.DESCRIPTOR,])




_NEWQUESTIONREQUEST = _descriptor.Descriptor(
  name='NewQuestionRequest',
  full_name='staff_bot.NewQuestionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='groupId', full_name='staff_bot.NewQuestionRequest.groupId', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='question', full_name='staff_bot.NewQuestionRequest.question', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=44,
  serialized_end=99,
)


_NOTIFYBOTSTATUSCHANGETELEGRAMREQUEST = _descriptor.Descriptor(
  name='NotifyBotStatusChangeTelegramRequest',
  full_name='staff_bot.NotifyBotStatusChangeTelegramRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='groupId', full_name='staff_bot.NotifyBotStatusChangeTelegramRequest.groupId', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='enabled', full_name='staff_bot.NotifyBotStatusChangeTelegramRequest.enabled', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=101,
  serialized_end=173,
)

DESCRIPTOR.message_types_by_name['NewQuestionRequest'] = _NEWQUESTIONREQUEST
DESCRIPTOR.message_types_by_name['NotifyBotStatusChangeTelegramRequest'] = _NOTIFYBOTSTATUSCHANGETELEGRAMREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

NewQuestionRequest = _reflection.GeneratedProtocolMessageType('NewQuestionRequest', (_message.Message,), {
  'DESCRIPTOR' : _NEWQUESTIONREQUEST,
  '__module__' : 'staff_bot_pb2'
  # @@protoc_insertion_point(class_scope:staff_bot.NewQuestionRequest)
  })
_sym_db.RegisterMessage(NewQuestionRequest)

NotifyBotStatusChangeTelegramRequest = _reflection.GeneratedProtocolMessageType('NotifyBotStatusChangeTelegramRequest', (_message.Message,), {
  'DESCRIPTOR' : _NOTIFYBOTSTATUSCHANGETELEGRAMREQUEST,
  '__module__' : 'staff_bot_pb2'
  # @@protoc_insertion_point(class_scope:staff_bot.NotifyBotStatusChangeTelegramRequest)
  })
_sym_db.RegisterMessage(NotifyBotStatusChangeTelegramRequest)


DESCRIPTOR._options = None

_TELEGRAMSTAFFBOT = _descriptor.ServiceDescriptor(
  name='TelegramStaffBot',
  full_name='staff_bot.TelegramStaffBot',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=176,
  serialized_end=376,
  methods=[
  _descriptor.MethodDescriptor(
    name='SendNewQuestion',
    full_name='staff_bot.TelegramStaffBot.SendNewQuestion',
    index=0,
    containing_service=None,
    input_type=_NEWQUESTIONREQUEST,
    output_type=common__pb2._EMPTYRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='NotifyBotStatusTelegramChange',
    full_name='staff_bot.TelegramStaffBot.NotifyBotStatusTelegramChange',
    index=1,
    containing_service=None,
    input_type=_NOTIFYBOTSTATUSCHANGETELEGRAMREQUEST,
    output_type=common__pb2._EMPTYRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_TELEGRAMSTAFFBOT)

DESCRIPTOR.services_by_name['TelegramStaffBot'] = _TELEGRAMSTAFFBOT

# @@protoc_insertion_point(module_scope)