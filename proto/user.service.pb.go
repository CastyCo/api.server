// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.service.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	messages "movie-night/proto/messages"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UpdateStateRequest struct {
	State                messages.PERSONAL_STATE `protobuf:"varint,1,opt,name=state,proto3,enum=messages.PERSONAL_STATE" json:"state,omitempty"`
	AuthRequest          *AuthenticateRequest    `protobuf:"bytes,2,opt,name=auth_request,json=authRequest,proto3" json:"auth_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *UpdateStateRequest) Reset()         { *m = UpdateStateRequest{} }
func (m *UpdateStateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateStateRequest) ProtoMessage()    {}
func (*UpdateStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ed9320702f2c5, []int{0}
}

func (m *UpdateStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStateRequest.Unmarshal(m, b)
}
func (m *UpdateStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStateRequest.Merge(m, src)
}
func (m *UpdateStateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateStateRequest.Size(m)
}
func (m *UpdateStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStateRequest proto.InternalMessageInfo

func (m *UpdateStateRequest) GetState() messages.PERSONAL_STATE {
	if m != nil {
		return m.State
	}
	return messages.PERSONAL_STATE_OFFLINE
}

func (m *UpdateStateRequest) GetAuthRequest() *AuthenticateRequest {
	if m != nil {
		return m.AuthRequest
	}
	return nil
}

type UpdateActivityRequest struct {
	Activity             *messages.Activity   `protobuf:"bytes,1,opt,name=activity,proto3" json:"activity,omitempty"`
	AuthRequest          *AuthenticateRequest `protobuf:"bytes,2,opt,name=auth_request,json=authRequest,proto3" json:"auth_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpdateActivityRequest) Reset()         { *m = UpdateActivityRequest{} }
func (m *UpdateActivityRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateActivityRequest) ProtoMessage()    {}
func (*UpdateActivityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ed9320702f2c5, []int{1}
}

func (m *UpdateActivityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateActivityRequest.Unmarshal(m, b)
}
func (m *UpdateActivityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateActivityRequest.Marshal(b, m, deterministic)
}
func (m *UpdateActivityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateActivityRequest.Merge(m, src)
}
func (m *UpdateActivityRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateActivityRequest.Size(m)
}
func (m *UpdateActivityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateActivityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateActivityRequest proto.InternalMessageInfo

func (m *UpdateActivityRequest) GetActivity() *messages.Activity {
	if m != nil {
		return m.Activity
	}
	return nil
}

func (m *UpdateActivityRequest) GetAuthRequest() *AuthenticateRequest {
	if m != nil {
		return m.AuthRequest
	}
	return nil
}

type FriendRequest struct {
	FriendId             string               `protobuf:"bytes,1,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	AuthRequest          *AuthenticateRequest `protobuf:"bytes,2,opt,name=auth_request,json=authRequest,proto3" json:"auth_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FriendRequest) Reset()         { *m = FriendRequest{} }
func (m *FriendRequest) String() string { return proto.CompactTextString(m) }
func (*FriendRequest) ProtoMessage()    {}
func (*FriendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ed9320702f2c5, []int{2}
}

func (m *FriendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendRequest.Unmarshal(m, b)
}
func (m *FriendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendRequest.Marshal(b, m, deterministic)
}
func (m *FriendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendRequest.Merge(m, src)
}
func (m *FriendRequest) XXX_Size() int {
	return xxx_messageInfo_FriendRequest.Size(m)
}
func (m *FriendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FriendRequest proto.InternalMessageInfo

func (m *FriendRequest) GetFriendId() string {
	if m != nil {
		return m.FriendId
	}
	return ""
}

func (m *FriendRequest) GetAuthRequest() *AuthenticateRequest {
	if m != nil {
		return m.AuthRequest
	}
	return nil
}

type CreateUserRequest struct {
	User                 *messages.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ed9320702f2c5, []int{3}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *messages.User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetFriendRequest struct {
	User                 *messages.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetFriendRequest) Reset()         { *m = GetFriendRequest{} }
func (m *GetFriendRequest) String() string { return proto.CompactTextString(m) }
func (*GetFriendRequest) ProtoMessage()    {}
func (*GetFriendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ed9320702f2c5, []int{4}
}

func (m *GetFriendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFriendRequest.Unmarshal(m, b)
}
func (m *GetFriendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFriendRequest.Marshal(b, m, deterministic)
}
func (m *GetFriendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFriendRequest.Merge(m, src)
}
func (m *GetFriendRequest) XXX_Size() int {
	return xxx_messageInfo_GetFriendRequest.Size(m)
}
func (m *GetFriendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFriendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFriendRequest proto.InternalMessageInfo

func (m *GetFriendRequest) GetUser() *messages.User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserResponse struct {
	Code                 int64          `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status               string         `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Message              string         `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Result               *messages.User `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ed9320702f2c5, []int{5}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetUserResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GetUserResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetUserResponse) GetResult() *messages.User {
	if m != nil {
		return m.Result
	}
	return nil
}

type FriendsResponse struct {
	Code                 int64            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status               string           `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Message              string           `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Result               []*messages.User `protobuf:"bytes,4,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FriendsResponse) Reset()         { *m = FriendsResponse{} }
func (m *FriendsResponse) String() string { return proto.CompactTextString(m) }
func (*FriendsResponse) ProtoMessage()    {}
func (*FriendsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ed9320702f2c5, []int{6}
}

func (m *FriendsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendsResponse.Unmarshal(m, b)
}
func (m *FriendsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendsResponse.Marshal(b, m, deterministic)
}
func (m *FriendsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendsResponse.Merge(m, src)
}
func (m *FriendsResponse) XXX_Size() int {
	return xxx_messageInfo_FriendsResponse.Size(m)
}
func (m *FriendsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FriendsResponse proto.InternalMessageInfo

func (m *FriendsResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *FriendsResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *FriendsResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *FriendsResponse) GetResult() []*messages.User {
	if m != nil {
		return m.Result
	}
	return nil
}

type FriendResponse struct {
	Code                 int64          `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status               string         `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Message              string         `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Result               *messages.User `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FriendResponse) Reset()         { *m = FriendResponse{} }
func (m *FriendResponse) String() string { return proto.CompactTextString(m) }
func (*FriendResponse) ProtoMessage()    {}
func (*FriendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ed9320702f2c5, []int{7}
}

func (m *FriendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendResponse.Unmarshal(m, b)
}
func (m *FriendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendResponse.Marshal(b, m, deterministic)
}
func (m *FriendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendResponse.Merge(m, src)
}
func (m *FriendResponse) XXX_Size() int {
	return xxx_messageInfo_FriendResponse.Size(m)
}
func (m *FriendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FriendResponse proto.InternalMessageInfo

func (m *FriendResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *FriendResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *FriendResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *FriendResponse) GetResult() *messages.User {
	if m != nil {
		return m.Result
	}
	return nil
}

type SearchUserRequest struct {
	Keyword              string               `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	AuthRequest          *AuthenticateRequest `protobuf:"bytes,2,opt,name=auth_request,json=authRequest,proto3" json:"auth_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SearchUserRequest) Reset()         { *m = SearchUserRequest{} }
func (m *SearchUserRequest) String() string { return proto.CompactTextString(m) }
func (*SearchUserRequest) ProtoMessage()    {}
func (*SearchUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ed9320702f2c5, []int{8}
}

func (m *SearchUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchUserRequest.Unmarshal(m, b)
}
func (m *SearchUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchUserRequest.Marshal(b, m, deterministic)
}
func (m *SearchUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchUserRequest.Merge(m, src)
}
func (m *SearchUserRequest) XXX_Size() int {
	return xxx_messageInfo_SearchUserRequest.Size(m)
}
func (m *SearchUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchUserRequest proto.InternalMessageInfo

func (m *SearchUserRequest) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func (m *SearchUserRequest) GetAuthRequest() *AuthenticateRequest {
	if m != nil {
		return m.AuthRequest
	}
	return nil
}

type SearchUserResponse struct {
	Code                 int64            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status               string           `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Message              string           `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Result               []*messages.User `protobuf:"bytes,4,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SearchUserResponse) Reset()         { *m = SearchUserResponse{} }
func (m *SearchUserResponse) String() string { return proto.CompactTextString(m) }
func (*SearchUserResponse) ProtoMessage()    {}
func (*SearchUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ed9320702f2c5, []int{9}
}

func (m *SearchUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchUserResponse.Unmarshal(m, b)
}
func (m *SearchUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchUserResponse.Marshal(b, m, deterministic)
}
func (m *SearchUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchUserResponse.Merge(m, src)
}
func (m *SearchUserResponse) XXX_Size() int {
	return xxx_messageInfo_SearchUserResponse.Size(m)
}
func (m *SearchUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchUserResponse proto.InternalMessageInfo

func (m *SearchUserResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SearchUserResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *SearchUserResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *SearchUserResponse) GetResult() []*messages.User {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateStateRequest)(nil), "proto.UpdateStateRequest")
	proto.RegisterType((*UpdateActivityRequest)(nil), "proto.UpdateActivityRequest")
	proto.RegisterType((*FriendRequest)(nil), "proto.FriendRequest")
	proto.RegisterType((*CreateUserRequest)(nil), "proto.CreateUserRequest")
	proto.RegisterType((*GetFriendRequest)(nil), "proto.GetFriendRequest")
	proto.RegisterType((*GetUserResponse)(nil), "proto.GetUserResponse")
	proto.RegisterType((*FriendsResponse)(nil), "proto.FriendsResponse")
	proto.RegisterType((*FriendResponse)(nil), "proto.FriendResponse")
	proto.RegisterType((*SearchUserRequest)(nil), "proto.SearchUserRequest")
	proto.RegisterType((*SearchUserResponse)(nil), "proto.SearchUserResponse")
}

func init() { proto.RegisterFile("user.service.proto", fileDescriptor_1e7ed9320702f2c5) }

var fileDescriptor_1e7ed9320702f2c5 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xdd, 0x6b, 0xd3, 0x50,
	0x14, 0x27, 0xb6, 0x6b, 0xd7, 0x13, 0xcd, 0xdc, 0x99, 0x1b, 0x59, 0xf4, 0x61, 0xf4, 0x41, 0xf6,
	0x14, 0xa1, 0x82, 0x9f, 0x14, 0x2d, 0x32, 0x87, 0x20, 0x2a, 0xc9, 0xf6, 0x5c, 0xb2, 0xe4, 0x68,
	0xc3, 0xb6, 0xa6, 0xde, 0x7b, 0x53, 0xd9, 0x8b, 0x82, 0x82, 0xff, 0x94, 0xff, 0x9c, 0xe4, 0xde,
	0x9b, 0xa4, 0x69, 0x42, 0x19, 0x94, 0xb1, 0xa7, 0xe4, 0x7c, 0xfe, 0x7e, 0xe7, 0x9e, 0x0f, 0xc0,
	0x94, 0x13, 0x73, 0x39, 0xb1, 0x79, 0x1c, 0x92, 0x3b, 0x63, 0x89, 0x48, 0x70, 0x43, 0x7e, 0x1c,
	0x38, 0x0b, 0xb8, 0x56, 0x39, 0x18, 0xa4, 0x62, 0x52, 0x75, 0x73, 0x76, 0x2e, 0x89, 0xf3, 0xe0,
	0x1b, 0xf1, 0x27, 0x32, 0x87, 0x54, 0xf6, 0xff, 0x18, 0x80, 0xa7, 0xb3, 0x28, 0x10, 0xe4, 0x8b,
	0x40, 0x90, 0x47, 0xdf, 0x53, 0xe2, 0x02, 0x5d, 0xd8, 0xe0, 0x99, 0x6c, 0x1b, 0x07, 0xc6, 0xa1,
	0x35, 0xb0, 0xdd, 0x3c, 0xd6, 0xfd, 0x72, 0xe4, 0xf9, 0x9f, 0x3f, 0x8d, 0x3e, 0x8e, 0xfd, 0x93,
	0xd1, 0xc9, 0x91, 0xa7, 0xdc, 0x70, 0x08, 0x77, 0x33, 0xc4, 0x31, 0x53, 0xf1, 0xf6, 0x9d, 0x03,
	0xe3, 0xd0, 0x1c, 0x38, 0x0a, 0xc4, 0x1d, 0xa5, 0x62, 0x42, 0x53, 0x11, 0x87, 0x25, 0x82, 0x67,
	0x66, 0xfe, 0x5a, 0xe8, 0xff, 0x35, 0x60, 0x57, 0xb1, 0x18, 0x85, 0x22, 0x9e, 0xc7, 0xe2, 0xaa,
	0x24, 0xb2, 0x19, 0x68, 0x95, 0xe4, 0x62, 0x0e, 0xb0, 0xe4, 0x52, 0x38, 0x17, 0x3e, 0xeb, 0x12,
	0x39, 0x87, 0x7b, 0xef, 0x59, 0x4c, 0xd3, 0x28, 0xc7, 0x7f, 0x08, 0xbd, 0xaf, 0x52, 0x31, 0x8e,
	0x23, 0x49, 0xa0, 0xe7, 0x6d, 0x2a, 0xc5, 0x87, 0x68, 0x5d, 0xb0, 0xe7, 0xb0, 0xfd, 0x8e, 0x51,
	0x20, 0xe8, 0x94, 0x13, 0xcb, 0x01, 0xfb, 0xd0, 0xce, 0xda, 0xa3, 0x8b, 0xb5, 0xca, 0x62, 0xa5,
	0x93, 0xb4, 0xf5, 0x9f, 0xc1, 0xfd, 0x63, 0x12, 0x55, 0xa2, 0xd7, 0x89, 0xfb, 0x05, 0x5b, 0xc7,
	0x24, 0x14, 0x1a, 0x9f, 0x25, 0x53, 0x4e, 0x88, 0xd0, 0x0e, 0x93, 0x48, 0xf5, 0xb9, 0xe5, 0xc9,
	0x7f, 0xdc, 0x83, 0x4e, 0xd6, 0xd5, 0x94, 0xcb, 0x82, 0x7a, 0x9e, 0x96, 0xd0, 0x86, 0xae, 0xce,
	0x6a, 0xb7, 0xa4, 0x21, 0x17, 0xf1, 0x31, 0x74, 0x18, 0xf1, 0xf4, 0x42, 0xd8, 0xed, 0x46, 0x78,
	0x6d, 0xcd, 0x08, 0x28, 0xd6, 0xfc, 0x06, 0x09, 0xb4, 0x56, 0x10, 0xf8, 0x09, 0x56, 0xfe, 0x6c,
	0xb7, 0xf2, 0x00, 0x17, 0xb0, 0xed, 0x53, 0xc0, 0xc2, 0xc9, 0x62, 0xcb, 0x6d, 0xe8, 0x9e, 0xd3,
	0xd5, 0x8f, 0x84, 0xe5, 0x13, 0x96, 0x8b, 0xeb, 0x0e, 0xd8, 0x6f, 0x03, 0x70, 0x11, 0xee, 0x36,
	0x9e, 0x7c, 0xf0, 0xaf, 0x0d, 0x66, 0xa6, 0xf0, 0xd5, 0x31, 0xc2, 0xd7, 0x00, 0xe5, 0xd4, 0xa3,
	0xad, 0x6b, 0xa9, 0x2d, 0x82, 0xb3, 0xb3, 0x50, 0x65, 0x41, 0xfd, 0x25, 0x98, 0x0b, 0xd7, 0x0a,
	0xf7, 0xb5, 0x4f, 0xfd, 0x82, 0x39, 0x5b, 0xda, 0x54, 0x84, 0xbe, 0x01, 0xab, 0x7a, 0x62, 0xf0,
	0x51, 0x25, 0x7a, 0xe9, 0xf2, 0xd4, 0x13, 0x0c, 0xc1, 0xf2, 0xe8, 0x32, 0x99, 0x97, 0x09, 0x56,
	0x34, 0xa2, 0x29, 0xbc, 0xab, 0x97, 0x6f, 0x65, 0xdc, 0x9e, 0xb6, 0x2d, 0x2f, 0xea, 0x0b, 0xe8,
	0x15, 0x3b, 0x8f, 0x0f, 0xb4, 0x53, 0xe5, 0x04, 0x38, 0xbb, 0x4b, 0x5a, 0x1d, 0xf9, 0x2a, 0x9b,
	0xb9, 0x69, 0x54, 0x3d, 0x17, 0xcd, 0x19, 0x6a, 0xa4, 0xdf, 0x02, 0x14, 0xa8, 0xfc, 0x5a, 0xbc,
	0x97, 0xf7, 0x7b, 0x08, 0x1d, 0x35, 0x82, 0x45, 0xab, 0x6b, 0x0b, 0xe0, 0xec, 0x37, 0x58, 0x54,
	0xf8, 0x59, 0x47, 0x5a, 0x9e, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x74, 0x64, 0x8b, 0x2d, 0xf8,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	UpdateState(ctx context.Context, in *UpdateStateRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateActivity(ctx context.Context, in *UpdateActivityRequest, opts ...grpc.CallOption) (*Response, error)
	RemoveActivity(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*Response, error)
	GetUser(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetFriend(ctx context.Context, in *FriendRequest, opts ...grpc.CallOption) (*FriendResponse, error)
	SendFriendRequest(ctx context.Context, in *FriendRequest, opts ...grpc.CallOption) (*Response, error)
	GetFriends(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*FriendsResponse, error)
	Search(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateState(ctx context.Context, in *UpdateStateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.UserService/UpdateState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateActivity(ctx context.Context, in *UpdateActivityRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.UserService/UpdateActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RemoveActivity(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.UserService/RemoveActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetFriend(ctx context.Context, in *FriendRequest, opts ...grpc.CallOption) (*FriendResponse, error) {
	out := new(FriendResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SendFriendRequest(ctx context.Context, in *FriendRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.UserService/SendFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetFriends(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*FriendsResponse, error) {
	out := new(FriendsResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Search(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error) {
	out := new(SearchUserResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*AuthResponse, error)
	UpdateState(context.Context, *UpdateStateRequest) (*Response, error)
	UpdateActivity(context.Context, *UpdateActivityRequest) (*Response, error)
	RemoveActivity(context.Context, *AuthenticateRequest) (*Response, error)
	GetUser(context.Context, *AuthenticateRequest) (*GetUserResponse, error)
	GetFriend(context.Context, *FriendRequest) (*FriendResponse, error)
	SendFriendRequest(context.Context, *FriendRequest) (*Response, error)
	GetFriends(context.Context, *AuthenticateRequest) (*FriendsResponse, error)
	Search(context.Context, *SearchUserRequest) (*SearchUserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServiceServer) UpdateState(ctx context.Context, req *UpdateStateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateState not implemented")
}
func (*UnimplementedUserServiceServer) UpdateActivity(ctx context.Context, req *UpdateActivityRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateActivity not implemented")
}
func (*UnimplementedUserServiceServer) RemoveActivity(ctx context.Context, req *AuthenticateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveActivity not implemented")
}
func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *AuthenticateRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServiceServer) GetFriend(ctx context.Context, req *FriendRequest) (*FriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriend not implemented")
}
func (*UnimplementedUserServiceServer) SendFriendRequest(ctx context.Context, req *FriendRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendFriendRequest not implemented")
}
func (*UnimplementedUserServiceServer) GetFriends(ctx context.Context, req *AuthenticateRequest) (*FriendsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriends not implemented")
}
func (*UnimplementedUserServiceServer) Search(ctx context.Context, req *SearchUserRequest) (*SearchUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/UpdateState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateState(ctx, req.(*UpdateStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/UpdateActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateActivity(ctx, req.(*UpdateActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RemoveActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RemoveActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/RemoveActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RemoveActivity(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetFriend(ctx, req.(*FriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SendFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SendFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/SendFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SendFriendRequest(ctx, req.(*FriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetFriends(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Search(ctx, req.(*SearchUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateState",
			Handler:    _UserService_UpdateState_Handler,
		},
		{
			MethodName: "UpdateActivity",
			Handler:    _UserService_UpdateActivity_Handler,
		},
		{
			MethodName: "RemoveActivity",
			Handler:    _UserService_RemoveActivity_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetFriend",
			Handler:    _UserService_GetFriend_Handler,
		},
		{
			MethodName: "SendFriendRequest",
			Handler:    _UserService_SendFriendRequest_Handler,
		},
		{
			MethodName: "GetFriends",
			Handler:    _UserService_GetFriends_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _UserService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.service.proto",
}
