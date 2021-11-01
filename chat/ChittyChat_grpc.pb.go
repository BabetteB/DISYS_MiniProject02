// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package chat

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

// ChittyChatServiceClient is the client API for ChittyChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChittyChatServiceClient interface {
	Publish(ctx context.Context, opts ...grpc.CallOption) (ChittyChatService_PublishClient, error)
	Broadcast(ctx context.Context, opts ...grpc.CallOption) (ChittyChatService_BroadcastClient, error)
}

type chittyChatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChittyChatServiceClient(cc grpc.ClientConnInterface) ChittyChatServiceClient {
	return &chittyChatServiceClient{cc}
}

func (c *chittyChatServiceClient) Publish(ctx context.Context, opts ...grpc.CallOption) (ChittyChatService_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChittyChatService_ServiceDesc.Streams[0], "/main.ChittyChatService/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &chittyChatServicePublishClient{stream}
	return x, nil
}

type ChittyChatService_PublishClient interface {
	Send(*ClientMessage) error
	Recv() (*StatusMessage, error)
	grpc.ClientStream
}

type chittyChatServicePublishClient struct {
	grpc.ClientStream
}

func (x *chittyChatServicePublishClient) Send(m *ClientMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chittyChatServicePublishClient) Recv() (*StatusMessage, error) {
	m := new(StatusMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chittyChatServiceClient) Broadcast(ctx context.Context, opts ...grpc.CallOption) (ChittyChatService_BroadcastClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChittyChatService_ServiceDesc.Streams[1], "/main.ChittyChatService/Broadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &chittyChatServiceBroadcastClient{stream}
	return x, nil
}

type ChittyChatService_BroadcastClient interface {
	Send(*StatusMessage) error
	Recv() (*ChatRoomMessages, error)
	grpc.ClientStream
}

type chittyChatServiceBroadcastClient struct {
	grpc.ClientStream
}

func (x *chittyChatServiceBroadcastClient) Send(m *StatusMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chittyChatServiceBroadcastClient) Recv() (*ChatRoomMessages, error) {
	m := new(ChatRoomMessages)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChittyChatServiceServer is the server API for ChittyChatService service.
// All implementations must embed UnimplementedChittyChatServiceServer
// for forward compatibility
type ChittyChatServiceServer interface {
	Publish(ChittyChatService_PublishServer) error
	Broadcast(ChittyChatService_BroadcastServer) error
	mustEmbedUnimplementedChittyChatServiceServer()
}

// UnimplementedChittyChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChittyChatServiceServer struct {
}

func (UnimplementedChittyChatServiceServer) Publish(ChittyChatService_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedChittyChatServiceServer) Broadcast(ChittyChatService_BroadcastServer) error {
	return status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedChittyChatServiceServer) mustEmbedUnimplementedChittyChatServiceServer() {}

// UnsafeChittyChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChittyChatServiceServer will
// result in compilation errors.
type UnsafeChittyChatServiceServer interface {
	mustEmbedUnimplementedChittyChatServiceServer()
}

func RegisterChittyChatServiceServer(s grpc.ServiceRegistrar, srv ChittyChatServiceServer) {
	s.RegisterService(&ChittyChatService_ServiceDesc, srv)
}

func _ChittyChatService_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChittyChatServiceServer).Publish(&chittyChatServicePublishServer{stream})
}

type ChittyChatService_PublishServer interface {
	Send(*StatusMessage) error
	Recv() (*ClientMessage, error)
	grpc.ServerStream
}

type chittyChatServicePublishServer struct {
	grpc.ServerStream
}

func (x *chittyChatServicePublishServer) Send(m *StatusMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chittyChatServicePublishServer) Recv() (*ClientMessage, error) {
	m := new(ClientMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChittyChatService_Broadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChittyChatServiceServer).Broadcast(&chittyChatServiceBroadcastServer{stream})
}

type ChittyChatService_BroadcastServer interface {
	Send(*ChatRoomMessages) error
	Recv() (*StatusMessage, error)
	grpc.ServerStream
}

type chittyChatServiceBroadcastServer struct {
	grpc.ServerStream
}

func (x *chittyChatServiceBroadcastServer) Send(m *ChatRoomMessages) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chittyChatServiceBroadcastServer) Recv() (*StatusMessage, error) {
	m := new(StatusMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChittyChatService_ServiceDesc is the grpc.ServiceDesc for ChittyChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChittyChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.ChittyChatService",
	HandlerType: (*ChittyChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Publish",
			Handler:       _ChittyChatService_Publish_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Broadcast",
			Handler:       _ChittyChatService_Broadcast_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ChittyChat.proto",
}
