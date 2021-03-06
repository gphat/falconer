// Code generated by protoc-gen-go. DO NOT EDIT.
// source: falconer.proto

/*
Package falconer is a generated protocol buffer package.

It is generated from these files:
	falconer.proto

It has these top-level messages:
	FindSpanRequest
	SpanBatch
	SpanResponse
	TraceRequest
*/
package falconer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ssf "github.com/stripe/veneur/ssf"

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

type FindSpanRequest struct {
	Name string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tags map[string]string `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FindSpanRequest) Reset()                    { *m = FindSpanRequest{} }
func (m *FindSpanRequest) String() string            { return proto.CompactTextString(m) }
func (*FindSpanRequest) ProtoMessage()               {}
func (*FindSpanRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FindSpanRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FindSpanRequest) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type SpanBatch struct {
	Spans []*ssf.SSFSpan `protobuf:"bytes,1,rep,name=spans" json:"spans,omitempty"`
}

func (m *SpanBatch) Reset()                    { *m = SpanBatch{} }
func (m *SpanBatch) String() string            { return proto.CompactTextString(m) }
func (*SpanBatch) ProtoMessage()               {}
func (*SpanBatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SpanBatch) GetSpans() []*ssf.SSFSpan {
	if m != nil {
		return m.Spans
	}
	return nil
}

// This clearly needs something else in it. Likely some backpressure stuff.
type SpanResponse struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *SpanResponse) Reset()                    { *m = SpanResponse{} }
func (m *SpanResponse) String() string            { return proto.CompactTextString(m) }
func (*SpanResponse) ProtoMessage()               {}
func (*SpanResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SpanResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type TraceRequest struct {
	TraceID int64 `protobuf:"varint,1,opt,name=traceID" json:"traceID,omitempty"`
}

func (m *TraceRequest) Reset()                    { *m = TraceRequest{} }
func (m *TraceRequest) String() string            { return proto.CompactTextString(m) }
func (*TraceRequest) ProtoMessage()               {}
func (*TraceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TraceRequest) GetTraceID() int64 {
	if m != nil {
		return m.TraceID
	}
	return 0
}

func init() {
	proto.RegisterType((*FindSpanRequest)(nil), "falconer.FindSpanRequest")
	proto.RegisterType((*SpanBatch)(nil), "falconer.SpanBatch")
	proto.RegisterType((*SpanResponse)(nil), "falconer.SpanResponse")
	proto.RegisterType((*TraceRequest)(nil), "falconer.TraceRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Falconer service

type FalconerClient interface {
	SendSpans(ctx context.Context, opts ...grpc.CallOption) (Falconer_SendSpansClient, error)
	FindSpans(ctx context.Context, in *FindSpanRequest, opts ...grpc.CallOption) (Falconer_FindSpansClient, error)
	GetTrace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (Falconer_GetTraceClient, error)
	WatchSpans(ctx context.Context, in *FindSpanRequest, opts ...grpc.CallOption) (Falconer_WatchSpansClient, error)
}

type falconerClient struct {
	cc *grpc.ClientConn
}

func NewFalconerClient(cc *grpc.ClientConn) FalconerClient {
	return &falconerClient{cc}
}

func (c *falconerClient) SendSpans(ctx context.Context, opts ...grpc.CallOption) (Falconer_SendSpansClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Falconer_serviceDesc.Streams[0], c.cc, "/falconer.Falconer/SendSpans", opts...)
	if err != nil {
		return nil, err
	}
	x := &falconerSendSpansClient{stream}
	return x, nil
}

type Falconer_SendSpansClient interface {
	Send(*SpanBatch) error
	CloseAndRecv() (*SpanResponse, error)
	grpc.ClientStream
}

type falconerSendSpansClient struct {
	grpc.ClientStream
}

func (x *falconerSendSpansClient) Send(m *SpanBatch) error {
	return x.ClientStream.SendMsg(m)
}

func (x *falconerSendSpansClient) CloseAndRecv() (*SpanResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SpanResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *falconerClient) FindSpans(ctx context.Context, in *FindSpanRequest, opts ...grpc.CallOption) (Falconer_FindSpansClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Falconer_serviceDesc.Streams[1], c.cc, "/falconer.Falconer/FindSpans", opts...)
	if err != nil {
		return nil, err
	}
	x := &falconerFindSpansClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Falconer_FindSpansClient interface {
	Recv() (*ssf.SSFSpan, error)
	grpc.ClientStream
}

type falconerFindSpansClient struct {
	grpc.ClientStream
}

func (x *falconerFindSpansClient) Recv() (*ssf.SSFSpan, error) {
	m := new(ssf.SSFSpan)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *falconerClient) GetTrace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (Falconer_GetTraceClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Falconer_serviceDesc.Streams[2], c.cc, "/falconer.Falconer/GetTrace", opts...)
	if err != nil {
		return nil, err
	}
	x := &falconerGetTraceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Falconer_GetTraceClient interface {
	Recv() (*ssf.SSFSpan, error)
	grpc.ClientStream
}

type falconerGetTraceClient struct {
	grpc.ClientStream
}

func (x *falconerGetTraceClient) Recv() (*ssf.SSFSpan, error) {
	m := new(ssf.SSFSpan)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *falconerClient) WatchSpans(ctx context.Context, in *FindSpanRequest, opts ...grpc.CallOption) (Falconer_WatchSpansClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Falconer_serviceDesc.Streams[3], c.cc, "/falconer.Falconer/WatchSpans", opts...)
	if err != nil {
		return nil, err
	}
	x := &falconerWatchSpansClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Falconer_WatchSpansClient interface {
	Recv() (*ssf.SSFSpan, error)
	grpc.ClientStream
}

type falconerWatchSpansClient struct {
	grpc.ClientStream
}

func (x *falconerWatchSpansClient) Recv() (*ssf.SSFSpan, error) {
	m := new(ssf.SSFSpan)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Falconer service

type FalconerServer interface {
	SendSpans(Falconer_SendSpansServer) error
	FindSpans(*FindSpanRequest, Falconer_FindSpansServer) error
	GetTrace(*TraceRequest, Falconer_GetTraceServer) error
	WatchSpans(*FindSpanRequest, Falconer_WatchSpansServer) error
}

func RegisterFalconerServer(s *grpc.Server, srv FalconerServer) {
	s.RegisterService(&_Falconer_serviceDesc, srv)
}

func _Falconer_SendSpans_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FalconerServer).SendSpans(&falconerSendSpansServer{stream})
}

type Falconer_SendSpansServer interface {
	SendAndClose(*SpanResponse) error
	Recv() (*SpanBatch, error)
	grpc.ServerStream
}

type falconerSendSpansServer struct {
	grpc.ServerStream
}

func (x *falconerSendSpansServer) SendAndClose(m *SpanResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *falconerSendSpansServer) Recv() (*SpanBatch, error) {
	m := new(SpanBatch)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Falconer_FindSpans_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindSpanRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FalconerServer).FindSpans(m, &falconerFindSpansServer{stream})
}

type Falconer_FindSpansServer interface {
	Send(*ssf.SSFSpan) error
	grpc.ServerStream
}

type falconerFindSpansServer struct {
	grpc.ServerStream
}

func (x *falconerFindSpansServer) Send(m *ssf.SSFSpan) error {
	return x.ServerStream.SendMsg(m)
}

func _Falconer_GetTrace_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TraceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FalconerServer).GetTrace(m, &falconerGetTraceServer{stream})
}

type Falconer_GetTraceServer interface {
	Send(*ssf.SSFSpan) error
	grpc.ServerStream
}

type falconerGetTraceServer struct {
	grpc.ServerStream
}

func (x *falconerGetTraceServer) Send(m *ssf.SSFSpan) error {
	return x.ServerStream.SendMsg(m)
}

func _Falconer_WatchSpans_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindSpanRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FalconerServer).WatchSpans(m, &falconerWatchSpansServer{stream})
}

type Falconer_WatchSpansServer interface {
	Send(*ssf.SSFSpan) error
	grpc.ServerStream
}

type falconerWatchSpansServer struct {
	grpc.ServerStream
}

func (x *falconerWatchSpansServer) Send(m *ssf.SSFSpan) error {
	return x.ServerStream.SendMsg(m)
}

var _Falconer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "falconer.Falconer",
	HandlerType: (*FalconerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendSpans",
			Handler:       _Falconer_SendSpans_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindSpans",
			Handler:       _Falconer_FindSpans_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetTrace",
			Handler:       _Falconer_GetTrace_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchSpans",
			Handler:       _Falconer_WatchSpans_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "falconer.proto",
}

func init() { proto.RegisterFile("falconer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4e, 0xa3, 0x50,
	0x14, 0xee, 0xed, 0xcf, 0x0c, 0x9c, 0x69, 0x66, 0x26, 0x77, 0x26, 0x06, 0x59, 0x35, 0xb8, 0x41,
	0x17, 0x60, 0xaa, 0x49, 0xab, 0x71, 0x65, 0xb4, 0xc6, 0x2d, 0x34, 0x71, 0x7d, 0x8b, 0xa7, 0x94,
	0xd8, 0x5e, 0x90, 0x73, 0x69, 0xd2, 0x27, 0xf1, 0x41, 0x7d, 0x01, 0xc3, 0xa5, 0x60, 0x6b, 0x74,
	0xe1, 0xee, 0x7c, 0x07, 0xbe, 0xf3, 0xfd, 0xe4, 0xc2, 0xef, 0xb9, 0x58, 0x46, 0xa9, 0xc4, 0xdc,
	0xcb, 0xf2, 0x54, 0xa5, 0xdc, 0xa8, 0xb1, 0x7d, 0x1c, 0x27, 0x6a, 0x51, 0xcc, 0xbc, 0x28, 0x5d,
	0xf9, 0xa4, 0xf2, 0x24, 0x43, 0x7f, 0x8d, 0x12, 0x8b, 0xdc, 0x27, 0x9a, 0xfb, 0x24, 0x56, 0xd9,
	0x12, 0x2b, 0x92, 0xf3, 0xc2, 0xe0, 0xcf, 0x24, 0x91, 0x8f, 0x61, 0x26, 0x64, 0x80, 0xcf, 0x05,
	0x92, 0xe2, 0x1c, 0xba, 0x52, 0xac, 0xd0, 0x62, 0x03, 0xe6, 0x9a, 0x81, 0x9e, 0xf9, 0x08, 0xba,
	0x4a, 0xc4, 0x64, 0xb5, 0x07, 0x1d, 0xf7, 0xd7, 0xf0, 0xc8, 0x6b, 0xb4, 0x3f, 0x90, 0xbd, 0xa9,
	0x88, 0xe9, 0x56, 0xaa, 0x7c, 0x13, 0x68, 0x82, 0x3d, 0x02, 0xb3, 0x59, 0xf1, 0xbf, 0xd0, 0x79,
	0xc2, 0xcd, 0xf6, 0x70, 0x39, 0xf2, 0xff, 0xd0, 0x5b, 0x8b, 0x65, 0x81, 0x56, 0x5b, 0xef, 0x2a,
	0x70, 0xd9, 0x1e, 0x33, 0xc7, 0x07, 0xb3, 0xbc, 0x7b, 0x2d, 0x54, 0xb4, 0xe0, 0x0e, 0xf4, 0x28,
	0x13, 0x92, 0x2c, 0xa6, 0xf5, 0xfb, 0x1e, 0xd1, 0xdc, 0x0b, 0xc3, 0x89, 0x56, 0xae, 0x3e, 0x39,
	0x27, 0xd0, 0xaf, 0x8c, 0x50, 0x96, 0x4a, 0x42, 0x6e, 0x83, 0x11, 0xe7, 0x88, 0x2a, 0x91, 0xf1,
	0x56, 0xb1, 0xc1, 0x8e, 0x0b, 0xfd, 0x69, 0x2e, 0x22, 0xac, 0x23, 0x5b, 0xf0, 0x53, 0x95, 0xf8,
	0xfe, 0x46, 0xff, 0xda, 0x09, 0x6a, 0x38, 0x7c, 0x65, 0x60, 0x4c, 0xb6, 0x61, 0xf9, 0x15, 0x98,
	0x21, 0x56, 0x79, 0x89, 0xff, 0x7b, 0x2f, 0xa1, 0x31, 0x6a, 0x1f, 0xec, 0x2f, 0x6b, 0x33, 0x4e,
	0xcb, 0x65, 0x7c, 0x0c, 0x66, 0xdd, 0x16, 0xf1, 0xc3, 0x2f, 0x2b, 0xb4, 0xf7, 0xd2, 0x39, 0xad,
	0x53, 0xc6, 0xcf, 0xc1, 0xb8, 0x43, 0xa5, 0x1d, 0xf3, 0x1d, 0x85, 0xdd, 0x08, 0x9f, 0xb0, 0x2e,
	0x00, 0x1e, 0x4a, 0x53, 0xdf, 0x17, 0x9c, 0xfd, 0xd0, 0xaf, 0xe3, 0xec, 0x2d, 0x00, 0x00, 0xff,
	0xff, 0x70, 0x25, 0x83, 0x2c, 0x64, 0x02, 0x00, 0x00,
}
