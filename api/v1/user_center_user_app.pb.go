// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: user_center_user_app.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type AddUserAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AppId  uint32 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *AddUserAppRequest) Reset() {
	*x = AddUserAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_user_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserAppRequest) ProtoMessage() {}

func (x *AddUserAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_user_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserAppRequest.ProtoReflect.Descriptor instead.
func (*AddUserAppRequest) Descriptor() ([]byte, []int) {
	return file_user_center_user_app_proto_rawDescGZIP(), []int{0}
}

func (x *AddUserAppRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddUserAppRequest) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

type DeleteUserAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AppId  uint32 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *DeleteUserAppRequest) Reset() {
	*x = DeleteUserAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_user_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserAppRequest) ProtoMessage() {}

func (x *DeleteUserAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_user_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserAppRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserAppRequest) Descriptor() ([]byte, []int) {
	return file_user_center_user_app_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteUserAppRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeleteUserAppRequest) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

var File_user_center_user_app_proto protoreflect.FileDescriptor

var file_user_center_user_app_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x55, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20,
	0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x06, 0x61, 0x70, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02,
	0x20, 0x00, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_center_user_app_proto_rawDescOnce sync.Once
	file_user_center_user_app_proto_rawDescData = file_user_center_user_app_proto_rawDesc
)

func file_user_center_user_app_proto_rawDescGZIP() []byte {
	file_user_center_user_app_proto_rawDescOnce.Do(func() {
		file_user_center_user_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_center_user_app_proto_rawDescData)
	})
	return file_user_center_user_app_proto_rawDescData
}

var file_user_center_user_app_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_center_user_app_proto_goTypes = []interface{}{
	(*AddUserAppRequest)(nil),    // 0: user_center.AddUserAppRequest
	(*DeleteUserAppRequest)(nil), // 1: user_center.DeleteUserAppRequest
}
var file_user_center_user_app_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_center_user_app_proto_init() }
func file_user_center_user_app_proto_init() {
	if File_user_center_user_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_center_user_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserAppRequest); i {
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
		file_user_center_user_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserAppRequest); i {
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
			RawDescriptor: file_user_center_user_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_center_user_app_proto_goTypes,
		DependencyIndexes: file_user_center_user_app_proto_depIdxs,
		MessageInfos:      file_user_center_user_app_proto_msgTypes,
	}.Build()
	File_user_center_user_app_proto = out.File
	file_user_center_user_app_proto_rawDesc = nil
	file_user_center_user_app_proto_goTypes = nil
	file_user_center_user_app_proto_depIdxs = nil
}