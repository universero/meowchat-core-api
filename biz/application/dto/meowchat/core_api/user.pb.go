// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.23.4
// source: meowchat/core_api/user.proto

package core_api

import (
	basic "github.com/xh-polaris/meowchat-core-api/biz/application/dto/basic"
	system "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/system"
	user "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	Nickname  string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty" form:"nickname" query:"nickname"`
	AvatarUrl string `protobuf:"bytes,3,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty" form:"avatarUrl" query:"avatarUrl"`
	Motto     string `protobuf:"bytes,4,opt,name=motto,proto3" json:"motto,omitempty" form:"motto" query:"motto"`
	Follower  int64  `protobuf:"varint,5,opt,name=follower,proto3" json:"follower,omitempty" form:"follower" query:"follower"`
	Following int64  `protobuf:"varint,6,opt,name=following,proto3" json:"following,omitempty" form:"following" query:"following"`
	Article   int64  `protobuf:"varint,7,opt,name=article,proto3" json:"article,omitempty" form:"article" query:"article"`
	Like      int64  `protobuf:"varint,8,opt,name=like,proto3" json:"like,omitempty" form:"like" query:"like"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *User) GetMotto() string {
	if x != nil {
		return x.Motto
	}
	return ""
}

func (x *User) GetFollower() int64 {
	if x != nil {
		return x.Follower
	}
	return 0
}

func (x *User) GetFollowing() int64 {
	if x != nil {
		return x.Following
	}
	return 0
}

func (x *User) GetArticle() int64 {
	if x != nil {
		return x.Article
	}
	return 0
}

func (x *User) GetLike() int64 {
	if x != nil {
		return x.Like
	}
	return 0
}

type UserPreviewWithRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  *user.UserPreview `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty" form:"user" query:"user"`
	Roles []*system.Role    `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty" form:"roles" query:"roles"`
}

func (x *UserPreviewWithRole) Reset() {
	*x = UserPreviewWithRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPreviewWithRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPreviewWithRole) ProtoMessage() {}

func (x *UserPreviewWithRole) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPreviewWithRole.ProtoReflect.Descriptor instead.
func (*UserPreviewWithRole) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserPreviewWithRole) GetUser() *user.UserPreview {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserPreviewWithRole) GetRoles() []*system.Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

type GetUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId *string `protobuf:"bytes,1,opt,name=userId,proto3,oneof" json:"userId,omitempty" form:"userId" query:"userId"`
}

func (x *GetUserInfoReq) Reset() {
	*x = GetUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReq) ProtoMessage() {}

func (x *GetUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReq.ProtoReflect.Descriptor instead.
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserInfoReq) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

type GetUserInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User        *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty" form:"user" query:"user"`
	EnableDebug bool  `protobuf:"varint,2,opt,name=enableDebug,proto3" json:"enableDebug,omitempty" form:"enableDebug" query:"enableDebug"`
}

func (x *GetUserInfoResp) Reset() {
	*x = GetUserInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResp) ProtoMessage() {}

func (x *GetUserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResp.ProtoReflect.Descriptor instead.
func (*GetUserInfoResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserInfoResp) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GetUserInfoResp) GetEnableDebug() bool {
	if x != nil {
		return x.EnableDebug
	}
	return false
}

type UpdateUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarUrl *string `protobuf:"bytes,1,opt,name=avatarUrl,proto3,oneof" json:"avatarUrl,omitempty" form:"avatarUrl" query:"avatarUrl"`
	Nickname  *string `protobuf:"bytes,2,opt,name=nickname,proto3,oneof" json:"nickname,omitempty" form:"nickname" query:"nickname"`
	Motto     *string `protobuf:"bytes,3,opt,name=motto,proto3,oneof" json:"motto,omitempty" form:"motto" query:"motto"`
}

func (x *UpdateUserInfoReq) Reset() {
	*x = UpdateUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfoReq) ProtoMessage() {}

func (x *UpdateUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfoReq.ProtoReflect.Descriptor instead.
func (*UpdateUserInfoReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserInfoReq) GetAvatarUrl() string {
	if x != nil && x.AvatarUrl != nil {
		return *x.AvatarUrl
	}
	return ""
}

func (x *UpdateUserInfoReq) GetNickname() string {
	if x != nil && x.Nickname != nil {
		return *x.Nickname
	}
	return ""
}

func (x *UpdateUserInfoReq) GetMotto() string {
	if x != nil && x.Motto != nil {
		return *x.Motto
	}
	return ""
}

type UpdateUserInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUserInfoResp) Reset() {
	*x = UpdateUserInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfoResp) ProtoMessage() {}

func (x *UpdateUserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfoResp.ProtoReflect.Descriptor instead.
func (*UpdateUserInfoResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{5}
}

type SearchUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword          string                   `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty" form:"keyword" query:"keyword"`
	PaginationOption *basic.PaginationOptions `protobuf:"bytes,2,opt,name=paginationOption,proto3" json:"paginationOption,omitempty" form:"paginationOption" query:"paginationOption"`
}

func (x *SearchUserReq) Reset() {
	*x = SearchUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserReq) ProtoMessage() {}

func (x *SearchUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserReq.ProtoReflect.Descriptor instead.
func (*SearchUserReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{6}
}

func (x *SearchUserReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SearchUserReq) GetPaginationOption() *basic.PaginationOptions {
	if x != nil {
		return x.PaginationOption
	}
	return nil
}

type SearchUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*user.UserPreview `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty" form:"users" query:"users"`
	Total int64               `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty" form:"total" query:"total"`
	Token string              `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty" form:"token" query:"token"`
}

func (x *SearchUserResp) Reset() {
	*x = SearchUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserResp) ProtoMessage() {}

func (x *SearchUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserResp.ProtoReflect.Descriptor instead.
func (*SearchUserResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{7}
}

func (x *SearchUserResp) GetUsers() []*user.UserPreview {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *SearchUserResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SearchUserResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SearchUserForAdminReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword          string                   `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty" form:"keyword" query:"keyword"`
	PaginationOption *basic.PaginationOptions `protobuf:"bytes,2,opt,name=paginationOption,proto3" json:"paginationOption,omitempty" form:"paginationOption" query:"paginationOption"`
}

func (x *SearchUserForAdminReq) Reset() {
	*x = SearchUserForAdminReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserForAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserForAdminReq) ProtoMessage() {}

func (x *SearchUserForAdminReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserForAdminReq.ProtoReflect.Descriptor instead.
func (*SearchUserForAdminReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{8}
}

func (x *SearchUserForAdminReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SearchUserForAdminReq) GetPaginationOption() *basic.PaginationOptions {
	if x != nil {
		return x.PaginationOption
	}
	return nil
}

type SearchUserForAdminResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserPreviewWithRole `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty" form:"users" query:"users"`
	Total int64                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty" form:"total" query:"total"`
	Token string                 `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty" form:"token" query:"token"`
}

func (x *SearchUserForAdminResp) Reset() {
	*x = SearchUserForAdminResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserForAdminResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserForAdminResp) ProtoMessage() {}

func (x *SearchUserForAdminResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserForAdminResp.ProtoReflect.Descriptor instead.
func (*SearchUserForAdminResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{9}
}

func (x *SearchUserForAdminResp) GetUsers() []*UserPreviewWithRole {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *SearchUserForAdminResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SearchUserForAdminResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_meowchat_core_api_user_proto protoreflect.FileDescriptor

var file_meowchat_core_api_user_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x1a, 0x16, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x65, 0x6f, 0x77, 0x63,
	0x68, 0x61, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x6c, 0x69, 0x6b, 0x65, 0x22, 0x72, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x57, 0x69, 0x74, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x6f, 0x77,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x60, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x65, 0x62, 0x75,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x44,
	0x65, 0x62, 0x75, 0x67, 0x22, 0x97, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x09, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x05, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x22, 0x14,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x6f, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x44, 0x0a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6e, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x77, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18,
	0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x44, 0x0a, 0x10, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x10, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x82,
	0x01, 0x0a, 0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x72,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x57, 0x69, 0x74, 0x68, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x42, 0x83, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f,
	0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x61, 0x70, 0x69, 0x42, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69,
	0x73, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_meowchat_core_api_user_proto_rawDescOnce sync.Once
	file_meowchat_core_api_user_proto_rawDescData = file_meowchat_core_api_user_proto_rawDesc
)

func file_meowchat_core_api_user_proto_rawDescGZIP() []byte {
	file_meowchat_core_api_user_proto_rawDescOnce.Do(func() {
		file_meowchat_core_api_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowchat_core_api_user_proto_rawDescData)
	})
	return file_meowchat_core_api_user_proto_rawDescData
}

var file_meowchat_core_api_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_meowchat_core_api_user_proto_goTypes = []interface{}{
	(*User)(nil),                    // 0: meowchat.core_api.User
	(*UserPreviewWithRole)(nil),     // 1: meowchat.core_api.UserPreviewWithRole
	(*GetUserInfoReq)(nil),          // 2: meowchat.core_api.GetUserInfoReq
	(*GetUserInfoResp)(nil),         // 3: meowchat.core_api.GetUserInfoResp
	(*UpdateUserInfoReq)(nil),       // 4: meowchat.core_api.UpdateUserInfoReq
	(*UpdateUserInfoResp)(nil),      // 5: meowchat.core_api.UpdateUserInfoResp
	(*SearchUserReq)(nil),           // 6: meowchat.core_api.SearchUserReq
	(*SearchUserResp)(nil),          // 7: meowchat.core_api.SearchUserResp
	(*SearchUserForAdminReq)(nil),   // 8: meowchat.core_api.SearchUserForAdminReq
	(*SearchUserForAdminResp)(nil),  // 9: meowchat.core_api.SearchUserForAdminResp
	(*user.UserPreview)(nil),        // 10: meowchat.user.UserPreview
	(*system.Role)(nil),             // 11: meowchat.system.Role
	(*basic.PaginationOptions)(nil), // 12: basic.PaginationOptions
}
var file_meowchat_core_api_user_proto_depIdxs = []int32{
	10, // 0: meowchat.core_api.UserPreviewWithRole.user:type_name -> meowchat.user.UserPreview
	11, // 1: meowchat.core_api.UserPreviewWithRole.roles:type_name -> meowchat.system.Role
	0,  // 2: meowchat.core_api.GetUserInfoResp.user:type_name -> meowchat.core_api.User
	12, // 3: meowchat.core_api.SearchUserReq.paginationOption:type_name -> basic.PaginationOptions
	10, // 4: meowchat.core_api.SearchUserResp.users:type_name -> meowchat.user.UserPreview
	12, // 5: meowchat.core_api.SearchUserForAdminReq.paginationOption:type_name -> basic.PaginationOptions
	1,  // 6: meowchat.core_api.SearchUserForAdminResp.users:type_name -> meowchat.core_api.UserPreviewWithRole
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func file_meowchat_core_api_user_proto_init() {
	if File_meowchat_core_api_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meowchat_core_api_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_meowchat_core_api_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPreviewWithRole); i {
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
		file_meowchat_core_api_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoReq); i {
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
		file_meowchat_core_api_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoResp); i {
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
		file_meowchat_core_api_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInfoReq); i {
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
		file_meowchat_core_api_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInfoResp); i {
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
		file_meowchat_core_api_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserReq); i {
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
		file_meowchat_core_api_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserResp); i {
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
		file_meowchat_core_api_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserForAdminReq); i {
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
		file_meowchat_core_api_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserForAdminResp); i {
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
	file_meowchat_core_api_user_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_meowchat_core_api_user_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meowchat_core_api_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meowchat_core_api_user_proto_goTypes,
		DependencyIndexes: file_meowchat_core_api_user_proto_depIdxs,
		MessageInfos:      file_meowchat_core_api_user_proto_msgTypes,
	}.Build()
	File_meowchat_core_api_user_proto = out.File
	file_meowchat_core_api_user_proto_rawDesc = nil
	file_meowchat_core_api_user_proto_goTypes = nil
	file_meowchat_core_api_user_proto_depIdxs = nil
}
