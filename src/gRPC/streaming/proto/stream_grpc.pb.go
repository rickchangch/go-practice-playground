// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamServiceClient interface {
	// Simple RPC
	GetFeature(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamResponse, error)
	// Server-side streaming RPC
	ListFeatures(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (StreamService_ListFeaturesClient, error)
	// Client-side streaming RPC
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (StreamService_RecordRouteClient, error)
	// Bidirectional streaming RPC
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (StreamService_RouteChatClient, error)
}

type streamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamServiceClient(cc grpc.ClientConnInterface) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) GetFeature(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamResponse, error) {
	out := new(StreamResponse)
	err := c.cc.Invoke(ctx, "/streamingpb.StreamService/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamServiceClient) ListFeatures(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (StreamService_ListFeaturesClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[0], "/streamingpb.StreamService/ListFeatures", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceListFeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_ListFeaturesClient interface {
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServiceListFeaturesClient struct {
	grpc.ClientStream
}

func (x *streamServiceListFeaturesClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (StreamService_RecordRouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[1], "/streamingpb.StreamService/RecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceRecordRouteClient{stream}
	return x, nil
}

type StreamService_RecordRouteClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServiceRecordRouteClient struct {
	grpc.ClientStream
}

func (x *streamServiceRecordRouteClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceRecordRouteClient) CloseAndRecv() (*StreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (StreamService_RouteChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[2], "/streamingpb.StreamService/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceRouteChatClient{stream}
	return x, nil
}

type StreamService_RouteChatClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServiceRouteChatClient struct {
	grpc.ClientStream
}

func (x *streamServiceRouteChatClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceRouteChatClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
// All implementations must embed UnimplementedStreamServiceServer
// for forward compatibility
type StreamServiceServer interface {
	// Simple RPC
	GetFeature(context.Context, *StreamRequest) (*StreamResponse, error)
	// Server-side streaming RPC
	ListFeatures(*StreamRequest, StreamService_ListFeaturesServer) error
	// Client-side streaming RPC
	RecordRoute(StreamService_RecordRouteServer) error
	// Bidirectional streaming RPC
	RouteChat(StreamService_RouteChatServer) error
	mustEmbedUnimplementedStreamServiceServer()
}

// UnimplementedStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (UnimplementedStreamServiceServer) GetFeature(context.Context, *StreamRequest) (*StreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
}
func (UnimplementedStreamServiceServer) ListFeatures(*StreamRequest, StreamService_ListFeaturesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFeatures not implemented")
}
func (UnimplementedStreamServiceServer) RecordRoute(StreamService_RecordRouteServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordRoute not implemented")
}
func (UnimplementedStreamServiceServer) RouteChat(StreamService_RouteChatServer) error {
	return status.Errorf(codes.Unimplemented, "method RouteChat not implemented")
}
func (UnimplementedStreamServiceServer) mustEmbedUnimplementedStreamServiceServer() {}

// UnsafeStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServiceServer will
// result in compilation errors.
type UnsafeStreamServiceServer interface {
	mustEmbedUnimplementedStreamServiceServer()
}

func RegisterStreamServiceServer(s grpc.ServiceRegistrar, srv StreamServiceServer) {
	s.RegisterService(&StreamService_ServiceDesc, srv)
}

func _StreamService_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/streamingpb.StreamService/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).GetFeature(ctx, req.(*StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamService_ListFeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).ListFeatures(m, &streamServiceListFeaturesServer{stream})
}

type StreamService_ListFeaturesServer interface {
	Send(*StreamResponse) error
	grpc.ServerStream
}

type streamServiceListFeaturesServer struct {
	grpc.ServerStream
}

func (x *streamServiceListFeaturesServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).RecordRoute(&streamServiceRecordRouteServer{stream})
}

type StreamService_RecordRouteServer interface {
	SendAndClose(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamServiceRecordRouteServer struct {
	grpc.ServerStream
}

func (x *streamServiceRecordRouteServer) SendAndClose(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceRecordRouteServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).RouteChat(&streamServiceRouteChatServer{stream})
}

type StreamService_RouteChatServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamServiceRouteChatServer struct {
	grpc.ServerStream
}

func (x *streamServiceRouteChatServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceRouteChatServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamService_ServiceDesc is the grpc.ServiceDesc for StreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "streamingpb.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _StreamService_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeatures",
			Handler:       _StreamService_ListFeatures_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _StreamService_RecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteChat",
			Handler:       _StreamService_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}
