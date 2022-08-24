// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package capture

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

// CaptureClient is the client API for Capture service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaptureClient interface {
	Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error)
}

type captureClient struct {
	cc grpc.ClientConnInterface
}

func NewCaptureClient(cc grpc.ClientConnInterface) CaptureClient {
	return &captureClient{cc}
}

func (c *captureClient) Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error) {
	out := new(FlushResponse)
	err := c.cc.Invoke(ctx, "/magma.capture.Capture/Flush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaptureServer is the server API for Capture service.
// All implementations must embed UnimplementedCaptureServer
// for forward compatibility
type CaptureServer interface {
	Flush(context.Context, *FlushRequest) (*FlushResponse, error)
	mustEmbedUnimplementedCaptureServer()
}

// UnimplementedCaptureServer must be embedded to have forward compatible implementations.
type UnimplementedCaptureServer struct {
}

func (UnimplementedCaptureServer) Flush(context.Context, *FlushRequest) (*FlushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (UnimplementedCaptureServer) mustEmbedUnimplementedCaptureServer() {}

// UnsafeCaptureServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaptureServer will
// result in compilation errors.
type UnsafeCaptureServer interface {
	mustEmbedUnimplementedCaptureServer()
}

func RegisterCaptureServer(s grpc.ServiceRegistrar, srv CaptureServer) {
	s.RegisterService(&Capture_ServiceDesc, srv)
}

func _Capture_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptureServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.capture.Capture/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptureServer).Flush(ctx, req.(*FlushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Capture_ServiceDesc is the grpc.ServiceDesc for Capture service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Capture_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.capture.Capture",
	HandlerType: (*CaptureServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Flush",
			Handler:    _Capture_Flush_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "capture.proto",
}
