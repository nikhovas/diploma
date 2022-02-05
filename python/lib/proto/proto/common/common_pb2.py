# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='common.proto',
  package='common',
  syntax='proto3',
  serialized_options=b'Z/github.com/nikhovas/diploma/go/lib/proto/common',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0c\x63ommon.proto\x12\x06\x63ommon\"\x0f\n\rEmptyResponse\"?\n\x0bMsgLocation\x12\x0f\n\x07service\x18\x01 \x01(\t\x12\x0f\n\x07groupId\x18\x02 \x01(\x03\x12\x0e\n\x06userId\x18\x03 \x01(\x03\"n\n\x19WaitingQuesionInformation\x12(\n\x0bmsgLocation\x18\x01 \x01(\x0b\x32\x13.common.MsgLocation\x12\x10\n\x08question\x18\x02 \x01(\t\x12\x15\n\rquestionMsgId\x18\x03 \x01(\x03\x42\x31Z/github.com/nikhovas/diploma/go/lib/proto/commonb\x06proto3'
)




_EMPTYRESPONSE = _descriptor.Descriptor(
  name='EmptyResponse',
  full_name='common.EmptyResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
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
  serialized_start=24,
  serialized_end=39,
)


_MSGLOCATION = _descriptor.Descriptor(
  name='MsgLocation',
  full_name='common.MsgLocation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='service', full_name='common.MsgLocation.service', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='groupId', full_name='common.MsgLocation.groupId', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='userId', full_name='common.MsgLocation.userId', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=41,
  serialized_end=104,
)


_WAITINGQUESIONINFORMATION = _descriptor.Descriptor(
  name='WaitingQuesionInformation',
  full_name='common.WaitingQuesionInformation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='msgLocation', full_name='common.WaitingQuesionInformation.msgLocation', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='question', full_name='common.WaitingQuesionInformation.question', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='questionMsgId', full_name='common.WaitingQuesionInformation.questionMsgId', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=106,
  serialized_end=216,
)

_WAITINGQUESIONINFORMATION.fields_by_name['msgLocation'].message_type = _MSGLOCATION
DESCRIPTOR.message_types_by_name['EmptyResponse'] = _EMPTYRESPONSE
DESCRIPTOR.message_types_by_name['MsgLocation'] = _MSGLOCATION
DESCRIPTOR.message_types_by_name['WaitingQuesionInformation'] = _WAITINGQUESIONINFORMATION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

EmptyResponse = _reflection.GeneratedProtocolMessageType('EmptyResponse', (_message.Message,), {
  'DESCRIPTOR' : _EMPTYRESPONSE,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:common.EmptyResponse)
  })
_sym_db.RegisterMessage(EmptyResponse)

MsgLocation = _reflection.GeneratedProtocolMessageType('MsgLocation', (_message.Message,), {
  'DESCRIPTOR' : _MSGLOCATION,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:common.MsgLocation)
  })
_sym_db.RegisterMessage(MsgLocation)

WaitingQuesionInformation = _reflection.GeneratedProtocolMessageType('WaitingQuesionInformation', (_message.Message,), {
  'DESCRIPTOR' : _WAITINGQUESIONINFORMATION,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:common.WaitingQuesionInformation)
  })
_sym_db.RegisterMessage(WaitingQuesionInformation)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
