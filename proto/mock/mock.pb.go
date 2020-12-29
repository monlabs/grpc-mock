// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mock.proto

/*
Package grpc_mock_api is a generated protocol buffer package.

It is generated from these files:
	mock.proto

It has these top-level messages:
	AddStubsRequest
	AddStubsResponse
	FindStubsRequest
	FindStubsResponse
	DeleteStubsRequest
	DeleteStubsResponse
	Stub
	Input
	Output
*/
package grpc_mock_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddStubsRequest struct {
	Stubs []*Stub `protobuf:"bytes,1,rep,name=stubs" json:"stubs,omitempty"`
}

func (m *AddStubsRequest) Reset()                    { *m = AddStubsRequest{} }
func (m *AddStubsRequest) String() string            { return proto.CompactTextString(m) }
func (*AddStubsRequest) ProtoMessage()               {}
func (*AddStubsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddStubsRequest) GetStubs() []*Stub {
	if m != nil {
		return m.Stubs
	}
	return nil
}

type AddStubsResponse struct {
}

func (m *AddStubsResponse) Reset()                    { *m = AddStubsResponse{} }
func (m *AddStubsResponse) String() string            { return proto.CompactTextString(m) }
func (*AddStubsResponse) ProtoMessage()               {}
func (*AddStubsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type FindStubsRequest struct {
	Service string                   `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Method  string                   `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
	In      *google_protobuf1.Struct `protobuf:"bytes,3,opt,name=in" json:"in,omitempty"`
}

func (m *FindStubsRequest) Reset()                    { *m = FindStubsRequest{} }
func (m *FindStubsRequest) String() string            { return proto.CompactTextString(m) }
func (*FindStubsRequest) ProtoMessage()               {}
func (*FindStubsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FindStubsRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *FindStubsRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *FindStubsRequest) GetIn() *google_protobuf1.Struct {
	if m != nil {
		return m.In
	}
	return nil
}

type FindStubsResponse struct {
	Stubs []*Stub `protobuf:"bytes,1,rep,name=stubs" json:"stubs,omitempty"`
}

func (m *FindStubsResponse) Reset()                    { *m = FindStubsResponse{} }
func (m *FindStubsResponse) String() string            { return proto.CompactTextString(m) }
func (*FindStubsResponse) ProtoMessage()               {}
func (*FindStubsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FindStubsResponse) GetStubs() []*Stub {
	if m != nil {
		return m.Stubs
	}
	return nil
}

type DeleteStubsRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Method  string `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
}

func (m *DeleteStubsRequest) Reset()                    { *m = DeleteStubsRequest{} }
func (m *DeleteStubsRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteStubsRequest) ProtoMessage()               {}
func (*DeleteStubsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteStubsRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *DeleteStubsRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

type DeleteStubsResponse struct {
}

func (m *DeleteStubsResponse) Reset()                    { *m = DeleteStubsResponse{} }
func (m *DeleteStubsResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteStubsResponse) ProtoMessage()               {}
func (*DeleteStubsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Stub struct {
	Service string  `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Method  string  `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
	In      *Input  `protobuf:"bytes,3,opt,name=in" json:"in,omitempty"`
	Out     *Output `protobuf:"bytes,4,opt,name=out" json:"out,omitempty"`
}

func (m *Stub) Reset()                    { *m = Stub{} }
func (m *Stub) String() string            { return proto.CompactTextString(m) }
func (*Stub) ProtoMessage()               {}
func (*Stub) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Stub) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Stub) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Stub) GetIn() *Input {
	if m != nil {
		return m.In
	}
	return nil
}

func (m *Stub) GetOut() *Output {
	if m != nil {
		return m.Out
	}
	return nil
}

type Input struct {
	// Types that are valid to be assigned to Rule:
	//	*Input_Equals
	//	*Input_Contains
	//	*Input_Matches
	Rule isInput_Rule `protobuf_oneof:"rule"`
}

func (m *Input) Reset()                    { *m = Input{} }
func (m *Input) String() string            { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()               {}
func (*Input) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isInput_Rule interface {
	isInput_Rule()
}

type Input_Equals struct {
	Equals *google_protobuf1.Struct `protobuf:"bytes,1,opt,name=equals,oneof"`
}
type Input_Contains struct {
	Contains *google_protobuf1.Struct `protobuf:"bytes,2,opt,name=contains,oneof"`
}
type Input_Matches struct {
	Matches *google_protobuf1.Struct `protobuf:"bytes,3,opt,name=matches,oneof"`
}

func (*Input_Equals) isInput_Rule()   {}
func (*Input_Contains) isInput_Rule() {}
func (*Input_Matches) isInput_Rule()  {}

func (m *Input) GetRule() isInput_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *Input) GetEquals() *google_protobuf1.Struct {
	if x, ok := m.GetRule().(*Input_Equals); ok {
		return x.Equals
	}
	return nil
}

func (m *Input) GetContains() *google_protobuf1.Struct {
	if x, ok := m.GetRule().(*Input_Contains); ok {
		return x.Contains
	}
	return nil
}

func (m *Input) GetMatches() *google_protobuf1.Struct {
	if x, ok := m.GetRule().(*Input_Matches); ok {
		return x.Matches
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Input) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Input_OneofMarshaler, _Input_OneofUnmarshaler, _Input_OneofSizer, []interface{}{
		(*Input_Equals)(nil),
		(*Input_Contains)(nil),
		(*Input_Matches)(nil),
	}
}

func _Input_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Input)
	// rule
	switch x := m.Rule.(type) {
	case *Input_Equals:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Equals); err != nil {
			return err
		}
	case *Input_Contains:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Contains); err != nil {
			return err
		}
	case *Input_Matches:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Matches); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Input.Rule has unexpected type %T", x)
	}
	return nil
}

func _Input_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Input)
	switch tag {
	case 1: // rule.equals
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Struct)
		err := b.DecodeMessage(msg)
		m.Rule = &Input_Equals{msg}
		return true, err
	case 2: // rule.contains
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Struct)
		err := b.DecodeMessage(msg)
		m.Rule = &Input_Contains{msg}
		return true, err
	case 3: // rule.matches
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Struct)
		err := b.DecodeMessage(msg)
		m.Rule = &Input_Matches{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Input_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Input)
	// rule
	switch x := m.Rule.(type) {
	case *Input_Equals:
		s := proto.Size(x.Equals)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Input_Contains:
		s := proto.Size(x.Contains)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Input_Matches:
		s := proto.Size(x.Matches)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Output struct {
	Data  *google_protobuf1.Struct `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Code  int32                    `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	Error string                   `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *Output) Reset()                    { *m = Output{} }
func (m *Output) String() string            { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()               {}
func (*Output) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Output) GetData() *google_protobuf1.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Output) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Output) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*AddStubsRequest)(nil), "github.com.monlabs.grpc_mock.AddStubsRequest")
	proto.RegisterType((*AddStubsResponse)(nil), "github.com.monlabs.grpc_mock.AddStubsResponse")
	proto.RegisterType((*FindStubsRequest)(nil), "github.com.monlabs.grpc_mock.FindStubsRequest")
	proto.RegisterType((*FindStubsResponse)(nil), "github.com.monlabs.grpc_mock.FindStubsResponse")
	proto.RegisterType((*DeleteStubsRequest)(nil), "github.com.monlabs.grpc_mock.DeleteStubsRequest")
	proto.RegisterType((*DeleteStubsResponse)(nil), "github.com.monlabs.grpc_mock.DeleteStubsResponse")
	proto.RegisterType((*Stub)(nil), "github.com.monlabs.grpc_mock.Stub")
	proto.RegisterType((*Input)(nil), "github.com.monlabs.grpc_mock.Input")
	proto.RegisterType((*Output)(nil), "github.com.monlabs.grpc_mock.Output")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Mock service

type MockClient interface {
	AddStubs(ctx context.Context, in *AddStubsRequest, opts ...grpc.CallOption) (*AddStubsResponse, error)
	FindStubs(ctx context.Context, in *FindStubsRequest, opts ...grpc.CallOption) (*FindStubsResponse, error)
	DeleteStubs(ctx context.Context, in *DeleteStubsRequest, opts ...grpc.CallOption) (*DeleteStubsResponse, error)
}

type mockClient struct {
	cc *grpc.ClientConn
}

func NewMockClient(cc *grpc.ClientConn) MockClient {
	return &mockClient{cc}
}

func (c *mockClient) AddStubs(ctx context.Context, in *AddStubsRequest, opts ...grpc.CallOption) (*AddStubsResponse, error) {
	out := new(AddStubsResponse)
	err := grpc.Invoke(ctx, "/github.com.monlabs.grpc_mock.Mock/AddStubs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mockClient) FindStubs(ctx context.Context, in *FindStubsRequest, opts ...grpc.CallOption) (*FindStubsResponse, error) {
	out := new(FindStubsResponse)
	err := grpc.Invoke(ctx, "/github.com.monlabs.grpc_mock.Mock/FindStubs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mockClient) DeleteStubs(ctx context.Context, in *DeleteStubsRequest, opts ...grpc.CallOption) (*DeleteStubsResponse, error) {
	out := new(DeleteStubsResponse)
	err := grpc.Invoke(ctx, "/github.com.monlabs.grpc_mock.Mock/DeleteStubs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Mock service

type MockServer interface {
	AddStubs(context.Context, *AddStubsRequest) (*AddStubsResponse, error)
	FindStubs(context.Context, *FindStubsRequest) (*FindStubsResponse, error)
	DeleteStubs(context.Context, *DeleteStubsRequest) (*DeleteStubsResponse, error)
}

func RegisterMockServer(s *grpc.Server, srv MockServer) {
	s.RegisterService(&_Mock_serviceDesc, srv)
}

func _Mock_AddStubs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStubsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockServer).AddStubs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.monlabs.grpc_mock.Mock/AddStubs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockServer).AddStubs(ctx, req.(*AddStubsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mock_FindStubs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindStubsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockServer).FindStubs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.monlabs.grpc_mock.Mock/FindStubs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockServer).FindStubs(ctx, req.(*FindStubsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mock_DeleteStubs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStubsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockServer).DeleteStubs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.monlabs.grpc_mock.Mock/DeleteStubs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockServer).DeleteStubs(ctx, req.(*DeleteStubsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.monlabs.grpc_mock.Mock",
	HandlerType: (*MockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddStubs",
			Handler:    _Mock_AddStubs_Handler,
		},
		{
			MethodName: "FindStubs",
			Handler:    _Mock_FindStubs_Handler,
		},
		{
			MethodName: "DeleteStubs",
			Handler:    _Mock_DeleteStubs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mock.proto",
}

func init() { proto.RegisterFile("mock.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xeb, 0xc4, 0x49, 0x9b, 0x89, 0x04, 0xed, 0x10, 0x8a, 0x15, 0xf5, 0x10, 0x19, 0x24,
	0xa2, 0x22, 0xd6, 0x24, 0x11, 0x08, 0xc1, 0x89, 0x0a, 0x55, 0x45, 0xa8, 0x42, 0x72, 0x6f, 0x5c,
	0xa2, 0x8d, 0xbd, 0x24, 0x56, 0xed, 0x5d, 0xd7, 0xbb, 0xdb, 0x6b, 0x25, 0x24, 0xee, 0x48, 0x3c,
	0x03, 0x57, 0x5e, 0x86, 0x57, 0xe0, 0x41, 0x90, 0xd7, 0xce, 0xbf, 0x22, 0x9c, 0x42, 0x6f, 0x1e,
	0xef, 0xf7, 0xcd, 0xfe, 0x66, 0x76, 0x06, 0x20, 0x11, 0xc1, 0x39, 0x49, 0x33, 0xa1, 0x04, 0x1e,
	0x4c, 0x23, 0x35, 0xd3, 0x13, 0x12, 0x88, 0x84, 0x24, 0x82, 0xc7, 0x74, 0x22, 0xc9, 0x34, 0x4b,
	0x83, 0x71, 0xae, 0xe9, 0x1e, 0x4c, 0x85, 0x98, 0xc6, 0xcc, 0xa3, 0x69, 0xe4, 0x51, 0xce, 0x85,
	0xa2, 0x2a, 0x12, 0x5c, 0x16, 0xde, 0xc5, 0xa9, 0x89, 0x26, 0xfa, 0x93, 0x27, 0x55, 0xa6, 0x03,
	0x55, 0x9c, 0xba, 0xef, 0xe1, 0xee, 0x9b, 0x30, 0x3c, 0x53, 0x7a, 0x22, 0x7d, 0x76, 0xa1, 0x99,
	0x54, 0xf8, 0x12, 0x1a, 0x32, 0x8f, 0x1d, 0xab, 0x57, 0xef, 0xb7, 0x87, 0x2e, 0xa9, 0xba, 0x9c,
	0xe4, 0x56, 0xbf, 0x30, 0xb8, 0x08, 0xbb, 0xcb, 0x64, 0x32, 0x15, 0x5c, 0x32, 0x37, 0x81, 0xdd,
	0xe3, 0x88, 0xaf, 0xdf, 0xe0, 0xc0, 0xb6, 0x64, 0xd9, 0x65, 0x14, 0x30, 0xc7, 0xea, 0x59, 0xfd,
	0x96, 0x3f, 0x0f, 0x71, 0x1f, 0x9a, 0x09, 0x53, 0x33, 0x11, 0x3a, 0x35, 0x73, 0x50, 0x46, 0xf8,
	0x18, 0x6a, 0x11, 0x77, 0xea, 0x3d, 0xab, 0xdf, 0x1e, 0x3e, 0x20, 0x45, 0x45, 0x64, 0x5e, 0x11,
	0x39, 0x33, 0x15, 0xf9, 0xb5, 0x88, 0xbb, 0xa7, 0xb0, 0xb7, 0x72, 0x5d, 0xc1, 0x70, 0x8b, 0x8a,
	0x8e, 0x01, 0xdf, 0xb2, 0x98, 0x29, 0x76, 0x3b, 0x7e, 0xf7, 0x3e, 0xdc, 0x5b, 0xcb, 0x53, 0x36,
	0xe7, 0xbb, 0x05, 0x76, 0xfe, 0xe7, 0x3f, 0x3a, 0x32, 0x5a, 0xe9, 0xc8, 0xc3, 0xea, 0x82, 0xde,
	0xf1, 0x54, 0x9b, 0xee, 0xe0, 0x0b, 0xa8, 0x0b, 0xad, 0x1c, 0xdb, 0xb8, 0x1e, 0x55, 0xbb, 0x3e,
	0x68, 0x95, 0xdb, 0x72, 0x83, 0xfb, 0xc3, 0x82, 0x86, 0xc9, 0x82, 0x03, 0x68, 0xb2, 0x0b, 0x4d,
	0x63, 0x69, 0x38, 0xff, 0xfe, 0x18, 0x27, 0x5b, 0x7e, 0x29, 0xc4, 0xe7, 0xb0, 0x13, 0x08, 0xae,
	0x68, 0xc4, 0xa5, 0xa9, 0xa1, 0xd2, 0xb4, 0x90, 0xe2, 0x08, 0xb6, 0x13, 0xaa, 0x82, 0x19, 0x93,
	0x1b, 0xde, 0xfd, 0x64, 0xcb, 0x9f, 0x2b, 0x8f, 0x9a, 0x60, 0x67, 0x3a, 0x66, 0xee, 0x18, 0x9a,
	0x05, 0x3f, 0x3e, 0x01, 0x3b, 0xa4, 0x8a, 0x6e, 0xc0, 0xf5, 0x8d, 0x08, 0x11, 0xec, 0x40, 0x84,
	0xcc, 0x60, 0x36, 0x7c, 0xf3, 0x8d, 0x1d, 0x68, 0xb0, 0x2c, 0x13, 0x99, 0xa1, 0x68, 0xf9, 0x45,
	0x30, 0xfc, 0x5a, 0x07, 0xfb, 0x54, 0x04, 0xe7, 0x78, 0x05, 0x3b, 0xf3, 0x99, 0xc7, 0xa7, 0xd5,
	0x1d, 0xbd, 0xb6, 0x68, 0x5d, 0x72, 0x53, 0x79, 0x39, 0x2d, 0x9d, 0xcf, 0x3f, 0x7f, 0x7d, 0xab,
	0xdd, 0x71, 0x5b, 0xde, 0xe5, 0xc0, 0x33, 0xf3, 0xf9, 0xca, 0x3a, 0xc4, 0x2b, 0x68, 0x2d, 0x26,
	0x1e, 0x37, 0xa4, 0xbc, 0xbe, 0x89, 0x5d, 0xef, 0xc6, 0xfa, 0x92, 0x61, 0xcf, 0x30, 0xb4, 0x71,
	0xc9, 0x80, 0x5f, 0x2c, 0x68, 0xaf, 0x0c, 0x37, 0x3e, 0xab, 0xce, 0xf9, 0xe7, 0x3e, 0x75, 0x07,
	0xff, 0xe0, 0x58, 0xe7, 0x38, 0x5c, 0x72, 0x1c, 0xed, 0x7f, 0xec, 0x10, 0x2f, 0x37, 0xbc, 0x5e,
	0x58, 0xc7, 0x34, 0x8d, 0x26, 0x4d, 0xf3, 0xd6, 0xa3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2c,
	0x37, 0x54, 0x11, 0x51, 0x05, 0x00, 0x00,
}