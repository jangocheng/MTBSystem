// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cms.ext.proto

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

type UserLoginReq struct {
	User     string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *UserLoginReq) Reset()                    { *m = UserLoginReq{} }
func (m *UserLoginReq) String() string            { return proto.CompactTextString(m) }
func (*UserLoginReq) ProtoMessage()               {}
func (*UserLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *UserLoginReq) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserLoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserLoginRsp struct {
	AdminNum   int64  `protobuf:"varint,2,opt,name=adminNum" json:"adminNum,omitempty"`
	CinemaName string `protobuf:"bytes,3,opt,name=cinemaName" json:"cinemaName,omitempty"`
}

func (m *UserLoginRsp) Reset()                    { *m = UserLoginRsp{} }
func (m *UserLoginRsp) String() string            { return proto.CompactTextString(m) }
func (*UserLoginRsp) ProtoMessage()               {}
func (*UserLoginRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *UserLoginRsp) GetAdminNum() int64 {
	if m != nil {
		return m.AdminNum
	}
	return 0
}

func (m *UserLoginRsp) GetCinemaName() string {
	if m != nil {
		return m.CinemaName
	}
	return ""
}

type UpdateMessageReq struct {
	Table string `protobuf:"bytes,1,opt,name=table" json:"table,omitempty"`
	Json  string `protobuf:"bytes,2,opt,name=json" json:"json,omitempty"`
	Num   string `protobuf:"bytes,3,opt,name=num" json:"num,omitempty"`
}

func (m *UpdateMessageReq) Reset()                    { *m = UpdateMessageReq{} }
func (m *UpdateMessageReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateMessageReq) ProtoMessage()               {}
func (*UpdateMessageReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *UpdateMessageReq) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *UpdateMessageReq) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

func (m *UpdateMessageReq) GetNum() string {
	if m != nil {
		return m.Num
	}
	return ""
}

type UpdateMessageRsp struct {
}

func (m *UpdateMessageRsp) Reset()                    { *m = UpdateMessageRsp{} }
func (m *UpdateMessageRsp) String() string            { return proto.CompactTextString(m) }
func (*UpdateMessageRsp) ProtoMessage()               {}
func (*UpdateMessageRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func init() {
	proto.RegisterType((*UserLoginReq)(nil), "pb.UserLoginReq")
	proto.RegisterType((*UserLoginRsp)(nil), "pb.UserLoginRsp")
	proto.RegisterType((*UpdateMessageReq)(nil), "pb.UpdateMessageReq")
	proto.RegisterType((*UpdateMessageRsp)(nil), "pb.UpdateMessageRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CMSServiceExt service

type CMSServiceExtClient interface {
	UserLogin(ctx context.Context, in *UserLoginReq, opts ...client.CallOption) (*UserLoginRsp, error)
	UpdateMessage(ctx context.Context, in *UpdateMessageReq, opts ...client.CallOption) (*UpdateMessageRsp, error)
}

type cMSServiceExtClient struct {
	c           client.Client
	serviceName string
}

func NewCMSServiceExtClient(serviceName string, c client.Client) CMSServiceExtClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "pb"
	}
	return &cMSServiceExtClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *cMSServiceExtClient) UserLogin(ctx context.Context, in *UserLoginReq, opts ...client.CallOption) (*UserLoginRsp, error) {
	req := c.c.NewRequest(c.serviceName, "CMSServiceExt.UserLogin", in)
	out := new(UserLoginRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cMSServiceExtClient) UpdateMessage(ctx context.Context, in *UpdateMessageReq, opts ...client.CallOption) (*UpdateMessageRsp, error) {
	req := c.c.NewRequest(c.serviceName, "CMSServiceExt.UpdateMessage", in)
	out := new(UpdateMessageRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CMSServiceExt service

type CMSServiceExtHandler interface {
	UserLogin(context.Context, *UserLoginReq, *UserLoginRsp) error
	UpdateMessage(context.Context, *UpdateMessageReq, *UpdateMessageRsp) error
}

func RegisterCMSServiceExtHandler(s server.Server, hdlr CMSServiceExtHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&CMSServiceExt{hdlr}, opts...))
}

type CMSServiceExt struct {
	CMSServiceExtHandler
}

func (h *CMSServiceExt) UserLogin(ctx context.Context, in *UserLoginReq, out *UserLoginRsp) error {
	return h.CMSServiceExtHandler.UserLogin(ctx, in, out)
}

func (h *CMSServiceExt) UpdateMessage(ctx context.Context, in *UpdateMessageReq, out *UpdateMessageRsp) error {
	return h.CMSServiceExtHandler.UpdateMessage(ctx, in, out)
}

func init() { proto.RegisterFile("cms.ext.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x4b, 0xc4, 0x40,
	0x10, 0xc5, 0xcd, 0x45, 0xc5, 0x1b, 0x0c, 0x84, 0xe1, 0x8a, 0x90, 0x42, 0x24, 0x95, 0x55, 0x40,
	0xad, 0xb5, 0x11, 0x1b, 0xf1, 0x52, 0xe4, 0xb8, 0x0f, 0xb0, 0x49, 0x86, 0x23, 0xe2, 0xfe, 0x71,
	0x67, 0xa3, 0xd7, 0xfa, 0xcd, 0x65, 0xd7, 0x10, 0x62, 0xb0, 0x7b, 0xef, 0xf1, 0xf6, 0xf1, 0xdb,
	0x81, 0xa4, 0x95, 0x5c, 0xd2, 0xd1, 0x95, 0xc6, 0x6a, 0xa7, 0x71, 0x65, 0x9a, 0xe2, 0x11, 0x2e,
	0xf7, 0x4c, 0xf6, 0x55, 0x1f, 0x7a, 0x55, 0xd3, 0x07, 0x22, 0x9c, 0x0e, 0x4c, 0x36, 0x8b, 0xae,
	0xa3, 0x9b, 0x75, 0x1d, 0x34, 0xe6, 0x70, 0x61, 0x04, 0xf3, 0x97, 0xb6, 0x5d, 0xb6, 0x0a, 0xf9,
	0xe4, 0x8b, 0x97, 0xf9, 0x7b, 0x36, 0xbe, 0x2b, 0x3a, 0xd9, 0xab, 0x6a, 0x90, 0xa1, 0x1b, 0xd7,
	0x93, 0xc7, 0x2b, 0x80, 0xb6, 0x57, 0x24, 0x45, 0x25, 0x24, 0x65, 0x71, 0x58, 0x9a, 0x25, 0x45,
	0x05, 0xe9, 0xde, 0x74, 0xc2, 0xd1, 0x96, 0x98, 0xc5, 0x81, 0x3c, 0xcf, 0x06, 0xce, 0x9c, 0x68,
	0xde, 0x69, 0x04, 0xfa, 0x35, 0x9e, 0xf2, 0x8d, 0xb5, 0x1a, 0x69, 0x82, 0xc6, 0x14, 0x62, 0x35,
	0xc8, 0x71, 0xd6, 0xcb, 0x02, 0x97, 0x7b, 0x6c, 0xee, 0xbe, 0x23, 0x48, 0x9e, 0xb6, 0xbb, 0x1d,
	0xd9, 0xcf, 0xbe, 0xa5, 0xe7, 0xa3, 0xc3, 0x5b, 0x58, 0x4f, 0x3f, 0xc0, 0xb4, 0x34, 0x4d, 0x39,
	0x3f, 0x48, 0xbe, 0x48, 0xd8, 0x14, 0x27, 0xf8, 0x00, 0xc9, 0x9f, 0x61, 0xdc, 0x84, 0xd2, 0x82,
	0x3d, 0xff, 0x27, 0xf5, 0xcf, 0x9b, 0xf3, 0x70, 0xfe, 0xfb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x6f, 0x26, 0x99, 0xb3, 0x8f, 0x01, 0x00, 0x00,
}
