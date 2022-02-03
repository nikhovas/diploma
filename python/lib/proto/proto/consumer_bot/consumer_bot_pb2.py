# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: consumer_bot.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.common import common_pb2 as common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='consumer_bot.proto',
  package='consumer_bot',
  syntax='proto3',
  serialized_options=b'Z5github.com/nikhovas/diploma/go/lib/proto/consumer_bot',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x12\x63onsumer_bot.proto\x12\x0c\x63onsumer_bot\x1a\x0c\x63ommon.proto\"F\n\x12MessageDestination\x12\x0f\n\x07service\x18\x01 \x01(\t\x12\x0f\n\x07groupId\x18\x02 \x01(\x03\x12\x0e\n\x06userId\x18\x03 \x01(\x03\"f\n\x18SimpleMessageInformation\x12<\n\x12messageDestination\x18\x01 \x01(\x0b\x32 .consumer_bot.MessageDestination\x12\x0c\n\x04text\x18\x02 \x01(\t\"}\n\x17ReplyMessageInformation\x12<\n\x12messageDestination\x18\x01 \x01(\x0b\x32 .consumer_bot.MessageDestination\x12\x0c\n\x04text\x18\x02 \x01(\t\x12\x16\n\x0ereplyMessageId\x18\x03 \x01(\x04\"B\n\x11\x42otsActionRequest\x12\x0f\n\x07groupId\x18\x01 \x01(\x03\x12\x12\n\x05token\x18\x02 \x01(\tH\x00\x88\x01\x01\x42\x08\n\x06_token*$\n\rBotActionEnum\x12\x07\n\x03\x41\x44\x44\x10\x00\x12\n\n\x06REMOVE\x10\x01\x32\xbf\x02\n\x08VkServer\x12T\n\x11SendSimpleMessage\x12&.consumer_bot.SimpleMessageInformation\x1a\x15.common.EmptyResponse\"\x00\x12R\n\x10SendReplyMessage\x12%.consumer_bot.ReplyMessageInformation\x1a\x15.common.EmptyResponse\"\x00\x12\x42\n\x06\x41\x64\x64\x42ot\x12\x1f.consumer_bot.BotsActionRequest\x1a\x15.common.EmptyResponse\"\x00\x12\x45\n\tRemoveBot\x12\x1f.consumer_bot.BotsActionRequest\x1a\x15.common.EmptyResponse\"\x00\x42\x37Z5github.com/nikhovas/diploma/go/lib/proto/consumer_botb\x06proto3'
  ,
  dependencies=[common__pb2.DESCRIPTOR,])

_BOTACTIONENUM = _descriptor.EnumDescriptor(
  name='BotActionEnum',
  full_name='consumer_bot.BotActionEnum',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ADD', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='REMOVE', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=421,
  serialized_end=457,
)
_sym_db.RegisterEnumDescriptor(_BOTACTIONENUM)

BotActionEnum = enum_type_wrapper.EnumTypeWrapper(_BOTACTIONENUM)
ADD = 0
REMOVE = 1



_MESSAGEDESTINATION = _descriptor.Descriptor(
  name='MessageDestination',
  full_name='consumer_bot.MessageDestination',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='service', full_name='consumer_bot.MessageDestination.service', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='groupId', full_name='consumer_bot.MessageDestination.groupId', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='userId', full_name='consumer_bot.MessageDestination.userId', index=2,
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
  serialized_start=50,
  serialized_end=120,
)


_SIMPLEMESSAGEINFORMATION = _descriptor.Descriptor(
  name='SimpleMessageInformation',
  full_name='consumer_bot.SimpleMessageInformation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='messageDestination', full_name='consumer_bot.SimpleMessageInformation.messageDestination', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='text', full_name='consumer_bot.SimpleMessageInformation.text', index=1,
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
  serialized_start=122,
  serialized_end=224,
)


_REPLYMESSAGEINFORMATION = _descriptor.Descriptor(
  name='ReplyMessageInformation',
  full_name='consumer_bot.ReplyMessageInformation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='messageDestination', full_name='consumer_bot.ReplyMessageInformation.messageDestination', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='text', full_name='consumer_bot.ReplyMessageInformation.text', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='replyMessageId', full_name='consumer_bot.ReplyMessageInformation.replyMessageId', index=2,
      number=3, type=4, cpp_type=4, label=1,
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
  serialized_start=226,
  serialized_end=351,
)


_BOTSACTIONREQUEST = _descriptor.Descriptor(
  name='BotsActionRequest',
  full_name='consumer_bot.BotsActionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='groupId', full_name='consumer_bot.BotsActionRequest.groupId', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='token', full_name='consumer_bot.BotsActionRequest.token', index=1,
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
    _descriptor.OneofDescriptor(
      name='_token', full_name='consumer_bot.BotsActionRequest._token',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=353,
  serialized_end=419,
)

_SIMPLEMESSAGEINFORMATION.fields_by_name['messageDestination'].message_type = _MESSAGEDESTINATION
_REPLYMESSAGEINFORMATION.fields_by_name['messageDestination'].message_type = _MESSAGEDESTINATION
_BOTSACTIONREQUEST.oneofs_by_name['_token'].fields.append(
  _BOTSACTIONREQUEST.fields_by_name['token'])
_BOTSACTIONREQUEST.fields_by_name['token'].containing_oneof = _BOTSACTIONREQUEST.oneofs_by_name['_token']
DESCRIPTOR.message_types_by_name['MessageDestination'] = _MESSAGEDESTINATION
DESCRIPTOR.message_types_by_name['SimpleMessageInformation'] = _SIMPLEMESSAGEINFORMATION
DESCRIPTOR.message_types_by_name['ReplyMessageInformation'] = _REPLYMESSAGEINFORMATION
DESCRIPTOR.message_types_by_name['BotsActionRequest'] = _BOTSACTIONREQUEST
DESCRIPTOR.enum_types_by_name['BotActionEnum'] = _BOTACTIONENUM
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

MessageDestination = _reflection.GeneratedProtocolMessageType('MessageDestination', (_message.Message,), {
  'DESCRIPTOR' : _MESSAGEDESTINATION,
  '__module__' : 'consumer_bot_pb2'
  # @@protoc_insertion_point(class_scope:consumer_bot.MessageDestination)
  })
_sym_db.RegisterMessage(MessageDestination)

SimpleMessageInformation = _reflection.GeneratedProtocolMessageType('SimpleMessageInformation', (_message.Message,), {
  'DESCRIPTOR' : _SIMPLEMESSAGEINFORMATION,
  '__module__' : 'consumer_bot_pb2'
  # @@protoc_insertion_point(class_scope:consumer_bot.SimpleMessageInformation)
  })
_sym_db.RegisterMessage(SimpleMessageInformation)

ReplyMessageInformation = _reflection.GeneratedProtocolMessageType('ReplyMessageInformation', (_message.Message,), {
  'DESCRIPTOR' : _REPLYMESSAGEINFORMATION,
  '__module__' : 'consumer_bot_pb2'
  # @@protoc_insertion_point(class_scope:consumer_bot.ReplyMessageInformation)
  })
_sym_db.RegisterMessage(ReplyMessageInformation)

BotsActionRequest = _reflection.GeneratedProtocolMessageType('BotsActionRequest', (_message.Message,), {
  'DESCRIPTOR' : _BOTSACTIONREQUEST,
  '__module__' : 'consumer_bot_pb2'
  # @@protoc_insertion_point(class_scope:consumer_bot.BotsActionRequest)
  })
_sym_db.RegisterMessage(BotsActionRequest)


DESCRIPTOR._options = None

_VKSERVER = _descriptor.ServiceDescriptor(
  name='VkServer',
  full_name='consumer_bot.VkServer',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=460,
  serialized_end=779,
  methods=[
  _descriptor.MethodDescriptor(
    name='SendSimpleMessage',
    full_name='consumer_bot.VkServer.SendSimpleMessage',
    index=0,
    containing_service=None,
    input_type=_SIMPLEMESSAGEINFORMATION,
    output_type=common__pb2._EMPTYRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='SendReplyMessage',
    full_name='consumer_bot.VkServer.SendReplyMessage',
    index=1,
    containing_service=None,
    input_type=_REPLYMESSAGEINFORMATION,
    output_type=common__pb2._EMPTYRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='AddBot',
    full_name='consumer_bot.VkServer.AddBot',
    index=2,
    containing_service=None,
    input_type=_BOTSACTIONREQUEST,
    output_type=common__pb2._EMPTYRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='RemoveBot',
    full_name='consumer_bot.VkServer.RemoveBot',
    index=3,
    containing_service=None,
    input_type=_BOTSACTIONREQUEST,
    output_type=common__pb2._EMPTYRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_VKSERVER)

DESCRIPTOR.services_by_name['VkServer'] = _VKSERVER

# @@protoc_insertion_point(module_scope)