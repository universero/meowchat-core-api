// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.23.4
// source: meowchat/core_api/sts.proto

package core_api

import (
	_ "github.com/xh-polaris/meowchat-core-api/biz/application/dto/http"
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

type ApplySignedUrlReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix *string `protobuf:"bytes,1,opt,name=prefix,proto3,oneof" form:"prefix" json:"prefix,omitempty"`
	Suffix *string `protobuf:"bytes,2,opt,name=suffix,proto3,oneof" form:"suffix" json:"suffix,omitempty"`
}

func (x *ApplySignedUrlReq) Reset() {
	*x = ApplySignedUrlReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_sts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplySignedUrlReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplySignedUrlReq) ProtoMessage() {}

func (x *ApplySignedUrlReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_sts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplySignedUrlReq.ProtoReflect.Descriptor instead.
func (*ApplySignedUrlReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_sts_proto_rawDescGZIP(), []int{0}
}

func (x *ApplySignedUrlReq) GetPrefix() string {
	if x != nil && x.Prefix != nil {
		return *x.Prefix
	}
	return ""
}

func (x *ApplySignedUrlReq) GetSuffix() string {
	if x != nil && x.Suffix != nil {
		return *x.Suffix
	}
	return ""
}

type ApplySignedUrlResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url          string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty" form:"url" query:"url"`
	SessionToken string `protobuf:"bytes,2,opt,name=sessionToken,proto3" json:"sessionToken,omitempty" form:"sessionToken" query:"sessionToken"`
}

func (x *ApplySignedUrlResp) Reset() {
	*x = ApplySignedUrlResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_sts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplySignedUrlResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplySignedUrlResp) ProtoMessage() {}

func (x *ApplySignedUrlResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_sts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplySignedUrlResp.ProtoReflect.Descriptor instead.
func (*ApplySignedUrlResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_sts_proto_rawDescGZIP(), []int{1}
}

func (x *ApplySignedUrlResp) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ApplySignedUrlResp) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

type ApplySignedUrlAsCommunityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunityId string `protobuf:"bytes,1,opt,name=communityId,proto3" json:"communityId,omitempty" form:"communityId" query:"communityId"`
	Prefix      string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty" form:"prefix" query:"prefix"`
	Suffix      string `protobuf:"bytes,3,opt,name=suffix,proto3" json:"suffix,omitempty" form:"suffix" query:"suffix"`
}

func (x *ApplySignedUrlAsCommunityReq) Reset() {
	*x = ApplySignedUrlAsCommunityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_sts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplySignedUrlAsCommunityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplySignedUrlAsCommunityReq) ProtoMessage() {}

func (x *ApplySignedUrlAsCommunityReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_sts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplySignedUrlAsCommunityReq.ProtoReflect.Descriptor instead.
func (*ApplySignedUrlAsCommunityReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_sts_proto_rawDescGZIP(), []int{2}
}

func (x *ApplySignedUrlAsCommunityReq) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

func (x *ApplySignedUrlAsCommunityReq) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *ApplySignedUrlAsCommunityReq) GetSuffix() string {
	if x != nil {
		return x.Suffix
	}
	return ""
}

type ApplySignedUrlAsCommunityResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url          string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty" form:"url" query:"url"`
	SessionToken string `protobuf:"bytes,2,opt,name=sessionToken,proto3" json:"sessionToken,omitempty" form:"sessionToken" query:"sessionToken"`
}

func (x *ApplySignedUrlAsCommunityResp) Reset() {
	*x = ApplySignedUrlAsCommunityResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_sts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplySignedUrlAsCommunityResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplySignedUrlAsCommunityResp) ProtoMessage() {}

func (x *ApplySignedUrlAsCommunityResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_sts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplySignedUrlAsCommunityResp.ProtoReflect.Descriptor instead.
func (*ApplySignedUrlAsCommunityResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_sts_proto_rawDescGZIP(), []int{3}
}

func (x *ApplySignedUrlAsCommunityResp) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ApplySignedUrlAsCommunityResp) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

var File_meowchat_core_api_sts_proto protoreflect.FileDescriptor

var file_meowchat_core_api_sts_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6d,
	0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x1a, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7b, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xca, 0xbb, 0x18, 0x06, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x48, 0x00, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x12,
	0x27, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xca, 0xbb, 0x18, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x48, 0x01, 0x52, 0x06, 0x73,
	0x75, 0x66, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x22, 0x4a,
	0x0a, 0x12, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x70, 0x0a, 0x1c, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x41, 0x73, 0x43, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x22, 0x55, 0x0a, 0x1d,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x41, 0x73,
	0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x22, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x62,
	0x69, 0x7a, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64,
	0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meowchat_core_api_sts_proto_rawDescOnce sync.Once
	file_meowchat_core_api_sts_proto_rawDescData = file_meowchat_core_api_sts_proto_rawDesc
)

func file_meowchat_core_api_sts_proto_rawDescGZIP() []byte {
	file_meowchat_core_api_sts_proto_rawDescOnce.Do(func() {
		file_meowchat_core_api_sts_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowchat_core_api_sts_proto_rawDescData)
	})
	return file_meowchat_core_api_sts_proto_rawDescData
}

var file_meowchat_core_api_sts_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_meowchat_core_api_sts_proto_goTypes = []interface{}{
	(*ApplySignedUrlReq)(nil),             // 0: meowchat.core_api.ApplySignedUrlReq
	(*ApplySignedUrlResp)(nil),            // 1: meowchat.core_api.ApplySignedUrlResp
	(*ApplySignedUrlAsCommunityReq)(nil),  // 2: meowchat.core_api.ApplySignedUrlAsCommunityReq
	(*ApplySignedUrlAsCommunityResp)(nil), // 3: meowchat.core_api.ApplySignedUrlAsCommunityResp
}
var file_meowchat_core_api_sts_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func file_meowchat_core_api_sts_proto_init() {
	if File_meowchat_core_api_sts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meowchat_core_api_sts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplySignedUrlReq); i {
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
		file_meowchat_core_api_sts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplySignedUrlResp); i {
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
		file_meowchat_core_api_sts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplySignedUrlAsCommunityReq); i {
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
		file_meowchat_core_api_sts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplySignedUrlAsCommunityResp); i {
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
	file_meowchat_core_api_sts_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meowchat_core_api_sts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meowchat_core_api_sts_proto_goTypes,
		DependencyIndexes: file_meowchat_core_api_sts_proto_depIdxs,
		MessageInfos:      file_meowchat_core_api_sts_proto_msgTypes,
	}.Build()
	File_meowchat_core_api_sts_proto = out.File
	file_meowchat_core_api_sts_proto_rawDesc = nil
	file_meowchat_core_api_sts_proto_goTypes = nil
	file_meowchat_core_api_sts_proto_depIdxs = nil
}
