//
// Copyright © 2022 Durudex
//
// This file is part of Durudex: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// Durudex is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with Durudex. If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: internal/delivery/grpc/pb/user.proto

package pb

import (
	timestamp "github.com/durudex/dugopb/types/timestamp"
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

type GetUserByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserByIDRequest) Reset() {
	*x = GetUserByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_pb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIDRequest) ProtoMessage() {}

func (x *GetUserByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_pb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIDRequest.ProtoReflect.Descriptor instead.
func (*GetUserByIDRequest) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserByIDRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type GetUserByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string               `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastVisit *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_visit,json=lastVisit,proto3" json:"last_visit,omitempty"`
	Verified  bool                 `protobuf:"varint,4,opt,name=verified,proto3" json:"verified,omitempty"`
	AvatarUrl *string              `protobuf:"bytes,5,opt,name=avatar_url,json=avatarUrl,proto3,oneof" json:"avatar_url,omitempty"`
}

func (x *GetUserByIDResponse) Reset() {
	*x = GetUserByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_pb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIDResponse) ProtoMessage() {}

func (x *GetUserByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_pb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIDResponse.ProtoReflect.Descriptor instead.
func (*GetUserByIDResponse) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserByIDResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetUserByIDResponse) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GetUserByIDResponse) GetLastVisit() *timestamp.Timestamp {
	if x != nil {
		return x.LastVisit
	}
	return nil
}

func (x *GetUserByIDResponse) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *GetUserByIDResponse) GetAvatarUrl() string {
	if x != nil && x.AvatarUrl != nil {
		return *x.AvatarUrl
	}
	return ""
}

type ForgotPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ForgotPasswordRequest) Reset() {
	*x = ForgotPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_pb_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgotPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgotPasswordRequest) ProtoMessage() {}

func (x *ForgotPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_pb_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgotPasswordRequest.ProtoReflect.Descriptor instead.
func (*ForgotPasswordRequest) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_pb_user_proto_rawDescGZIP(), []int{2}
}

func (x *ForgotPasswordRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ForgotPasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ForgotPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ForgotPasswordResponse) Reset() {
	*x = ForgotPasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_pb_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgotPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgotPasswordResponse) ProtoMessage() {}

func (x *ForgotPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_pb_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgotPasswordResponse.ProtoReflect.Descriptor instead.
func (*ForgotPasswordResponse) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_pb_user_proto_rawDescGZIP(), []int{3}
}

var File_internal_delivery_grpc_pb_user_proto protoreflect.FileDescriptor

var file_internal_delivery_grpc_pb_user_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x1a,
	0x1d, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xf2, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64,
	0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x37, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x56, 0x69, 0x73, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x49, 0x0a, 0x15, 0x46, 0x6f, 0x72,
	0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xaa,
	0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1b, 0x2e,
	0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x75, 0x72,
	0x75, 0x64, 0x65, 0x78, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x67,
	0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x2e, 0x64, 0x75, 0x72,
	0x75, 0x64, 0x65, 0x78, 0x2e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x75, 0x72,
	0x75, 0x64, 0x65, 0x78, 0x2e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x41, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65,
	0x78, 0x2f, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_delivery_grpc_pb_user_proto_rawDescOnce sync.Once
	file_internal_delivery_grpc_pb_user_proto_rawDescData = file_internal_delivery_grpc_pb_user_proto_rawDesc
)

func file_internal_delivery_grpc_pb_user_proto_rawDescGZIP() []byte {
	file_internal_delivery_grpc_pb_user_proto_rawDescOnce.Do(func() {
		file_internal_delivery_grpc_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_delivery_grpc_pb_user_proto_rawDescData)
	})
	return file_internal_delivery_grpc_pb_user_proto_rawDescData
}

var file_internal_delivery_grpc_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_delivery_grpc_pb_user_proto_goTypes = []interface{}{
	(*GetUserByIDRequest)(nil),     // 0: durudex.GetUserByIDRequest
	(*GetUserByIDResponse)(nil),    // 1: durudex.GetUserByIDResponse
	(*ForgotPasswordRequest)(nil),  // 2: durudex.ForgotPasswordRequest
	(*ForgotPasswordResponse)(nil), // 3: durudex.ForgotPasswordResponse
	(*timestamp.Timestamp)(nil),    // 4: durudex.types.Timestamp
}
var file_internal_delivery_grpc_pb_user_proto_depIdxs = []int32{
	4, // 0: durudex.GetUserByIDResponse.created_at:type_name -> durudex.types.Timestamp
	4, // 1: durudex.GetUserByIDResponse.last_visit:type_name -> durudex.types.Timestamp
	0, // 2: durudex.UserService.GetUserByID:input_type -> durudex.GetUserByIDRequest
	2, // 3: durudex.UserService.ForgotPassword:input_type -> durudex.ForgotPasswordRequest
	1, // 4: durudex.UserService.GetUserByID:output_type -> durudex.GetUserByIDResponse
	3, // 5: durudex.UserService.ForgotPassword:output_type -> durudex.ForgotPasswordResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_delivery_grpc_pb_user_proto_init() }
func file_internal_delivery_grpc_pb_user_proto_init() {
	if File_internal_delivery_grpc_pb_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_delivery_grpc_pb_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIDRequest); i {
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
		file_internal_delivery_grpc_pb_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByIDResponse); i {
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
		file_internal_delivery_grpc_pb_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgotPasswordRequest); i {
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
		file_internal_delivery_grpc_pb_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgotPasswordResponse); i {
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
	file_internal_delivery_grpc_pb_user_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_delivery_grpc_pb_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_delivery_grpc_pb_user_proto_goTypes,
		DependencyIndexes: file_internal_delivery_grpc_pb_user_proto_depIdxs,
		MessageInfos:      file_internal_delivery_grpc_pb_user_proto_msgTypes,
	}.Build()
	File_internal_delivery_grpc_pb_user_proto = out.File
	file_internal_delivery_grpc_pb_user_proto_rawDesc = nil
	file_internal_delivery_grpc_pb_user_proto_goTypes = nil
	file_internal_delivery_grpc_pb_user_proto_depIdxs = nil
}
