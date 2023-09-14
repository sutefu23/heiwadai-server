// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: v1/user/Store.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	shared "server/api/v1/shared"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SoreIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *SoreIDRequest) Reset() {
	*x = SoreIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_user_Store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoreIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoreIDRequest) ProtoMessage() {}

func (x *SoreIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_user_Store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoreIDRequest.ProtoReflect.Descriptor instead.
func (*SoreIDRequest) Descriptor() ([]byte, []int) {
	return file_v1_user_Store_proto_rawDescGZIP(), []int{0}
}

func (x *SoreIDRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

var File_v1_user_Store_proto protoreflect.FileDescriptor

var file_v1_user_Store_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0d, 0x53, 0x6f, 0x72, 0x65, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x32, 0xa3, 0x02, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x6f, 0x72, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x73, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x79,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x53, 0x74,
	0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4d,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x53, 0x6f, 0x72, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x53, 0x74,
	0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x00, 0x42, 0x7e, 0x0a,
	0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x42, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x12,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0xa2, 0x02, 0x03, 0x53, 0x55, 0x58, 0xaa, 0x02, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c,
	0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x17, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x55, 0x73,
	0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_user_Store_proto_rawDescOnce sync.Once
	file_v1_user_Store_proto_rawDescData = file_v1_user_Store_proto_rawDesc
)

func file_v1_user_Store_proto_rawDescGZIP() []byte {
	file_v1_user_Store_proto_rawDescOnce.Do(func() {
		file_v1_user_Store_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_user_Store_proto_rawDescData)
	})
	return file_v1_user_Store_proto_rawDescData
}

var file_v1_user_Store_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v1_user_Store_proto_goTypes = []interface{}{
	(*SoreIDRequest)(nil),         // 0: server.user.SoreIDRequest
	(*emptypb.Empty)(nil),         // 1: google.protobuf.Empty
	(*shared.Store)(nil),          // 2: server.shared.Store
	(*shared.Stores)(nil),         // 3: server.shared.Stores
	(*shared.StayableStores)(nil), // 4: server.shared.StayableStores
	(*shared.StayableStore)(nil),  // 5: server.shared.StayableStore
}
var file_v1_user_Store_proto_depIdxs = []int32{
	0, // 0: server.user.StoreController.GetByID:input_type -> server.user.SoreIDRequest
	1, // 1: server.user.StoreController.GetAll:input_type -> google.protobuf.Empty
	1, // 2: server.user.StoreController.GetStayables:input_type -> google.protobuf.Empty
	0, // 3: server.user.StoreController.GetStayableByID:input_type -> server.user.SoreIDRequest
	2, // 4: server.user.StoreController.GetByID:output_type -> server.shared.Store
	3, // 5: server.user.StoreController.GetAll:output_type -> server.shared.Stores
	4, // 6: server.user.StoreController.GetStayables:output_type -> server.shared.StayableStores
	5, // 7: server.user.StoreController.GetStayableByID:output_type -> server.shared.StayableStore
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_user_Store_proto_init() }
func file_v1_user_Store_proto_init() {
	if File_v1_user_Store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_user_Store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoreIDRequest); i {
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
			RawDescriptor: file_v1_user_Store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_user_Store_proto_goTypes,
		DependencyIndexes: file_v1_user_Store_proto_depIdxs,
		MessageInfos:      file_v1_user_Store_proto_msgTypes,
	}.Build()
	File_v1_user_Store_proto = out.File
	file_v1_user_Store_proto_rawDesc = nil
	file_v1_user_Store_proto_goTypes = nil
	file_v1_user_Store_proto_depIdxs = nil
}
