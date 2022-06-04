// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pb/run.proto

package pb

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

type ChatNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ChatNote) Reset() {
	*x = ChatNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_run_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatNote) ProtoMessage() {}

func (x *ChatNote) ProtoReflect() protoreflect.Message {
	mi := &file_pb_run_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatNote.ProtoReflect.Descriptor instead.
func (*ChatNote) Descriptor() ([]byte, []int) {
	return file_pb_run_proto_rawDescGZIP(), []int{0}
}

func (x *ChatNote) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_pb_run_proto protoreflect.FileDescriptor

var file_pb_run_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x62, 0x2f, 0x72, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x1c, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x32, 0x2d, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x26, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x1a, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x79,
	0x6c, 0x61, 0x6e, 0x47, 0x72, 0x61, 0x68, 0x61, 0x6d, 0x2f, 0x72, 0x75, 0x6e, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_run_proto_rawDescOnce sync.Once
	file_pb_run_proto_rawDescData = file_pb_run_proto_rawDesc
)

func file_pb_run_proto_rawDescGZIP() []byte {
	file_pb_run_proto_rawDescOnce.Do(func() {
		file_pb_run_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_run_proto_rawDescData)
	})
	return file_pb_run_proto_rawDescData
}

var file_pb_run_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_run_proto_goTypes = []interface{}{
	(*ChatNote)(nil), // 0: pb.ChatNote
}
var file_pb_run_proto_depIdxs = []int32{
	0, // 0: pb.Run.Chat:input_type -> pb.ChatNote
	0, // 1: pb.Run.Chat:output_type -> pb.ChatNote
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_run_proto_init() }
func file_pb_run_proto_init() {
	if File_pb_run_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_run_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatNote); i {
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
			RawDescriptor: file_pb_run_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_run_proto_goTypes,
		DependencyIndexes: file_pb_run_proto_depIdxs,
		MessageInfos:      file_pb_run_proto_msgTypes,
	}.Build()
	File_pb_run_proto = out.File
	file_pb_run_proto_rawDesc = nil
	file_pb_run_proto_goTypes = nil
	file_pb_run_proto_depIdxs = nil
}
