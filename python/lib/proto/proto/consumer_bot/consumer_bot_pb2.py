# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: consumer_bot.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.common import common_pb2 as common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x12\x63onsumer_bot.proto\x12\x0c\x63onsumer_bot\x1a\x0c\x63ommon.proto\"R\n\x18SimpleMessageInformation\x12(\n\x0bmsgLocation\x18\x01 \x01(\x0b\x32\x13.common.MsgLocation\x12\x0c\n\x04text\x18\x02 \x01(\t\"\x87\x01\n\x17ReplyMessageInformation\x12(\n\x0bmsgLocation\x18\x01 \x01(\x0b\x32\x13.common.MsgLocation\x12\x0c\n\x04text\x18\x02 \x01(\t\x12\x16\n\x0ereplyMessageId\x18\x03 \x01(\x04\x12\x1c\n\x14replyUnsupportedText\x18\x04 \x01(\t\"^\n\x18SendSimpleMessageRequest\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x34\n\x04info\x18\x02 \x01(\x0b\x32&.consumer_bot.SimpleMessageInformation\"\\\n\x17SendReplyMessageRequest\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x33\n\x04info\x18\x02 \x01(\x0b\x32%.consumer_bot.ReplyMessageInformation2\xb4\x01\n\x08VkServer\x12T\n\x11SendSimpleMessage\x12&.consumer_bot.SendSimpleMessageRequest\x1a\x15.common.EmptyResponse\"\x00\x12R\n\x10SendReplyMessage\x12%.consumer_bot.SendReplyMessageRequest\x1a\x15.common.EmptyResponse\"\x00\x42\x37Z5github.com/nikhovas/diploma/go/lib/proto/consumer_botb\x06proto3')



_SIMPLEMESSAGEINFORMATION = DESCRIPTOR.message_types_by_name['SimpleMessageInformation']
_REPLYMESSAGEINFORMATION = DESCRIPTOR.message_types_by_name['ReplyMessageInformation']
_SENDSIMPLEMESSAGEREQUEST = DESCRIPTOR.message_types_by_name['SendSimpleMessageRequest']
_SENDREPLYMESSAGEREQUEST = DESCRIPTOR.message_types_by_name['SendReplyMessageRequest']
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

SendSimpleMessageRequest = _reflection.GeneratedProtocolMessageType('SendSimpleMessageRequest', (_message.Message,), {
  'DESCRIPTOR' : _SENDSIMPLEMESSAGEREQUEST,
  '__module__' : 'consumer_bot_pb2'
  # @@protoc_insertion_point(class_scope:consumer_bot.SendSimpleMessageRequest)
  })
_sym_db.RegisterMessage(SendSimpleMessageRequest)

SendReplyMessageRequest = _reflection.GeneratedProtocolMessageType('SendReplyMessageRequest', (_message.Message,), {
  'DESCRIPTOR' : _SENDREPLYMESSAGEREQUEST,
  '__module__' : 'consumer_bot_pb2'
  # @@protoc_insertion_point(class_scope:consumer_bot.SendReplyMessageRequest)
  })
_sym_db.RegisterMessage(SendReplyMessageRequest)

_VKSERVER = DESCRIPTOR.services_by_name['VkServer']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z5github.com/nikhovas/diploma/go/lib/proto/consumer_bot'
  _SIMPLEMESSAGEINFORMATION._serialized_start=50
  _SIMPLEMESSAGEINFORMATION._serialized_end=132
  _REPLYMESSAGEINFORMATION._serialized_start=135
  _REPLYMESSAGEINFORMATION._serialized_end=270
  _SENDSIMPLEMESSAGEREQUEST._serialized_start=272
  _SENDSIMPLEMESSAGEREQUEST._serialized_end=366
  _SENDREPLYMESSAGEREQUEST._serialized_start=368
  _SENDREPLYMESSAGEREQUEST._serialized_end=460
  _VKSERVER._serialized_start=463
  _VKSERVER._serialized_end=643
# @@protoc_insertion_point(module_scope)
