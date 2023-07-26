// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.23.4
// source: meowchat/basic/basic.proto

package basic

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

type DevicePlatform int32

const (
	DevicePlatform_Android        DevicePlatform = 0
	DevicePlatform_IOS            DevicePlatform = 1
	DevicePlatform_Windows        DevicePlatform = 2
	DevicePlatform_Mac            DevicePlatform = 3
	DevicePlatform_WechatDevtools DevicePlatform = 4
)

// Enum value maps for DevicePlatform.
var (
	DevicePlatform_name = map[int32]string{
		0: "Android",
		1: "IOS",
		2: "Windows",
		3: "Mac",
		4: "WechatDevtools",
	}
	DevicePlatform_value = map[string]int32{
		"Android":        0,
		"IOS":            1,
		"Windows":        2,
		"Mac":            3,
		"WechatDevtools": 4,
	}
)

func (x DevicePlatform) Enum() *DevicePlatform {
	p := new(DevicePlatform)
	*p = x
	return p
}

func (x DevicePlatform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DevicePlatform) Descriptor() protoreflect.EnumDescriptor {
	return file_meowchat_basic_basic_proto_enumTypes[0].Descriptor()
}

func (DevicePlatform) Type() protoreflect.EnumType {
	return &file_meowchat_basic_basic_proto_enumTypes[0]
}

func (x DevicePlatform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DevicePlatform.Descriptor instead.
func (DevicePlatform) EnumDescriptor() ([]byte, []int) {
	return file_meowchat_basic_basic_proto_rawDescGZIP(), []int{0}
}

type UserMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string          `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty" form:"userId" query:"userId"`
	AppId           int64           `protobuf:"varint,2,opt,name=appId,proto3" json:"appId,omitempty" form:"appId" query:"appId"`
	DeviceId        string          `protobuf:"bytes,3,opt,name=deviceId,proto3" json:"deviceId,omitempty" form:"deviceId" query:"deviceId"`
	SessionUserId   string          `protobuf:"bytes,4,opt,name=sessionUserId,proto3" json:"sessionUserId,omitempty" form:"sessionUserId" query:"sessionUserId"`
	SessionAppId    int64           `protobuf:"varint,5,opt,name=sessionAppId,proto3" json:"sessionAppId,omitempty" form:"sessionAppId" query:"sessionAppId"`
	SessionDeviceId string          `protobuf:"bytes,6,opt,name=sessionDeviceId,proto3" json:"sessionDeviceId,omitempty" form:"sessionDeviceId" query:"sessionDeviceId"`
	IsLogin         bool            `protobuf:"varint,7,opt,name=isLogin,proto3" json:"isLogin,omitempty" form:"isLogin" query:"isLogin"`
	WechatUserMeta  *WechatUserMeta `protobuf:"bytes,8,opt,name=wechatUserMeta,proto3,oneof" json:"wechatUserMeta,omitempty" form:"wechatUserMeta" query:"wechatUserMeta"`
}

func (x *UserMeta) Reset() {
	*x = UserMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_basic_basic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMeta) ProtoMessage() {}

func (x *UserMeta) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_basic_basic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMeta.ProtoReflect.Descriptor instead.
func (*UserMeta) Descriptor() ([]byte, []int) {
	return file_meowchat_basic_basic_proto_rawDescGZIP(), []int{0}
}

func (x *UserMeta) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserMeta) GetAppId() int64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *UserMeta) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *UserMeta) GetSessionUserId() string {
	if x != nil {
		return x.SessionUserId
	}
	return ""
}

func (x *UserMeta) GetSessionAppId() int64 {
	if x != nil {
		return x.SessionAppId
	}
	return 0
}

func (x *UserMeta) GetSessionDeviceId() string {
	if x != nil {
		return x.SessionDeviceId
	}
	return ""
}

func (x *UserMeta) GetIsLogin() bool {
	if x != nil {
		return x.IsLogin
	}
	return false
}

func (x *UserMeta) GetWechatUserMeta() *WechatUserMeta {
	if x != nil {
		return x.WechatUserMeta
	}
	return nil
}

// 客户端透传
type Extra struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientIP       string `protobuf:"bytes,1,opt,name=clientIP,proto3" json:"clientIP,omitempty" form:"clientIP" query:"clientIP"`
	Language       string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty" form:"language" query:"language"`
	Resolution     string `protobuf:"bytes,3,opt,name=resolution,proto3" json:"resolution,omitempty" form:"resolution" query:"resolution"`                 // 像素比
	DevicePlatform string `protobuf:"bytes,4,opt,name=devicePlatform,proto3" json:"devicePlatform,omitempty" form:"devicePlatform" query:"devicePlatform"` // 设备平台
	DeviceBrand    string `protobuf:"bytes,5,opt,name=deviceBrand,proto3" json:"deviceBrand,omitempty" form:"deviceBrand" query:"deviceBrand"`             // 设备品牌
	DeviceId       string `protobuf:"bytes,6,opt,name=deviceId,proto3" json:"deviceId,omitempty" form:"deviceId" query:"deviceId"`
	OperateSystem  string `protobuf:"bytes,7,opt,name=operateSystem,proto3" json:"operateSystem,omitempty" form:"operateSystem" query:"operateSystem"`
}

func (x *Extra) Reset() {
	*x = Extra{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_basic_basic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extra) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extra) ProtoMessage() {}

func (x *Extra) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_basic_basic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extra.ProtoReflect.Descriptor instead.
func (*Extra) Descriptor() ([]byte, []int) {
	return file_meowchat_basic_basic_proto_rawDescGZIP(), []int{1}
}

func (x *Extra) GetClientIP() string {
	if x != nil {
		return x.ClientIP
	}
	return ""
}

func (x *Extra) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Extra) GetResolution() string {
	if x != nil {
		return x.Resolution
	}
	return ""
}

func (x *Extra) GetDevicePlatform() string {
	if x != nil {
		return x.DevicePlatform
	}
	return ""
}

func (x *Extra) GetDeviceBrand() string {
	if x != nil {
		return x.DeviceBrand
	}
	return ""
}

func (x *Extra) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Extra) GetOperateSystem() string {
	if x != nil {
		return x.OperateSystem
	}
	return ""
}

type WechatUserMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId      string `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty" form:"appId" query:"appId"`
	OpenId     string `protobuf:"bytes,2,opt,name=openId,proto3" json:"openId,omitempty" form:"openId" query:"openId"`
	PlatformId string `protobuf:"bytes,3,opt,name=platformId,proto3" json:"platformId,omitempty" form:"platformId" query:"platformId"`
	UnionId    string `protobuf:"bytes,4,opt,name=unionId,proto3" json:"unionId,omitempty" form:"unionId" query:"unionId"`
}

func (x *WechatUserMeta) Reset() {
	*x = WechatUserMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_basic_basic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WechatUserMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WechatUserMeta) ProtoMessage() {}

func (x *WechatUserMeta) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_basic_basic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WechatUserMeta.ProtoReflect.Descriptor instead.
func (*WechatUserMeta) Descriptor() ([]byte, []int) {
	return file_meowchat_basic_basic_proto_rawDescGZIP(), []int{2}
}

func (x *WechatUserMeta) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *WechatUserMeta) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *WechatUserMeta) GetPlatformId() string {
	if x != nil {
		return x.PlatformId
	}
	return ""
}

func (x *WechatUserMeta) GetUnionId() string {
	if x != nil {
		return x.UnionId
	}
	return ""
}

type UserPreview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	Nickname  string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty" form:"nickname" query:"nickname"`
	AvatarUrl string `protobuf:"bytes,3,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty" form:"avatarUrl" query:"avatarUrl"`
}

func (x *UserPreview) Reset() {
	*x = UserPreview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_basic_basic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPreview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPreview) ProtoMessage() {}

func (x *UserPreview) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_basic_basic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPreview.ProtoReflect.Descriptor instead.
func (*UserPreview) Descriptor() ([]byte, []int) {
	return file_meowchat_basic_basic_proto_rawDescGZIP(), []int{3}
}

func (x *UserPreview) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserPreview) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserPreview) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleType    string  `protobuf:"bytes,1,opt,name=roleType,proto3" json:"roleType,omitempty" form:"roleType" query:"roleType"`
	CommunityId *string `protobuf:"bytes,2,opt,name=communityId,proto3,oneof" json:"communityId,omitempty" form:"communityId" query:"communityId"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_basic_basic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_basic_basic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_meowchat_basic_basic_proto_rawDescGZIP(), []int{4}
}

func (x *Role) GetRoleType() string {
	if x != nil {
		return x.RoleType
	}
	return ""
}

func (x *Role) GetCommunityId() string {
	if x != nil && x.CommunityId != nil {
		return *x.CommunityId
	}
	return ""
}

type PaginationOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      *int64  `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty" form:"page" query:"page"`
	Limit     *int64  `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty" form:"limit" query:"limit"`
	LastToken *string `protobuf:"bytes,3,opt,name=lastToken,proto3,oneof" json:"lastToken,omitempty" form:"lastToken" query:"lastToken"`
	Backward  *bool   `protobuf:"varint,4,opt,name=backward,proto3,oneof" json:"backward,omitempty" form:"backward" query:"backward"`
	Offset    *int64  `protobuf:"varint,5,opt,name=offset,proto3,oneof" json:"offset,omitempty" form:"offset" query:"offset"`
}

func (x *PaginationOptions) Reset() {
	*x = PaginationOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_basic_basic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationOptions) ProtoMessage() {}

func (x *PaginationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_basic_basic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationOptions.ProtoReflect.Descriptor instead.
func (*PaginationOptions) Descriptor() ([]byte, []int) {
	return file_meowchat_basic_basic_proto_rawDescGZIP(), []int{5}
}

func (x *PaginationOptions) GetPage() int64 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *PaginationOptions) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *PaginationOptions) GetLastToken() string {
	if x != nil && x.LastToken != nil {
		return *x.LastToken
	}
	return ""
}

func (x *PaginationOptions) GetBackward() bool {
	if x != nil && x.Backward != nil {
		return *x.Backward
	}
	return false
}

func (x *PaginationOptions) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

var File_meowchat_basic_basic_proto protoreflect.FileDescriptor

var file_meowchat_basic_basic_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x65,
	0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x22, 0xc2, 0x02, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x0f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x4b, 0x0a, 0x0e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x65, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x6f, 0x77,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x57, 0x65, 0x63, 0x68, 0x61,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0e, 0x77, 0x65, 0x63,
	0x68, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x11,
	0x0a, 0x0f, 0x5f, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x61, 0x22, 0xeb, 0x01, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x50, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22,
	0x78, 0x0a, 0x0e, 0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55,
	0x72, 0x6c, 0x22, 0x59, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0xe1, 0x01,
	0x0a, 0x11, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x62, 0x61,
	0x63, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x08,
	0x62, 0x61, 0x63, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x04, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x6c, 0x61, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x2a, 0x50, 0x0a, 0x0e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x49, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x69, 0x6e,
	0x64, 0x6f, 0x77, 0x73, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x61, 0x63, 0x10, 0x03, 0x12,
	0x12, 0x0a, 0x0e, 0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x44, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x10, 0x04, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x62,
	0x69, 0x7a, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64,
	0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meowchat_basic_basic_proto_rawDescOnce sync.Once
	file_meowchat_basic_basic_proto_rawDescData = file_meowchat_basic_basic_proto_rawDesc
)

func file_meowchat_basic_basic_proto_rawDescGZIP() []byte {
	file_meowchat_basic_basic_proto_rawDescOnce.Do(func() {
		file_meowchat_basic_basic_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowchat_basic_basic_proto_rawDescData)
	})
	return file_meowchat_basic_basic_proto_rawDescData
}

var file_meowchat_basic_basic_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_meowchat_basic_basic_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_meowchat_basic_basic_proto_goTypes = []interface{}{
	(DevicePlatform)(0),       // 0: meowchat.basic.DevicePlatform
	(*UserMeta)(nil),          // 1: meowchat.basic.UserMeta
	(*Extra)(nil),             // 2: meowchat.basic.Extra
	(*WechatUserMeta)(nil),    // 3: meowchat.basic.WechatUserMeta
	(*UserPreview)(nil),       // 4: meowchat.basic.UserPreview
	(*Role)(nil),              // 5: meowchat.basic.Role
	(*PaginationOptions)(nil), // 6: meowchat.basic.PaginationOptions
}
var file_meowchat_basic_basic_proto_depIdxs = []int32{
	3, // 0: meowchat.basic.UserMeta.wechatUserMeta:type_name -> meowchat.basic.WechatUserMeta
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func file_meowchat_basic_basic_proto_init() {
	if File_meowchat_basic_basic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meowchat_basic_basic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMeta); i {
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
		file_meowchat_basic_basic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extra); i {
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
		file_meowchat_basic_basic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WechatUserMeta); i {
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
		file_meowchat_basic_basic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPreview); i {
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
		file_meowchat_basic_basic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_meowchat_basic_basic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationOptions); i {
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
	file_meowchat_basic_basic_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_meowchat_basic_basic_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_meowchat_basic_basic_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meowchat_basic_basic_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meowchat_basic_basic_proto_goTypes,
		DependencyIndexes: file_meowchat_basic_basic_proto_depIdxs,
		EnumInfos:         file_meowchat_basic_basic_proto_enumTypes,
		MessageInfos:      file_meowchat_basic_basic_proto_msgTypes,
	}.Build()
	File_meowchat_basic_basic_proto = out.File
	file_meowchat_basic_basic_proto_rawDesc = nil
	file_meowchat_basic_basic_proto_goTypes = nil
	file_meowchat_basic_basic_proto_depIdxs = nil
}
