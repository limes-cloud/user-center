// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: user_center_user_service.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_user_center_user_service_proto protoreflect.FileDescriptor

var file_user_center_user_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfa, 0x1a, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x4f,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x59, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x73, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a,
	0x01, 0x2a, 0x1a, 0x1b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x5b, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12,
	0x1b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x5a, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x65, 0x0a, 0x0a, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a,
	0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x64, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x25,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x1a, 0x1a, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x61, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x2a, 0x1a, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x6e, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x27, 0x3a, 0x01, 0x2a, 0x22, 0x22, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x6b, 0x0a, 0x0a, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a,
	0x01, 0x2a, 0x22, 0x21, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x6e, 0x0a, 0x0b, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x66, 0x66, 0x6c,
	0x69, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01,
	0x2a, 0x22, 0x22, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6f, 0x66,
	0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x68, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x70, 0x70, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22,
	0x1e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x12,
	0x6b, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70,
	0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x2a, 0x1e, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x12, 0x66, 0x0a, 0x0a,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a,
	0x22, 0x22, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x12, 0x7f, 0x0a, 0x13, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x42, 0x69, 0x6e,
	0x64, 0x42, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x42, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x35,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x01, 0x2a, 0x22, 0x2a, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x69, 0x6e, 0x64, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x7c, 0x0a, 0x12, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x42, 0x69,
	0x6e, 0x64, 0x42, 0x79, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x1f, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x42, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x43, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x34, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a, 0x01, 0x2a, 0x22, 0x29, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x69, 0x6e, 0x64, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x12, 0x7c, 0x0a, 0x15, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x42, 0x69, 0x6e, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31,
	0x12, 0x2f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6e, 0x64, 0x2f, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2d, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x12, 0x8b, 0x01, 0x0a, 0x15, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x42, 0x69, 0x6e, 0x64, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x22, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x42, 0x69, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x3a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x3a, 0x01, 0x2a, 0x22, 0x2f,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6e, 0x64, 0x2f, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12,
	0x6f, 0x0a, 0x0d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x30,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x7f, 0x0a, 0x14, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x3b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x12, 0x33, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2d, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x12, 0x7b, 0x0a, 0x10, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2d, 0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x95,
	0x01, 0x0a, 0x15, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x36,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x85, 0x01, 0x0a, 0x17, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x3e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x38, 0x12, 0x36, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2d, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x6c,
	0x0a, 0x0c, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x19,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x7e, 0x0a, 0x11,
	0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x01, 0x2a,
	0x22, 0x2a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x63,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x8c, 0x01, 0x0a,
	0x14, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x3d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x37, 0x12, 0x35, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x2d, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x78, 0x0a, 0x0f, 0x43,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x67, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a, 0x01, 0x2a, 0x22, 0x1f, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x12, 0x63,
	0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23,
	0x22, 0x21, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_user_center_user_service_proto_goTypes = []interface{}{
	(*GetSimpleUserRequest)(nil),         // 0: user.GetSimpleUserRequest
	(*GetBaseUserRequest)(nil),           // 1: user.GetBaseUserRequest
	(*GetUserRequest)(nil),               // 2: user.GetUserRequest
	(*emptypb.Empty)(nil),                // 3: google.protobuf.Empty
	(*UpdateCurrentUserRequest)(nil),     // 4: user.UpdateCurrentUserRequest
	(*PageUserRequest)(nil),              // 5: user.PageUserRequest
	(*AddUserRequest)(nil),               // 6: user.AddUserRequest
	(*ImportUserRequest)(nil),            // 7: user.ImportUserRequest
	(*UpdateUserRequest)(nil),            // 8: user.UpdateUserRequest
	(*DeleteUserRequest)(nil),            // 9: user.DeleteUserRequest
	(*DisableUserRequest)(nil),           // 10: user.DisableUserRequest
	(*EnableUserRequest)(nil),            // 11: user.EnableUserRequest
	(*OfflineUserRequest)(nil),           // 12: user.OfflineUserRequest
	(*AddUserAppRequest)(nil),            // 13: user.AddUserAppRequest
	(*DeleteUserAppRequest)(nil),         // 14: user.DeleteUserAppRequest
	(*OAuthLoginRequest)(nil),            // 15: user.OAuthLoginRequest
	(*OAuthBindByPasswordRequest)(nil),   // 16: user.OAuthBindByPasswordRequest
	(*OAuthBindByCaptchaRequest)(nil),    // 17: user.OAuthBindByCaptchaRequest
	(*OAuthBindEmailCaptchaRequest)(nil), // 18: user.OAuthBindEmailCaptchaRequest
	(*PasswordLoginRequest)(nil),         // 19: user.PasswordLoginRequest
	(*PasswordRegisterRequest)(nil),      // 20: user.PasswordRegisterRequest
	(*PasswordRegisterCheckRequest)(nil), // 21: user.PasswordRegisterCheckRequest
	(*CaptchaLoginRequest)(nil),          // 22: user.CaptchaLoginRequest
	(*CaptchaLoginEmailRequest)(nil),     // 23: user.CaptchaLoginEmailRequest
	(*CaptchaRegisterEmailRequest)(nil),  // 24: user.CaptchaRegisterEmailRequest
	(*CaptchaRegisterRequest)(nil),       // 25: user.CaptchaRegisterRequest
	(*SimpleUser)(nil),                   // 26: user.SimpleUser
	(*BaseUser)(nil),                     // 27: user.BaseUser
	(*User)(nil),                         // 28: user.User
	(*PageUserReply)(nil),                // 29: user.PageUserReply
	(*AddUserReply)(nil),                 // 30: user.AddUserReply
	(*LoginReply)(nil),                   // 31: user.LoginReply
	(*BindReply)(nil),                    // 32: user.BindReply
	(*CaptchaReply)(nil),                 // 33: user.CaptchaReply
	(*RegisterReply)(nil),                // 34: user.RegisterReply
	(*PasswordRegisterCheckReply)(nil),   // 35: user.PasswordRegisterCheckReply
	(*ParseTokenReply)(nil),              // 36: user.ParseTokenReply
}
var file_user_center_user_service_proto_depIdxs = []int32{
	0,  // 0: user.Service.GetSimpleUser:input_type -> user.GetSimpleUserRequest
	1,  // 1: user.Service.GetBaseUser:input_type -> user.GetBaseUserRequest
	2,  // 2: user.Service.GetUser:input_type -> user.GetUserRequest
	3,  // 3: user.Service.GetCurrentUser:input_type -> google.protobuf.Empty
	4,  // 4: user.Service.UpdateCurrentUser:input_type -> user.UpdateCurrentUserRequest
	5,  // 5: user.Service.PageUser:input_type -> user.PageUserRequest
	6,  // 6: user.Service.AddUser:input_type -> user.AddUserRequest
	7,  // 7: user.Service.ImportUser:input_type -> user.ImportUserRequest
	8,  // 8: user.Service.UpdateUser:input_type -> user.UpdateUserRequest
	9,  // 9: user.Service.DeleteUser:input_type -> user.DeleteUserRequest
	10, // 10: user.Service.DisableUser:input_type -> user.DisableUserRequest
	11, // 11: user.Service.EnableUser:input_type -> user.EnableUserRequest
	12, // 12: user.Service.OfflineUser:input_type -> user.OfflineUserRequest
	13, // 13: user.Service.AddUserApp:input_type -> user.AddUserAppRequest
	14, // 14: user.Service.DeleteUserApp:input_type -> user.DeleteUserAppRequest
	15, // 15: user.Service.OAuthLogin:input_type -> user.OAuthLoginRequest
	16, // 16: user.Service.OAuthBindByPassword:input_type -> user.OAuthBindByPasswordRequest
	17, // 17: user.Service.OAuthBindByCaptcha:input_type -> user.OAuthBindByCaptchaRequest
	3,  // 18: user.Service.OAuthBindImageCaptcha:input_type -> google.protobuf.Empty
	18, // 19: user.Service.OAuthBindEmailCaptcha:input_type -> user.OAuthBindEmailCaptchaRequest
	19, // 20: user.Service.PasswordLogin:input_type -> user.PasswordLoginRequest
	3,  // 21: user.Service.PasswordLoginCaptcha:input_type -> google.protobuf.Empty
	20, // 22: user.Service.PasswordRegister:input_type -> user.PasswordRegisterRequest
	21, // 23: user.Service.PasswordRegisterCheck:input_type -> user.PasswordRegisterCheckRequest
	3,  // 24: user.Service.PasswordRegisterCaptcha:input_type -> google.protobuf.Empty
	22, // 25: user.Service.CaptchaLogin:input_type -> user.CaptchaLoginRequest
	23, // 26: user.Service.CaptchaLoginEmail:input_type -> user.CaptchaLoginEmailRequest
	24, // 27: user.Service.CaptchaRegisterEmail:input_type -> user.CaptchaRegisterEmailRequest
	25, // 28: user.Service.CaptchaRegister:input_type -> user.CaptchaRegisterRequest
	3,  // 29: user.Service.ParseToken:input_type -> google.protobuf.Empty
	3,  // 30: user.Service.RefreshToken:input_type -> google.protobuf.Empty
	26, // 31: user.Service.GetSimpleUser:output_type -> user.SimpleUser
	27, // 32: user.Service.GetBaseUser:output_type -> user.BaseUser
	28, // 33: user.Service.GetUser:output_type -> user.User
	28, // 34: user.Service.GetCurrentUser:output_type -> user.User
	3,  // 35: user.Service.UpdateCurrentUser:output_type -> google.protobuf.Empty
	29, // 36: user.Service.PageUser:output_type -> user.PageUserReply
	30, // 37: user.Service.AddUser:output_type -> user.AddUserReply
	3,  // 38: user.Service.ImportUser:output_type -> google.protobuf.Empty
	3,  // 39: user.Service.UpdateUser:output_type -> google.protobuf.Empty
	3,  // 40: user.Service.DeleteUser:output_type -> google.protobuf.Empty
	3,  // 41: user.Service.DisableUser:output_type -> google.protobuf.Empty
	3,  // 42: user.Service.EnableUser:output_type -> google.protobuf.Empty
	3,  // 43: user.Service.OfflineUser:output_type -> google.protobuf.Empty
	3,  // 44: user.Service.AddUserApp:output_type -> google.protobuf.Empty
	3,  // 45: user.Service.DeleteUserApp:output_type -> google.protobuf.Empty
	31, // 46: user.Service.OAuthLogin:output_type -> user.LoginReply
	32, // 47: user.Service.OAuthBindByPassword:output_type -> user.BindReply
	32, // 48: user.Service.OAuthBindByCaptcha:output_type -> user.BindReply
	33, // 49: user.Service.OAuthBindImageCaptcha:output_type -> user.CaptchaReply
	33, // 50: user.Service.OAuthBindEmailCaptcha:output_type -> user.CaptchaReply
	31, // 51: user.Service.PasswordLogin:output_type -> user.LoginReply
	33, // 52: user.Service.PasswordLoginCaptcha:output_type -> user.CaptchaReply
	34, // 53: user.Service.PasswordRegister:output_type -> user.RegisterReply
	35, // 54: user.Service.PasswordRegisterCheck:output_type -> user.PasswordRegisterCheckReply
	33, // 55: user.Service.PasswordRegisterCaptcha:output_type -> user.CaptchaReply
	31, // 56: user.Service.CaptchaLogin:output_type -> user.LoginReply
	33, // 57: user.Service.CaptchaLoginEmail:output_type -> user.CaptchaReply
	33, // 58: user.Service.CaptchaRegisterEmail:output_type -> user.CaptchaReply
	34, // 59: user.Service.CaptchaRegister:output_type -> user.RegisterReply
	36, // 60: user.Service.ParseToken:output_type -> user.ParseTokenReply
	31, // 61: user.Service.RefreshToken:output_type -> user.LoginReply
	31, // [31:62] is the sub-list for method output_type
	0,  // [0:31] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_user_center_user_service_proto_init() }
func file_user_center_user_service_proto_init() {
	if File_user_center_user_service_proto != nil {
		return
	}
	file_user_center_user_proto_init()
	file_user_center_auth_proto_init()
	file_user_center_user_app_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_center_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_center_user_service_proto_goTypes,
		DependencyIndexes: file_user_center_user_service_proto_depIdxs,
	}.Build()
	File_user_center_user_service_proto = out.File
	file_user_center_user_service_proto_rawDesc = nil
	file_user_center_user_service_proto_goTypes = nil
	file_user_center_user_service_proto_depIdxs = nil
}
