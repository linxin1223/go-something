// Code generated by protoc-gen-go. DO NOT EDIT.
// source: NetConfig.proto

package NetConfig

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type CfgMessage struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Filename             string   `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Config               []byte   `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	Websiteid            string   `protobuf:"bytes,4,opt,name=websiteid,proto3" json:"websiteid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CfgMessage) Reset()         { *m = CfgMessage{} }
func (m *CfgMessage) String() string { return proto.CompactTextString(m) }
func (*CfgMessage) ProtoMessage()    {}
func (*CfgMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd2839f1f124fe5f, []int{0}
}

func (m *CfgMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CfgMessage.Unmarshal(m, b)
}
func (m *CfgMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CfgMessage.Marshal(b, m, deterministic)
}
func (m *CfgMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CfgMessage.Merge(m, src)
}
func (m *CfgMessage) XXX_Size() int {
	return xxx_messageInfo_CfgMessage.Size(m)
}
func (m *CfgMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CfgMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CfgMessage proto.InternalMessageInfo

func (m *CfgMessage) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *CfgMessage) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *CfgMessage) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *CfgMessage) GetWebsiteid() string {
	if m != nil {
		return m.Websiteid
	}
	return ""
}

type NetConfigStatus struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetConfigStatus) Reset()         { *m = NetConfigStatus{} }
func (m *NetConfigStatus) String() string { return proto.CompactTextString(m) }
func (*NetConfigStatus) ProtoMessage()    {}
func (*NetConfigStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd2839f1f124fe5f, []int{1}
}

func (m *NetConfigStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetConfigStatus.Unmarshal(m, b)
}
func (m *NetConfigStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetConfigStatus.Marshal(b, m, deterministic)
}
func (m *NetConfigStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetConfigStatus.Merge(m, src)
}
func (m *NetConfigStatus) XXX_Size() int {
	return xxx_messageInfo_NetConfigStatus.Size(m)
}
func (m *NetConfigStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_NetConfigStatus.DiscardUnknown(m)
}

var xxx_messageInfo_NetConfigStatus proto.InternalMessageInfo

func (m *NetConfigStatus) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *NetConfigStatus) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*CfgMessage)(nil), "CfgMessage")
	proto.RegisterType((*NetConfigStatus)(nil), "NetConfigStatus")
}

func init() {
	proto.RegisterFile("NetConfig.proto", fileDescriptor_bd2839f1f124fe5f)
}

var fileDescriptor_bd2839f1f124fe5f = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xf7, 0x4b, 0x2d, 0x71,
	0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x2a, 0xe2, 0xe2, 0x72,
	0x4e, 0x4b, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe2, 0x62, 0x29, 0xa9, 0x2c,
	0x48, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0xb3, 0x85, 0xa4, 0xb8, 0x38, 0xd2, 0x32,
	0x73, 0x52, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x21,
	0x31, 0x2e, 0xb6, 0x64, 0xb0, 0x69, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x50, 0x9e, 0x90,
	0x0c, 0x17, 0x67, 0x79, 0x6a, 0x52, 0x71, 0x66, 0x49, 0x6a, 0x66, 0x8a, 0x04, 0x0b, 0x58, 0x13,
	0x42, 0x40, 0xc9, 0x19, 0xc9, 0x19, 0xc1, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x20, 0x83, 0x8a, 0xc1,
	0x2c, 0xa8, 0xd5, 0x50, 0x9e, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x7e, 0x5e, 0x49, 0x6a, 0x5e, 0x09,
	0xd8, 0x6e, 0x9e, 0x20, 0x18, 0xd7, 0x68, 0x09, 0x23, 0x17, 0x2f, 0xdc, 0x94, 0x90, 0xa2, 0xc4,
	0x3c, 0x21, 0x03, 0x2e, 0xde, 0x80, 0xd2, 0xe2, 0x0c, 0xb8, 0xa0, 0x10, 0xb7, 0x1e, 0xc2, 0x6b,
	0x52, 0x02, 0x7a, 0x68, 0x76, 0x2a, 0x31, 0x08, 0x69, 0x83, 0x74, 0xe4, 0xe4, 0xe0, 0xd0, 0x81,
	0xcc, 0x51, 0x62, 0x10, 0x32, 0xe1, 0x12, 0x04, 0x19, 0x1f, 0x90, 0x58, 0x54, 0x42, 0xbc, 0x15,
	0x49, 0x6c, 0xe0, 0x60, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x25, 0xc5, 0x6b, 0x4d, 0x79,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NetConfigTranClient is the client API for NetConfigTran service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetConfigTranClient interface {
	PushNetConfig(ctx context.Context, in *CfgMessage, opts ...grpc.CallOption) (*NetConfigStatus, error)
	PullNetConfig(ctx context.Context, in *CfgMessage, opts ...grpc.CallOption) (*CfgMessage, error)
	PushPartNetConfig(ctx context.Context, in *CfgMessage, opts ...grpc.CallOption) (*NetConfigStatus, error)
}

type netConfigTranClient struct {
	cc grpc.ClientConnInterface
}

func NewNetConfigTranClient(cc grpc.ClientConnInterface) NetConfigTranClient {
	return &netConfigTranClient{cc}
}

func (c *netConfigTranClient) PushNetConfig(ctx context.Context, in *CfgMessage, opts ...grpc.CallOption) (*NetConfigStatus, error) {
	out := new(NetConfigStatus)
	err := c.cc.Invoke(ctx, "/NetConfigTran/PushNetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netConfigTranClient) PullNetConfig(ctx context.Context, in *CfgMessage, opts ...grpc.CallOption) (*CfgMessage, error) {
	out := new(CfgMessage)
	err := c.cc.Invoke(ctx, "/NetConfigTran/PullNetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netConfigTranClient) PushPartNetConfig(ctx context.Context, in *CfgMessage, opts ...grpc.CallOption) (*NetConfigStatus, error) {
	out := new(NetConfigStatus)
	err := c.cc.Invoke(ctx, "/NetConfigTran/PushPartNetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetConfigTranServer is the server API for NetConfigTran service.
type NetConfigTranServer interface {
	PushNetConfig(context.Context, *CfgMessage) (*NetConfigStatus, error)
	PullNetConfig(context.Context, *CfgMessage) (*CfgMessage, error)
	PushPartNetConfig(context.Context, *CfgMessage) (*NetConfigStatus, error)
}

// UnimplementedNetConfigTranServer can be embedded to have forward compatible implementations.
type UnimplementedNetConfigTranServer struct {
}

func (*UnimplementedNetConfigTranServer) PushNetConfig(ctx context.Context, req *CfgMessage) (*NetConfigStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushNetConfig not implemented")
}
func (*UnimplementedNetConfigTranServer) PullNetConfig(ctx context.Context, req *CfgMessage) (*CfgMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullNetConfig not implemented")
}
func (*UnimplementedNetConfigTranServer) PushPartNetConfig(ctx context.Context, req *CfgMessage) (*NetConfigStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushPartNetConfig not implemented")
}

func RegisterNetConfigTranServer(s *grpc.Server, srv NetConfigTranServer) {
	s.RegisterService(&_NetConfigTran_serviceDesc, srv)
}

func _NetConfigTran_PushNetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CfgMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetConfigTranServer).PushNetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NetConfigTran/PushNetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetConfigTranServer).PushNetConfig(ctx, req.(*CfgMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetConfigTran_PullNetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CfgMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetConfigTranServer).PullNetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NetConfigTran/PullNetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetConfigTranServer).PullNetConfig(ctx, req.(*CfgMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetConfigTran_PushPartNetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CfgMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetConfigTranServer).PushPartNetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NetConfigTran/PushPartNetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetConfigTranServer).PushPartNetConfig(ctx, req.(*CfgMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetConfigTran_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NetConfigTran",
	HandlerType: (*NetConfigTranServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushNetConfig",
			Handler:    _NetConfigTran_PushNetConfig_Handler,
		},
		{
			MethodName: "PullNetConfig",
			Handler:    _NetConfigTran_PullNetConfig_Handler,
		},
		{
			MethodName: "PushPartNetConfig",
			Handler:    _NetConfigTran_PushPartNetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "NetConfig.proto",
}
