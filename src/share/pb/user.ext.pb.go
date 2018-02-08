// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.ext.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RegistAccountReq struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=userName" json:"userName,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *RegistAccountReq) Reset()                    { *m = RegistAccountReq{} }
func (m *RegistAccountReq) String() string            { return proto.CompactTextString(m) }
func (*RegistAccountReq) ProtoMessage()               {}
func (*RegistAccountReq) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *RegistAccountReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegistAccountReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RegistAccountReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegistAccountRsp struct {
}

func (m *RegistAccountRsp) Reset()                    { *m = RegistAccountRsp{} }
func (m *RegistAccountRsp) String() string            { return proto.CompactTextString(m) }
func (*RegistAccountRsp) ProtoMessage()               {}
func (*RegistAccountRsp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

type LoginAccountReq struct {
	UserName string `protobuf:"bytes,1,opt,name=userName" json:"userName,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginAccountReq) Reset()                    { *m = LoginAccountReq{} }
func (m *LoginAccountReq) String() string            { return proto.CompactTextString(m) }
func (*LoginAccountReq) ProtoMessage()               {}
func (*LoginAccountReq) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *LoginAccountReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginAccountReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginAccountRsp struct {
}

func (m *LoginAccountRsp) Reset()                    { *m = LoginAccountRsp{} }
func (m *LoginAccountRsp) String() string            { return proto.CompactTextString(m) }
func (*LoginAccountRsp) ProtoMessage()               {}
func (*LoginAccountRsp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

type ResetAccountReq struct {
}

func (m *ResetAccountReq) Reset()                    { *m = ResetAccountReq{} }
func (m *ResetAccountReq) String() string            { return proto.CompactTextString(m) }
func (*ResetAccountReq) ProtoMessage()               {}
func (*ResetAccountReq) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

type ResetAccountRsp struct {
}

func (m *ResetAccountRsp) Reset()                    { *m = ResetAccountRsp{} }
func (m *ResetAccountRsp) String() string            { return proto.CompactTextString(m) }
func (*ResetAccountRsp) ProtoMessage()               {}
func (*ResetAccountRsp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

type WantScoreReq struct {
	UserId  int64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	MovieId int64 `protobuf:"varint,2,opt,name=movieId" json:"movieId,omitempty"`
	Score   int64 `protobuf:"varint,3,opt,name=score" json:"score,omitempty"`
}

func (m *WantScoreReq) Reset()                    { *m = WantScoreReq{} }
func (m *WantScoreReq) String() string            { return proto.CompactTextString(m) }
func (*WantScoreReq) ProtoMessage()               {}
func (*WantScoreReq) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *WantScoreReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *WantScoreReq) GetMovieId() int64 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

func (m *WantScoreReq) GetScore() int64 {
	if m != nil {
		return m.Score
	}
	return 0
}

type WantScoreRsp struct {
}

func (m *WantScoreRsp) Reset()                    { *m = WantScoreRsp{} }
func (m *WantScoreRsp) String() string            { return proto.CompactTextString(m) }
func (*WantScoreRsp) ProtoMessage()               {}
func (*WantScoreRsp) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func init() {
	proto.RegisterType((*RegistAccountReq)(nil), "pb.RegistAccountReq")
	proto.RegisterType((*RegistAccountRsp)(nil), "pb.RegistAccountRsp")
	proto.RegisterType((*LoginAccountReq)(nil), "pb.LoginAccountReq")
	proto.RegisterType((*LoginAccountRsp)(nil), "pb.LoginAccountRsp")
	proto.RegisterType((*ResetAccountReq)(nil), "pb.ResetAccountReq")
	proto.RegisterType((*ResetAccountRsp)(nil), "pb.ResetAccountRsp")
	proto.RegisterType((*WantScoreReq)(nil), "pb.WantScoreReq")
	proto.RegisterType((*WantScoreRsp)(nil), "pb.WantScoreRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserServiceExt service

type UserServiceExtClient interface {
	// 注册用户
	RegistAccount(ctx context.Context, in *RegistAccountReq, opts ...client.CallOption) (*RegistAccountRsp, error)
	// 用户登录
	LoginAccount(ctx context.Context, in *LoginAccountReq, opts ...client.CallOption) (*LoginAccountRsp, error)
	// 密码重置
	ResetAccount(ctx context.Context, in *ResetAccountReq, opts ...client.CallOption) (*ResetAccountRsp, error)
	// 评分
	WantScore(ctx context.Context, in *WantScoreReq, opts ...client.CallOption) (*WantScoreRsp, error)
}

type userServiceExtClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceExtClient(serviceName string, c client.Client) UserServiceExtClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "pb"
	}
	return &userServiceExtClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceExtClient) RegistAccount(ctx context.Context, in *RegistAccountReq, opts ...client.CallOption) (*RegistAccountRsp, error) {
	req := c.c.NewRequest(c.serviceName, "UserServiceExt.RegistAccount", in)
	out := new(RegistAccountRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceExtClient) LoginAccount(ctx context.Context, in *LoginAccountReq, opts ...client.CallOption) (*LoginAccountRsp, error) {
	req := c.c.NewRequest(c.serviceName, "UserServiceExt.LoginAccount", in)
	out := new(LoginAccountRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceExtClient) ResetAccount(ctx context.Context, in *ResetAccountReq, opts ...client.CallOption) (*ResetAccountRsp, error) {
	req := c.c.NewRequest(c.serviceName, "UserServiceExt.ResetAccount", in)
	out := new(ResetAccountRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceExtClient) WantScore(ctx context.Context, in *WantScoreReq, opts ...client.CallOption) (*WantScoreRsp, error) {
	req := c.c.NewRequest(c.serviceName, "UserServiceExt.WantScore", in)
	out := new(WantScoreRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserServiceExt service

type UserServiceExtHandler interface {
	// 注册用户
	RegistAccount(context.Context, *RegistAccountReq, *RegistAccountRsp) error
	// 用户登录
	LoginAccount(context.Context, *LoginAccountReq, *LoginAccountRsp) error
	// 密码重置
	ResetAccount(context.Context, *ResetAccountReq, *ResetAccountRsp) error
	// 评分
	WantScore(context.Context, *WantScoreReq, *WantScoreRsp) error
}

func RegisterUserServiceExtHandler(s server.Server, hdlr UserServiceExtHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserServiceExt{hdlr}, opts...))
}

type UserServiceExt struct {
	UserServiceExtHandler
}

func (h *UserServiceExt) RegistAccount(ctx context.Context, in *RegistAccountReq, out *RegistAccountRsp) error {
	return h.UserServiceExtHandler.RegistAccount(ctx, in, out)
}

func (h *UserServiceExt) LoginAccount(ctx context.Context, in *LoginAccountReq, out *LoginAccountRsp) error {
	return h.UserServiceExtHandler.LoginAccount(ctx, in, out)
}

func (h *UserServiceExt) ResetAccount(ctx context.Context, in *ResetAccountReq, out *ResetAccountRsp) error {
	return h.UserServiceExtHandler.ResetAccount(ctx, in, out)
}

func (h *UserServiceExt) WantScore(ctx context.Context, in *WantScoreReq, out *WantScoreRsp) error {
	return h.UserServiceExtHandler.WantScore(ctx, in, out)
}

func init() { proto.RegisterFile("user.ext.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0xfd, 0x92, 0xe8, 0x2b, 0xf4, 0x54, 0x42, 0x31, 0x15, 0x8a, 0x32, 0x21, 0x4f, 0x4c, 0x91,
	0x80, 0x0d, 0x89, 0x81, 0x81, 0x21, 0x12, 0x62, 0x70, 0x05, 0xac, 0x24, 0xe9, 0xa9, 0x8a, 0x44,
	0x62, 0xe3, 0x73, 0x4b, 0x7f, 0x3b, 0x13, 0x72, 0x42, 0x2a, 0xe3, 0x20, 0xc6, 0xf7, 0xee, 0xde,
	0x7b, 0xb9, 0x17, 0x43, 0xbc, 0x21, 0xd4, 0x19, 0xee, 0x4c, 0xa6, 0xb4, 0x34, 0x92, 0x85, 0xaa,
	0xe4, 0xaf, 0x30, 0x17, 0xb8, 0xae, 0xc9, 0xdc, 0x55, 0x95, 0xdc, 0xb4, 0x46, 0xe0, 0x3b, 0x5b,
	0xc0, 0x7f, 0x6c, 0x8a, 0xfa, 0x2d, 0x09, 0xce, 0x83, 0x8b, 0xa9, 0xe8, 0x01, 0x4b, 0xe1, 0xd0,
	0xea, 0x1f, 0x8b, 0x06, 0x93, 0xb0, 0x1b, 0xec, 0xb1, 0x9d, 0xa9, 0x82, 0xe8, 0x43, 0xea, 0x55,
	0x12, 0xf5, 0xb3, 0x01, 0x73, 0xe6, 0x27, 0x90, 0xe2, 0x39, 0x1c, 0x3f, 0xc8, 0x75, 0xdd, 0x3a,
	0xa1, 0xae, 0x7d, 0xf0, 0x87, 0x7d, 0xe8, 0xd9, 0x9f, 0x78, 0x56, 0xa4, 0x2c, 0x25, 0x90, 0xd0,
	0x39, 0x69, 0x44, 0x91, 0xe2, 0xcf, 0x30, 0x7b, 0x29, 0x5a, 0xb3, 0xac, 0xa4, 0x46, 0xfb, 0x01,
	0x67, 0x30, 0xb1, 0x81, 0xf9, 0xaa, 0x8b, 0x8f, 0xc4, 0x37, 0x62, 0x09, 0x1c, 0x34, 0x72, 0x5b,
	0x63, 0xde, 0x67, 0x47, 0x62, 0x80, 0xb6, 0x27, 0xb2, 0xea, 0xee, 0xe4, 0x48, 0xf4, 0x80, 0xc7,
	0xae, 0x2f, 0xa9, 0xab, 0xcf, 0x00, 0xe2, 0x27, 0x42, 0xbd, 0x44, 0xbd, 0xad, 0x2b, 0xbc, 0xdf,
	0x19, 0x76, 0x0b, 0x47, 0x3f, 0x2a, 0x61, 0x8b, 0x4c, 0x95, 0x99, 0xff, 0x1f, 0xd2, 0x5f, 0x58,
	0x52, 0xfc, 0x1f, 0xbb, 0x81, 0x99, 0x7b, 0x32, 0x3b, 0xb5, 0x7b, 0x5e, 0x9f, 0xe9, 0x98, 0x1c,
	0xb4, 0x6e, 0x11, 0xbd, 0xd6, 0x6b, 0x2b, 0x1d, 0x93, 0x9d, 0xf6, 0x12, 0xa6, 0xfb, 0xcb, 0xd8,
	0xdc, 0xee, 0xb8, 0x05, 0xa6, 0x1e, 0x63, 0x25, 0xe5, 0xa4, 0x7b, 0x69, 0xd7, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xbb, 0x47, 0x0f, 0x71, 0x7b, 0x02, 0x00, 0x00,
}