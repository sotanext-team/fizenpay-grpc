// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: fizenpay-auth/role.proto

package fizenpay_auth

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

type CheckPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email  string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Path   string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	Method string `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *CheckPermissionRequest) Reset() {
	*x = CheckPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fizenpay_auth_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionRequest) ProtoMessage() {}

func (x *CheckPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fizenpay_auth_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionRequest.ProtoReflect.Descriptor instead.
func (*CheckPermissionRequest) Descriptor() ([]byte, []int) {
	return file_fizenpay_auth_role_proto_rawDescGZIP(), []int{0}
}

func (x *CheckPermissionRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CheckPermissionRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CheckPermissionRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CheckPermissionRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type CheckPermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasPermission bool `protobuf:"varint,1,opt,name=hasPermission,proto3" json:"hasPermission,omitempty"`
}

func (x *CheckPermissionResponse) Reset() {
	*x = CheckPermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fizenpay_auth_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionResponse) ProtoMessage() {}

func (x *CheckPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fizenpay_auth_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionResponse.ProtoReflect.Descriptor instead.
func (*CheckPermissionResponse) Descriptor() ([]byte, []int) {
	return file_fizenpay_auth_role_proto_rawDescGZIP(), []int{1}
}

func (x *CheckPermissionResponse) GetHasPermission() bool {
	if x != nil {
		return x.HasPermission
	}
	return false
}

var File_fizenpay_auth_role_proto protoreflect.FileDescriptor

var file_fizenpay_auth_role_proto_rawDesc = []byte{
	0x0a, 0x18, 0x66, 0x69, 0x7a, 0x65, 0x6e, 0x70, 0x61, 0x79, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x66, 0x69, 0x7a, 0x65,
	0x6e, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x22, 0x72, 0x0a, 0x16, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x3f, 0x0a,
	0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x68, 0x61, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x71,
	0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a,
	0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x2e, 0x66, 0x69, 0x7a, 0x65, 0x6e, 0x70, 0x61, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x66, 0x69, 0x7a, 0x65, 0x6e, 0x70,
	0x61, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x74, 0x61, 0x6e, 0x65, 0x78, 0x74, 0x2d, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x66, 0x69,
	0x7a, 0x65, 0x6e, 0x70, 0x61, 0x79, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x69, 0x7a, 0x65,
	0x6e, 0x70, 0x61, 0x79, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x3b, 0x66, 0x69, 0x7a, 0x65, 0x6e, 0x70,
	0x61, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fizenpay_auth_role_proto_rawDescOnce sync.Once
	file_fizenpay_auth_role_proto_rawDescData = file_fizenpay_auth_role_proto_rawDesc
)

func file_fizenpay_auth_role_proto_rawDescGZIP() []byte {
	file_fizenpay_auth_role_proto_rawDescOnce.Do(func() {
		file_fizenpay_auth_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_fizenpay_auth_role_proto_rawDescData)
	})
	return file_fizenpay_auth_role_proto_rawDescData
}

var file_fizenpay_auth_role_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fizenpay_auth_role_proto_goTypes = []interface{}{
	(*CheckPermissionRequest)(nil),  // 0: fizenpay_auth.CheckPermissionRequest
	(*CheckPermissionResponse)(nil), // 1: fizenpay_auth.CheckPermissionResponse
}
var file_fizenpay_auth_role_proto_depIdxs = []int32{
	0, // 0: fizenpay_auth.RoleService.CheckPermission:input_type -> fizenpay_auth.CheckPermissionRequest
	1, // 1: fizenpay_auth.RoleService.CheckPermission:output_type -> fizenpay_auth.CheckPermissionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fizenpay_auth_role_proto_init() }
func file_fizenpay_auth_role_proto_init() {
	if File_fizenpay_auth_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fizenpay_auth_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionRequest); i {
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
		file_fizenpay_auth_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionResponse); i {
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
			RawDescriptor: file_fizenpay_auth_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fizenpay_auth_role_proto_goTypes,
		DependencyIndexes: file_fizenpay_auth_role_proto_depIdxs,
		MessageInfos:      file_fizenpay_auth_role_proto_msgTypes,
	}.Build()
	File_fizenpay_auth_role_proto = out.File
	file_fizenpay_auth_role_proto_rawDesc = nil
	file_fizenpay_auth_role_proto_goTypes = nil
	file_fizenpay_auth_role_proto_depIdxs = nil
}