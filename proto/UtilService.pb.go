// Code generated by protoc-gen-go. DO NOT EDIT.
// source: UtilService.proto

package UtilService

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type DeepLinkResp struct {
	DeepLink             string   `protobuf:"bytes,1,opt,name=deepLink,proto3" json:"deepLink,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *DeepLinkResp) Reset()         { *m = DeepLinkResp{} }
func (m *DeepLinkResp) String() string { return proto.CompactTextString(m) }
func (*DeepLinkResp) ProtoMessage()    {}
func (*DeepLinkResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b77ec18d394a137, []int{0}
}

func (m *DeepLinkResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeepLinkResp.Unmarshal(m, b)
}
func (m *DeepLinkResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeepLinkResp.Marshal(b, m, deterministic)
}
func (m *DeepLinkResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeepLinkResp.Merge(m, src)
}
func (m *DeepLinkResp) XXX_Size() int {
	return xxx_messageInfo_DeepLinkResp.Size(m)
}
func (m *DeepLinkResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeepLinkResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeepLinkResp proto.InternalMessageInfo

func (m *DeepLinkResp) GetDeepLink() string {
	if m != nil {
		return m.DeepLink
	}
	return ""
}

type DeepLinkReq struct {
	Board                string   `protobuf:"bytes,1,opt,name=board,proto3" json:"board,omitempty"`
	PackageName          string   `protobuf:"bytes,2,opt,name=packageName,proto3" json:"packageName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *DeepLinkReq) Reset()         { *m = DeepLinkReq{} }
func (m *DeepLinkReq) String() string { return proto.CompactTextString(m) }
func (*DeepLinkReq) ProtoMessage()    {}
func (*DeepLinkReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b77ec18d394a137, []int{1}
}

func (m *DeepLinkReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeepLinkReq.Unmarshal(m, b)
}
func (m *DeepLinkReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeepLinkReq.Marshal(b, m, deterministic)
}
func (m *DeepLinkReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeepLinkReq.Merge(m, src)
}
func (m *DeepLinkReq) XXX_Size() int {
	return xxx_messageInfo_DeepLinkReq.Size(m)
}
func (m *DeepLinkReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeepLinkReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeepLinkReq proto.InternalMessageInfo

func (m *DeepLinkReq) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

func (m *DeepLinkReq) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

type DeadLinkTarget struct {
	Source               string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *DeadLinkTarget) Reset()         { *m = DeadLinkTarget{} }
func (m *DeadLinkTarget) String() string { return proto.CompactTextString(m) }
func (*DeadLinkTarget) ProtoMessage()    {}
func (*DeadLinkTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b77ec18d394a137, []int{2}
}

func (m *DeadLinkTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeadLinkTarget.Unmarshal(m, b)
}
func (m *DeadLinkTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeadLinkTarget.Marshal(b, m, deterministic)
}
func (m *DeadLinkTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeadLinkTarget.Merge(m, src)
}
func (m *DeadLinkTarget) XXX_Size() int {
	return xxx_messageInfo_DeadLinkTarget.Size(m)
}
func (m *DeadLinkTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_DeadLinkTarget.DiscardUnknown(m)
}

var xxx_messageInfo_DeadLinkTarget proto.InternalMessageInfo

func (m *DeadLinkTarget) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type DeadLinkResp struct {
	IsSuccessful         bool     `protobuf:"varint,1,opt,name=isSuccessful,proto3" json:"isSuccessful,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *DeadLinkResp) Reset()         { *m = DeadLinkResp{} }
func (m *DeadLinkResp) String() string { return proto.CompactTextString(m) }
func (*DeadLinkResp) ProtoMessage()    {}
func (*DeadLinkResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b77ec18d394a137, []int{3}
}

func (m *DeadLinkResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeadLinkResp.Unmarshal(m, b)
}
func (m *DeadLinkResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeadLinkResp.Marshal(b, m, deterministic)
}
func (m *DeadLinkResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeadLinkResp.Merge(m, src)
}
func (m *DeadLinkResp) XXX_Size() int {
	return xxx_messageInfo_DeadLinkResp.Size(m)
}
func (m *DeadLinkResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeadLinkResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeadLinkResp proto.InternalMessageInfo

func (m *DeadLinkResp) GetIsSuccessful() bool {
	if m != nil {
		return m.IsSuccessful
	}
	return false
}

type AppDeepLink struct {
	PackageName          string    `protobuf:"bytes,1,opt,name=packageName,proto3" json:"packageName,omitempty"`
	Traget               []*Target `protobuf:"bytes,2,rep,name=traget,proto3" json:"traget,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" bson:"-"`
	XXX_unrecognized     []byte    `json:"-" bson:"-"`
	XXX_sizecache        int32     `json:"-" bson:"-"`
}

func (m *AppDeepLink) Reset()         { *m = AppDeepLink{} }
func (m *AppDeepLink) String() string { return proto.CompactTextString(m) }
func (*AppDeepLink) ProtoMessage()    {}
func (*AppDeepLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b77ec18d394a137, []int{4}
}

func (m *AppDeepLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppDeepLink.Unmarshal(m, b)
}
func (m *AppDeepLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppDeepLink.Marshal(b, m, deterministic)
}
func (m *AppDeepLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppDeepLink.Merge(m, src)
}
func (m *AppDeepLink) XXX_Size() int {
	return xxx_messageInfo_AppDeepLink.Size(m)
}
func (m *AppDeepLink) XXX_DiscardUnknown() {
	xxx_messageInfo_AppDeepLink.DiscardUnknown(m)
}

var xxx_messageInfo_AppDeepLink proto.InternalMessageInfo

func (m *AppDeepLink) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *AppDeepLink) GetTraget() []*Target {
	if m != nil {
		return m.Traget
	}
	return nil
}

type Target struct {
	Board                []string `protobuf:"bytes,1,rep,name=board,proto3" json:"board,omitempty"`
	DeepLink             string   `protobuf:"bytes,2,opt,name=deepLink,proto3" json:"deepLink,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Target) Reset()         { *m = Target{} }
func (m *Target) String() string { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()    {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b77ec18d394a137, []int{5}
}

func (m *Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Target.Unmarshal(m, b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Target.Marshal(b, m, deterministic)
}
func (m *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(m, src)
}
func (m *Target) XXX_Size() int {
	return xxx_messageInfo_Target.Size(m)
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

func (m *Target) GetBoard() []string {
	if m != nil {
		return m.Board
	}
	return nil
}

func (m *Target) GetDeepLink() string {
	if m != nil {
		return m.DeepLink
	}
	return ""
}

func init() {
	proto.RegisterType((*DeepLinkResp)(nil), "UtilService.DeepLinkResp")
	proto.RegisterType((*DeepLinkReq)(nil), "UtilService.DeepLinkReq")
	proto.RegisterType((*DeadLinkTarget)(nil), "UtilService.DeadLinkTarget")
	proto.RegisterType((*DeadLinkResp)(nil), "UtilService.DeadLinkResp")
	proto.RegisterType((*AppDeepLink)(nil), "UtilService.AppDeepLink")
	proto.RegisterType((*Target)(nil), "UtilService.Target")
}

func init() { proto.RegisterFile("UtilService.proto", fileDescriptor_1b77ec18d394a137) }

var fileDescriptor_1b77ec18d394a137 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x4e, 0x02, 0x31,
	0x14, 0x86, 0x33, 0x43, 0x9c, 0xe0, 0x1b, 0xa2, 0xa1, 0x1a, 0x1d, 0xd1, 0x05, 0xe9, 0x8a, 0x60,
	0x02, 0x09, 0xee, 0xd8, 0x19, 0x31, 0x6e, 0x8c, 0x8b, 0x41, 0x57, 0xba, 0x29, 0x33, 0xcf, 0xb1,
	0x01, 0xa7, 0xb5, 0x2d, 0x1e, 0xc0, 0x2b, 0x78, 0x2f, 0x37, 0x5e, 0xc1, 0x83, 0x18, 0x66, 0x0a,
	0x74, 0x40, 0x97, 0xff, 0xeb, 0x9f, 0xbf, 0xdf, 0xfb, 0x5b, 0x68, 0x3e, 0x18, 0x3e, 0x1b, 0xa3,
	0x7a, 0xe7, 0x09, 0xf6, 0xa4, 0x12, 0x46, 0x90, 0xd0, 0x19, 0xb5, 0xce, 0x32, 0x21, 0xb2, 0x19,
	0xf6, 0x99, 0xe4, 0x7d, 0x96, 0xe7, 0xc2, 0x30, 0xc3, 0x45, 0xae, 0x4b, 0x2b, 0xed, 0x42, 0x63,
	0x84, 0x28, 0x6f, 0x79, 0x3e, 0x8d, 0x51, 0x4b, 0xd2, 0x82, 0x7a, 0x6a, 0x75, 0xe4, 0xb5, 0xbd,
	0xce, 0x6e, 0xbc, 0xd2, 0xf4, 0x1a, 0xc2, 0xb5, 0xf7, 0x8d, 0x1c, 0xc2, 0xce, 0x44, 0x30, 0x95,
	0x5a, 0x5f, 0x29, 0x48, 0x1b, 0x42, 0xc9, 0x92, 0x29, 0xcb, 0xf0, 0x8e, 0xbd, 0x62, 0xe4, 0x17,
	0x67, 0xee, 0x88, 0x76, 0x60, 0x6f, 0x84, 0x2c, 0x5d, 0xc4, 0xdc, 0x33, 0x95, 0xa1, 0x21, 0x47,
	0x10, 0x68, 0x31, 0x57, 0x09, 0xda, 0x28, 0xab, 0xe8, 0x60, 0x01, 0x57, 0x3a, 0x0b, 0x38, 0x0a,
	0x0d, 0xae, 0xc7, 0xf3, 0x24, 0x41, 0xad, 0x9f, 0xe7, 0xb3, 0xc2, 0x5d, 0x8f, 0x2b, 0x33, 0xfa,
	0x04, 0xe1, 0xa5, 0x94, 0x4b, 0xce, 0x4d, 0x1c, 0x6f, 0x0b, 0x87, 0x9c, 0x43, 0x60, 0x14, 0xcb,
	0xd0, 0x44, 0x7e, 0xbb, 0xd6, 0x09, 0x07, 0x07, 0x3d, 0xb7, 0xd0, 0x92, 0x30, 0xb6, 0x16, 0x3a,
	0x84, 0xc0, 0x32, 0x3b, 0xdb, 0xd7, 0xd6, 0xdb, 0xbb, 0xf5, 0xf9, 0xd5, 0xfa, 0x06, 0x5f, 0x1e,
	0xb8, 0x0f, 0x43, 0x26, 0xb0, 0xbf, 0xdc, 0xee, 0xea, 0x05, 0x93, 0x29, 0x2a, 0x72, 0x5a, 0xb9,
	0xbb, 0xda, 0x52, 0xeb, 0xe4, 0xcf, 0xc3, 0x45, 0x31, 0xf4, 0xf8, 0xe3, 0xfb, 0xe7, 0xd3, 0x6f,
	0xd2, 0x46, 0x3f, 0x45, 0x96, 0xda, 0xb4, 0xa1, 0xd7, 0x25, 0x8f, 0x10, 0xde, 0xa0, 0x59, 0xb5,
	0x11, 0x6d, 0x44, 0xac, 0x1e, 0x73, 0x2b, 0x7c, 0xfd, 0x25, 0xfe, 0x0d, 0x9f, 0x04, 0xc5, 0x17,
	0xba, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x6a, 0x7f, 0x5f, 0x82, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UtilServiceClient is the client API for UtilService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UtilServiceClient interface {
	//Schedular CRUD
	DeadLinkChecker(ctx context.Context, in *DeadLinkTarget, opts ...grpc.CallOption) (*DeadLinkResp, error)
	GetDeepLink(ctx context.Context, in *DeepLinkReq, opts ...grpc.CallOption) (*DeepLinkResp, error)
}

type utilServiceClient struct {
	cc *grpc.ClientConn
}

func NewUtilServiceClient(cc *grpc.ClientConn) UtilServiceClient {
	return &utilServiceClient{cc}
}

func (c *utilServiceClient) DeadLinkChecker(ctx context.Context, in *DeadLinkTarget, opts ...grpc.CallOption) (*DeadLinkResp, error) {
	out := new(DeadLinkResp)
	err := c.cc.Invoke(ctx, "/UtilService.UtilService/DeadLinkChecker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *utilServiceClient) GetDeepLink(ctx context.Context, in *DeepLinkReq, opts ...grpc.CallOption) (*DeepLinkResp, error) {
	out := new(DeepLinkResp)
	err := c.cc.Invoke(ctx, "/UtilService.UtilService/GetDeepLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UtilServiceServer is the server API for UtilService service.
type UtilServiceServer interface {
	//Schedular CRUD
	DeadLinkChecker(context.Context, *DeadLinkTarget) (*DeadLinkResp, error)
	GetDeepLink(context.Context, *DeepLinkReq) (*DeepLinkResp, error)
}

// UnimplementedUtilServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUtilServiceServer struct {
}

func (*UnimplementedUtilServiceServer) DeadLinkChecker(ctx context.Context, req *DeadLinkTarget) (*DeadLinkResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeadLinkChecker not implemented")
}
func (*UnimplementedUtilServiceServer) GetDeepLink(ctx context.Context, req *DeepLinkReq) (*DeepLinkResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeepLink not implemented")
}

func RegisterUtilServiceServer(s *grpc.Server, srv UtilServiceServer) {
	s.RegisterService(&_UtilService_serviceDesc, srv)
}

func _UtilService_DeadLinkChecker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeadLinkTarget)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilServiceServer).DeadLinkChecker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UtilService.UtilService/DeadLinkChecker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilServiceServer).DeadLinkChecker(ctx, req.(*DeadLinkTarget))
	}
	return interceptor(ctx, in, info, handler)
}

func _UtilService_GetDeepLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeepLinkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilServiceServer).GetDeepLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UtilService.UtilService/GetDeepLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilServiceServer).GetDeepLink(ctx, req.(*DeepLinkReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UtilService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UtilService.UtilService",
	HandlerType: (*UtilServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeadLinkChecker",
			Handler:    _UtilService_DeadLinkChecker_Handler,
		},
		{
			MethodName: "GetDeepLink",
			Handler:    _UtilService_GetDeepLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "UtilService.proto",
}
