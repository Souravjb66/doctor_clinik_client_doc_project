// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: first.proto

package __

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

// DoctorClient is the client API for Doctor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoctorClient interface {
	DoctorOn(ctx context.Context, opts ...grpc.CallOption) (Doctor_DoctorOnClient, error)
	ServerSend(ctx context.Context, in *Clientmsg, opts ...grpc.CallOption) (Doctor_ServerSendClient, error)
}

type doctorClient struct {
	cc grpc.ClientConnInterface
}

func NewDoctorClient(cc grpc.ClientConnInterface) DoctorClient {
	return &doctorClient{cc}
}

func (c *doctorClient) DoctorOn(ctx context.Context, opts ...grpc.CallOption) (Doctor_DoctorOnClient, error) {
	stream, err := c.cc.NewStream(ctx, &Doctor_ServiceDesc.Streams[0], "/Doctor/DoctorOn", opts...)
	if err != nil {
		return nil, err
	}
	x := &doctorDoctorOnClient{stream}
	return x, nil
}

type Doctor_DoctorOnClient interface {
	Send(*Docavailable) error
	CloseAndRecv() (*Serverrespond, error)
	grpc.ClientStream
}

type doctorDoctorOnClient struct {
	grpc.ClientStream
}

func (x *doctorDoctorOnClient) Send(m *Docavailable) error {
	return x.ClientStream.SendMsg(m)
}

func (x *doctorDoctorOnClient) CloseAndRecv() (*Serverrespond, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Serverrespond)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *doctorClient) ServerSend(ctx context.Context, in *Clientmsg, opts ...grpc.CallOption) (Doctor_ServerSendClient, error) {
	stream, err := c.cc.NewStream(ctx, &Doctor_ServiceDesc.Streams[1], "/Doctor/ServerSend", opts...)
	if err != nil {
		return nil, err
	}
	x := &doctorServerSendClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Doctor_ServerSendClient interface {
	Recv() (*Docinfo, error)
	grpc.ClientStream
}

type doctorServerSendClient struct {
	grpc.ClientStream
}

func (x *doctorServerSendClient) Recv() (*Docinfo, error) {
	m := new(Docinfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DoctorServer is the server API for Doctor service.
// All implementations must embed UnimplementedDoctorServer
// for forward compatibility
type DoctorServer interface {
	DoctorOn(Doctor_DoctorOnServer) error
	ServerSend(*Clientmsg, Doctor_ServerSendServer) error
	mustEmbedUnimplementedDoctorServer()
}

// UnimplementedDoctorServer must be embedded to have forward compatible implementations.
type UnimplementedDoctorServer struct {
}

func (UnimplementedDoctorServer) DoctorOn(Doctor_DoctorOnServer) error {
	return status.Errorf(codes.Unimplemented, "method DoctorOn not implemented")
}
func (UnimplementedDoctorServer) ServerSend(*Clientmsg, Doctor_ServerSendServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerSend not implemented")
}
func (UnimplementedDoctorServer) mustEmbedUnimplementedDoctorServer() {}

// UnsafeDoctorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoctorServer will
// result in compilation errors.
type UnsafeDoctorServer interface {
	mustEmbedUnimplementedDoctorServer()
}

func RegisterDoctorServer(s grpc.ServiceRegistrar, srv DoctorServer) {
	s.RegisterService(&Doctor_ServiceDesc, srv)
}

func _Doctor_DoctorOn_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DoctorServer).DoctorOn(&doctorDoctorOnServer{stream})
}

type Doctor_DoctorOnServer interface {
	SendAndClose(*Serverrespond) error
	Recv() (*Docavailable, error)
	grpc.ServerStream
}

type doctorDoctorOnServer struct {
	grpc.ServerStream
}

func (x *doctorDoctorOnServer) SendAndClose(m *Serverrespond) error {
	return x.ServerStream.SendMsg(m)
}

func (x *doctorDoctorOnServer) Recv() (*Docavailable, error) {
	m := new(Docavailable)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Doctor_ServerSend_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Clientmsg)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DoctorServer).ServerSend(m, &doctorServerSendServer{stream})
}

type Doctor_ServerSendServer interface {
	Send(*Docinfo) error
	grpc.ServerStream
}

type doctorServerSendServer struct {
	grpc.ServerStream
}

func (x *doctorServerSendServer) Send(m *Docinfo) error {
	return x.ServerStream.SendMsg(m)
}

// Doctor_ServiceDesc is the grpc.ServiceDesc for Doctor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Doctor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Doctor",
	HandlerType: (*DoctorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DoctorOn",
			Handler:       _Doctor_DoctorOn_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerSend",
			Handler:       _Doctor_ServerSend_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "first.proto",
}
