// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: user_center_user_extra.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserExtra struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    uint32     `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Keyword   string     `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Type      string     `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Value     *anypb.Any `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	CreatedAt uint32     `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *UserExtra) Reset() {
	*x = UserExtra{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_user_extra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserExtra) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserExtra) ProtoMessage() {}

func (x *UserExtra) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_user_extra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserExtra.ProtoReflect.Descriptor instead.
func (*UserExtra) Descriptor() ([]byte, []int) {
	return file_user_center_user_extra_proto_rawDescGZIP(), []int{0}
}

func (x *UserExtra) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserExtra) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserExtra) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *UserExtra) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *UserExtra) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *UserExtra) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type AllUserExtraRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AllUserExtraRequest) Reset() {
	*x = AllUserExtraRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_user_extra_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllUserExtraRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllUserExtraRequest) ProtoMessage() {}

func (x *AllUserExtraRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_user_extra_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllUserExtraRequest.ProtoReflect.Descriptor instead.
func (*AllUserExtraRequest) Descriptor() ([]byte, []int) {
	return file_user_center_user_extra_proto_rawDescGZIP(), []int{1}
}

func (x *AllUserExtraRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AllUserExtraReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*UserExtra `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *AllUserExtraReply) Reset() {
	*x = AllUserExtraReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_user_extra_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllUserExtraReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllUserExtraReply) ProtoMessage() {}

func (x *AllUserExtraReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_user_extra_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllUserExtraReply.ProtoReflect.Descriptor instead.
func (*AllUserExtraReply) Descriptor() ([]byte, []int) {
	return file_user_center_user_extra_proto_rawDescGZIP(), []int{2}
}

func (x *AllUserExtraReply) GetList() []*UserExtra {
	if x != nil {
		return x.List
	}
	return nil
}

type AddUserExtraRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint32     `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Keyword string     `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Type    string     `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Value   *anypb.Any `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AddUserExtraRequest) Reset() {
	*x = AddUserExtraRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_user_extra_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserExtraRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserExtraRequest) ProtoMessage() {}

func (x *AddUserExtraRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_user_extra_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserExtraRequest.ProtoReflect.Descriptor instead.
func (*AddUserExtraRequest) Descriptor() ([]byte, []int) {
	return file_user_center_user_extra_proto_rawDescGZIP(), []int{3}
}

func (x *AddUserExtraRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddUserExtraRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *AddUserExtraRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AddUserExtraRequest) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type DeleteUserExtraRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Keyword string `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *DeleteUserExtraRequest) Reset() {
	*x = DeleteUserExtraRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_user_extra_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserExtraRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserExtraRequest) ProtoMessage() {}

func (x *DeleteUserExtraRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_user_extra_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserExtraRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserExtraRequest) Descriptor() ([]byte, []int) {
	return file_user_center_user_extra_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteUserExtraRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeleteUserExtraRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

var File_user_center_user_extra_proto protoreflect.FileDescriptor

var file_user_center_user_extra_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xad, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x37, 0x0a, 0x13, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x74, 0x72, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x11, 0x41, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x78, 0x74, 0x72, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xad, 0x01, 0x0a, 0x13, 0x41, 0x64,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x74, 0x72, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x00, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x00, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xa2, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5d, 0x0a, 0x16, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x74, 0x72, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_center_user_extra_proto_rawDescOnce sync.Once
	file_user_center_user_extra_proto_rawDescData = file_user_center_user_extra_proto_rawDesc
)

func file_user_center_user_extra_proto_rawDescGZIP() []byte {
	file_user_center_user_extra_proto_rawDescOnce.Do(func() {
		file_user_center_user_extra_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_center_user_extra_proto_rawDescData)
	})
	return file_user_center_user_extra_proto_rawDescData
}

var file_user_center_user_extra_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_center_user_extra_proto_goTypes = []interface{}{
	(*UserExtra)(nil),              // 0: user_center.UserExtra
	(*AllUserExtraRequest)(nil),    // 1: user_center.AllUserExtraRequest
	(*AllUserExtraReply)(nil),      // 2: user_center.AllUserExtraReply
	(*AddUserExtraRequest)(nil),    // 3: user_center.AddUserExtraRequest
	(*DeleteUserExtraRequest)(nil), // 4: user_center.DeleteUserExtraRequest
	(*anypb.Any)(nil),              // 5: google.protobuf.Any
}
var file_user_center_user_extra_proto_depIdxs = []int32{
	5, // 0: user_center.UserExtra.value:type_name -> google.protobuf.Any
	0, // 1: user_center.AllUserExtraReply.list:type_name -> user_center.UserExtra
	5, // 2: user_center.AddUserExtraRequest.value:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_user_center_user_extra_proto_init() }
func file_user_center_user_extra_proto_init() {
	if File_user_center_user_extra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_center_user_extra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserExtra); i {
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
		file_user_center_user_extra_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllUserExtraRequest); i {
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
		file_user_center_user_extra_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllUserExtraReply); i {
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
		file_user_center_user_extra_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserExtraRequest); i {
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
		file_user_center_user_extra_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserExtraRequest); i {
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
			RawDescriptor: file_user_center_user_extra_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_center_user_extra_proto_goTypes,
		DependencyIndexes: file_user_center_user_extra_proto_depIdxs,
		MessageInfos:      file_user_center_user_extra_proto_msgTypes,
	}.Build()
	File_user_center_user_extra_proto = out.File
	file_user_center_user_extra_proto_rawDesc = nil
	file_user_center_user_extra_proto_goTypes = nil
	file_user_center_user_extra_proto_depIdxs = nil
}