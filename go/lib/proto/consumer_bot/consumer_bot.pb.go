// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: consumer_bot.proto

package consumer_bot

import (
	common "github.com/nikhovas/diploma/go/lib/proto/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SimpleMessageInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgLocation *common.MsgLocation `protobuf:"bytes,1,opt,name=msgLocation,proto3" json:"msgLocation,omitempty"`
	Text        string              `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SimpleMessageInformation) Reset() {
	*x = SimpleMessageInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_bot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleMessageInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleMessageInformation) ProtoMessage() {}

func (x *SimpleMessageInformation) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_bot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleMessageInformation.ProtoReflect.Descriptor instead.
func (*SimpleMessageInformation) Descriptor() ([]byte, []int) {
	return file_consumer_bot_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleMessageInformation) GetMsgLocation() *common.MsgLocation {
	if x != nil {
		return x.MsgLocation
	}
	return nil
}

func (x *SimpleMessageInformation) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ReplyMessageInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgLocation          *common.MsgLocation `protobuf:"bytes,1,opt,name=msgLocation,proto3" json:"msgLocation,omitempty"`
	Text                 string              `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	ReplyMessageId       uint64              `protobuf:"varint,3,opt,name=replyMessageId,proto3" json:"replyMessageId,omitempty"`
	ReplyUnsupportedText string              `protobuf:"bytes,4,opt,name=replyUnsupportedText,proto3" json:"replyUnsupportedText,omitempty"`
}

func (x *ReplyMessageInformation) Reset() {
	*x = ReplyMessageInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_bot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyMessageInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyMessageInformation) ProtoMessage() {}

func (x *ReplyMessageInformation) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_bot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyMessageInformation.ProtoReflect.Descriptor instead.
func (*ReplyMessageInformation) Descriptor() ([]byte, []int) {
	return file_consumer_bot_proto_rawDescGZIP(), []int{1}
}

func (x *ReplyMessageInformation) GetMsgLocation() *common.MsgLocation {
	if x != nil {
		return x.MsgLocation
	}
	return nil
}

func (x *ReplyMessageInformation) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ReplyMessageInformation) GetReplyMessageId() uint64 {
	if x != nil {
		return x.ReplyMessageId
	}
	return 0
}

func (x *ReplyMessageInformation) GetReplyUnsupportedText() string {
	if x != nil {
		return x.ReplyUnsupportedText
	}
	return ""
}

type SendSimpleMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string                    `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Info *SimpleMessageInformation `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SendSimpleMessageRequest) Reset() {
	*x = SendSimpleMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_bot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSimpleMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSimpleMessageRequest) ProtoMessage() {}

func (x *SendSimpleMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_bot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSimpleMessageRequest.ProtoReflect.Descriptor instead.
func (*SendSimpleMessageRequest) Descriptor() ([]byte, []int) {
	return file_consumer_bot_proto_rawDescGZIP(), []int{2}
}

func (x *SendSimpleMessageRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SendSimpleMessageRequest) GetInfo() *SimpleMessageInformation {
	if x != nil {
		return x.Info
	}
	return nil
}

type SendReplyMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string                   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Info *ReplyMessageInformation `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SendReplyMessageRequest) Reset() {
	*x = SendReplyMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_bot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReplyMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReplyMessageRequest) ProtoMessage() {}

func (x *SendReplyMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_bot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReplyMessageRequest.ProtoReflect.Descriptor instead.
func (*SendReplyMessageRequest) Descriptor() ([]byte, []int) {
	return file_consumer_bot_proto_rawDescGZIP(), []int{3}
}

func (x *SendReplyMessageRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SendReplyMessageRequest) GetInfo() *ReplyMessageInformation {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_consumer_bot_proto protoreflect.FileDescriptor

var file_consumer_bot_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x62,
	0x6f, 0x74, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x65, 0x0a, 0x18, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x0b,
	0x6d, 0x73, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x73, 0x67, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x6d, 0x73, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0xc0, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x0b, 0x6d, 0x73, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4d, 0x73, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x6d,
	0x73, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x26,
	0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x55,
	0x6e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x55, 0x6e, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x22, 0x6a, 0x0a, 0x18, 0x53, 0x65,
	0x6e, 0x64, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x74, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x68, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x62,
	0x6f, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x32, 0xb4, 0x01, 0x0a, 0x08, 0x56, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x54, 0x0a,
	0x11, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x62, 0x6f,
	0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x5f, 0x62, 0x6f, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x6b, 0x68, 0x6f, 0x76, 0x61, 0x73, 0x2f, 0x64,
	0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consumer_bot_proto_rawDescOnce sync.Once
	file_consumer_bot_proto_rawDescData = file_consumer_bot_proto_rawDesc
)

func file_consumer_bot_proto_rawDescGZIP() []byte {
	file_consumer_bot_proto_rawDescOnce.Do(func() {
		file_consumer_bot_proto_rawDescData = protoimpl.X.CompressGZIP(file_consumer_bot_proto_rawDescData)
	})
	return file_consumer_bot_proto_rawDescData
}

var file_consumer_bot_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_consumer_bot_proto_goTypes = []interface{}{
	(*SimpleMessageInformation)(nil), // 0: consumer_bot.SimpleMessageInformation
	(*ReplyMessageInformation)(nil),  // 1: consumer_bot.ReplyMessageInformation
	(*SendSimpleMessageRequest)(nil), // 2: consumer_bot.SendSimpleMessageRequest
	(*SendReplyMessageRequest)(nil),  // 3: consumer_bot.SendReplyMessageRequest
	(*common.MsgLocation)(nil),       // 4: common.MsgLocation
	(*common.EmptyResponse)(nil),     // 5: common.EmptyResponse
}
var file_consumer_bot_proto_depIdxs = []int32{
	4, // 0: consumer_bot.SimpleMessageInformation.msgLocation:type_name -> common.MsgLocation
	4, // 1: consumer_bot.ReplyMessageInformation.msgLocation:type_name -> common.MsgLocation
	0, // 2: consumer_bot.SendSimpleMessageRequest.info:type_name -> consumer_bot.SimpleMessageInformation
	1, // 3: consumer_bot.SendReplyMessageRequest.info:type_name -> consumer_bot.ReplyMessageInformation
	2, // 4: consumer_bot.VkServer.SendSimpleMessage:input_type -> consumer_bot.SendSimpleMessageRequest
	3, // 5: consumer_bot.VkServer.SendReplyMessage:input_type -> consumer_bot.SendReplyMessageRequest
	5, // 6: consumer_bot.VkServer.SendSimpleMessage:output_type -> common.EmptyResponse
	5, // 7: consumer_bot.VkServer.SendReplyMessage:output_type -> common.EmptyResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_consumer_bot_proto_init() }
func file_consumer_bot_proto_init() {
	if File_consumer_bot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_consumer_bot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleMessageInformation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_consumer_bot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyMessageInformation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_consumer_bot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSimpleMessageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_consumer_bot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReplyMessageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_consumer_bot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_consumer_bot_proto_goTypes,
		DependencyIndexes: file_consumer_bot_proto_depIdxs,
		MessageInfos:      file_consumer_bot_proto_msgTypes,
	}.Build()
	File_consumer_bot_proto = out.File
	file_consumer_bot_proto_rawDesc = nil
	file_consumer_bot_proto_goTypes = nil
	file_consumer_bot_proto_depIdxs = nil
}
