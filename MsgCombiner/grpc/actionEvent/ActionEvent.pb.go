// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: ActionEvent.proto

package actions

import (
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

type ActionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	BotId       string `protobuf:"bytes,2,opt,name=botId,proto3" json:"botId,omitempty"`
	ShopId      int64  `protobuf:"varint,3,opt,name=shopId,proto3" json:"shopId,omitempty"`
	Time        uint64 `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	ServiceName string `protobuf:"bytes,5,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	NeedOrder   bool   `protobuf:"varint,6,opt,name=needOrder,proto3" json:"needOrder,omitempty"`
}

func (x *ActionEvent) Reset() {
	*x = ActionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActionEvent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionEvent) ProtoMessage() {}

func (x *ActionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_ActionEvent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionEvent.ProtoReflect.Descriptor instead.
func (*ActionEvent) Descriptor() ([]byte, []int) {
	return file_ActionEvent_proto_rawDescGZIP(), []int{0}
}

func (x *ActionEvent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ActionEvent) GetBotId() string {
	if x != nil {
		return x.BotId
	}
	return ""
}

func (x *ActionEvent) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *ActionEvent) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *ActionEvent) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *ActionEvent) GetNeedOrder() bool {
	if x != nil {
		return x.NeedOrder
	}
	return false
}

var File_ActionEvent_proto protoreflect.FileDescriptor

var file_ActionEvent_proto_rawDesc = []byte{
	0x0a, 0x11, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x6f, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x74, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x6e, 0x65, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x6b, 0x68,
	0x6f, 0x76, 0x61, 0x73, 0x2f, 0x64, 0x69, 0x70, 0x6c, 0x6f, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ActionEvent_proto_rawDescOnce sync.Once
	file_ActionEvent_proto_rawDescData = file_ActionEvent_proto_rawDesc
)

func file_ActionEvent_proto_rawDescGZIP() []byte {
	file_ActionEvent_proto_rawDescOnce.Do(func() {
		file_ActionEvent_proto_rawDescData = protoimpl.X.CompressGZIP(file_ActionEvent_proto_rawDescData)
	})
	return file_ActionEvent_proto_rawDescData
}

var file_ActionEvent_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ActionEvent_proto_goTypes = []interface{}{
	(*ActionEvent)(nil), // 0: ActionEvent
}
var file_ActionEvent_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ActionEvent_proto_init() }
func file_ActionEvent_proto_init() {
	if File_ActionEvent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ActionEvent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionEvent); i {
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
			RawDescriptor: file_ActionEvent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ActionEvent_proto_goTypes,
		DependencyIndexes: file_ActionEvent_proto_depIdxs,
		MessageInfos:      file_ActionEvent_proto_msgTypes,
	}.Build()
	File_ActionEvent_proto = out.File
	file_ActionEvent_proto_rawDesc = nil
	file_ActionEvent_proto_goTypes = nil
	file_ActionEvent_proto_depIdxs = nil
}
