# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: consumer_messages.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from proto.common import common_pb2 as common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x17\x63onsumer_messages.proto\x12\x11\x63onsumer_messages\x1a\x0c\x63ommon.proto\"B\n\nReplyBlock\x12\x16\n\x0ereplyMessageId\x18\x01 \x01(\x04\x12\x1c\n\x14replyUnsupportedText\x18\x02 \x01(\t\"\x96\x01\n\x11MessageToConsumer\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12(\n\x0bmsgLocation\x18\x02 \x01(\x0b\x32\x13.common.MsgLocation\x12\x0c\n\x04text\x18\x03 \x01(\t\x12\x31\n\x05reply\x18\x04 \x01(\x0b\x32\x1d.consumer_messages.ReplyBlockH\x00\x88\x01\x01\x42\x08\n\x06_replyB<Z:github.com/nikhovas/diploma/go/lib/proto/consumer_messagesb\x06proto3')



_REPLYBLOCK = DESCRIPTOR.message_types_by_name['ReplyBlock']
_MESSAGETOCONSUMER = DESCRIPTOR.message_types_by_name['MessageToConsumer']
ReplyBlock = _reflection.GeneratedProtocolMessageType('ReplyBlock', (_message.Message,), {
  'DESCRIPTOR' : _REPLYBLOCK,
  '__module__' : 'consumer_messages_pb2'
  # @@protoc_insertion_point(class_scope:consumer_messages.ReplyBlock)
  })
_sym_db.RegisterMessage(ReplyBlock)

MessageToConsumer = _reflection.GeneratedProtocolMessageType('MessageToConsumer', (_message.Message,), {
  'DESCRIPTOR' : _MESSAGETOCONSUMER,
  '__module__' : 'consumer_messages_pb2'
  # @@protoc_insertion_point(class_scope:consumer_messages.MessageToConsumer)
  })
_sym_db.RegisterMessage(MessageToConsumer)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z:github.com/nikhovas/diploma/go/lib/proto/consumer_messages'
  _REPLYBLOCK._serialized_start=60
  _REPLYBLOCK._serialized_end=126
  _MESSAGETOCONSUMER._serialized_start=129
  _MESSAGETOCONSUMER._serialized_end=279
# @@protoc_insertion_point(module_scope)