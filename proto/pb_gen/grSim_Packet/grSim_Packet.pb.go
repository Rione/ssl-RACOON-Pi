// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: grSim_Packet.proto

package __

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

type GrSim_Packet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commands    *GrSim_Commands    `protobuf:"bytes,1,opt,name=commands" json:"commands,omitempty"`
	Replacement *GrSim_Replacement `protobuf:"bytes,2,opt,name=replacement" json:"replacement,omitempty"`
}

func (x *GrSim_Packet) Reset() {
	*x = GrSim_Packet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grSim_Packet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrSim_Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrSim_Packet) ProtoMessage() {}

func (x *GrSim_Packet) ProtoReflect() protoreflect.Message {
	mi := &file_grSim_Packet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrSim_Packet.ProtoReflect.Descriptor instead.
func (*GrSim_Packet) Descriptor() ([]byte, []int) {
	return file_grSim_Packet_proto_rawDescGZIP(), []int{0}
}

func (x *GrSim_Packet) GetCommands() *GrSim_Commands {
	if x != nil {
		return x.Commands
	}
	return nil
}

func (x *GrSim_Packet) GetReplacement() *GrSim_Replacement {
	if x != nil {
		return x.Replacement
	}
	return nil
}

var File_grSim_Packet_proto protoreflect.FileDescriptor

var file_grSim_Packet_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x72, 0x53, 0x69, 0x6d, 0x5f, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x72, 0x53, 0x69, 0x6d, 0x5f, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x72, 0x53, 0x69,
	0x6d, 0x5f, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0c, 0x67, 0x72, 0x53, 0x69, 0x6d, 0x5f, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x53, 0x69, 0x6d, 0x5f, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x12, 0x34, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x72, 0x53, 0x69, 0x6d, 0x5f, 0x52, 0x65,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f,
}

var (
	file_grSim_Packet_proto_rawDescOnce sync.Once
	file_grSim_Packet_proto_rawDescData = file_grSim_Packet_proto_rawDesc
)

func file_grSim_Packet_proto_rawDescGZIP() []byte {
	file_grSim_Packet_proto_rawDescOnce.Do(func() {
		file_grSim_Packet_proto_rawDescData = protoimpl.X.CompressGZIP(file_grSim_Packet_proto_rawDescData)
	})
	return file_grSim_Packet_proto_rawDescData
}

var file_grSim_Packet_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grSim_Packet_proto_goTypes = []interface{}{
	(*GrSim_Packet)(nil),      // 0: grSim_Packet
	(*GrSim_Commands)(nil),    // 1: grSim_Commands
	(*GrSim_Replacement)(nil), // 2: grSim_Replacement
}
var file_grSim_Packet_proto_depIdxs = []int32{
	1, // 0: grSim_Packet.commands:type_name -> grSim_Commands
	2, // 1: grSim_Packet.replacement:type_name -> grSim_Replacement
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grSim_Packet_proto_init() }
func file_grSim_Packet_proto_init() {
	if File_grSim_Packet_proto != nil {
		return
	}
	file_grSim_Commands_proto_init()
	file_grSim_Replacement_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_grSim_Packet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrSim_Packet); i {
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
			RawDescriptor: file_grSim_Packet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grSim_Packet_proto_goTypes,
		DependencyIndexes: file_grSim_Packet_proto_depIdxs,
		MessageInfos:      file_grSim_Packet_proto_msgTypes,
	}.Build()
	File_grSim_Packet_proto = out.File
	file_grSim_Packet_proto_rawDesc = nil
	file_grSim_Packet_proto_goTypes = nil
	file_grSim_Packet_proto_depIdxs = nil
}
