// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: blotservice/v1beta1/blotservice.proto

package blotservicepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BlotService_CreateGameSet_FullMethodName        = "/blotservice.v1beta1.BlotService/CreateGameSet"
	BlotService_JoinGameSet_FullMethodName          = "/blotservice.v1beta1.BlotService/JoinGameSet"
	BlotService_LeaveGameSet_FullMethodName         = "/blotservice.v1beta1.BlotService/LeaveGameSet"
	BlotService_StartGame_FullMethodName            = "/blotservice.v1beta1.BlotService/StartGame"
	BlotService_GetGameSetForPlayer_FullMethodName  = "/blotservice.v1beta1.BlotService/GetGameSetForPlayer"
	BlotService_GetGameSetsForPlayer_FullMethodName = "/blotservice.v1beta1.BlotService/GetGameSetsForPlayer"
)

// BlotServiceClient is the client API for BlotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlotServiceClient interface {
	CreateGameSet(ctx context.Context, in *CreateGameSetRequest, opts ...grpc.CallOption) (*CreateGameSetResponse, error)
	JoinGameSet(ctx context.Context, in *JoinGameSetRequest, opts ...grpc.CallOption) (*JoinGameSetResponse, error)
	LeaveGameSet(ctx context.Context, in *LeaveGameSetRequest, opts ...grpc.CallOption) (*LeaveGameSetResponse, error)
	StartGame(ctx context.Context, in *StartGameRequest, opts ...grpc.CallOption) (*StartGameResponse, error)
	GetGameSetForPlayer(ctx context.Context, in *GetGameSetForPlayerRequest, opts ...grpc.CallOption) (*GetGameSetForPlayerResponse, error)
	GetGameSetsForPlayer(ctx context.Context, in *GetGameSetsForPlayerRequest, opts ...grpc.CallOption) (*GetGameSetsForPlayerResponse, error)
}

type blotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlotServiceClient(cc grpc.ClientConnInterface) BlotServiceClient {
	return &blotServiceClient{cc}
}

func (c *blotServiceClient) CreateGameSet(ctx context.Context, in *CreateGameSetRequest, opts ...grpc.CallOption) (*CreateGameSetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateGameSetResponse)
	err := c.cc.Invoke(ctx, BlotService_CreateGameSet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blotServiceClient) JoinGameSet(ctx context.Context, in *JoinGameSetRequest, opts ...grpc.CallOption) (*JoinGameSetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinGameSetResponse)
	err := c.cc.Invoke(ctx, BlotService_JoinGameSet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blotServiceClient) LeaveGameSet(ctx context.Context, in *LeaveGameSetRequest, opts ...grpc.CallOption) (*LeaveGameSetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LeaveGameSetResponse)
	err := c.cc.Invoke(ctx, BlotService_LeaveGameSet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blotServiceClient) StartGame(ctx context.Context, in *StartGameRequest, opts ...grpc.CallOption) (*StartGameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartGameResponse)
	err := c.cc.Invoke(ctx, BlotService_StartGame_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blotServiceClient) GetGameSetForPlayer(ctx context.Context, in *GetGameSetForPlayerRequest, opts ...grpc.CallOption) (*GetGameSetForPlayerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGameSetForPlayerResponse)
	err := c.cc.Invoke(ctx, BlotService_GetGameSetForPlayer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blotServiceClient) GetGameSetsForPlayer(ctx context.Context, in *GetGameSetsForPlayerRequest, opts ...grpc.CallOption) (*GetGameSetsForPlayerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGameSetsForPlayerResponse)
	err := c.cc.Invoke(ctx, BlotService_GetGameSetsForPlayer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlotServiceServer is the server API for BlotService service.
// All implementations must embed UnimplementedBlotServiceServer
// for forward compatibility.
type BlotServiceServer interface {
	CreateGameSet(context.Context, *CreateGameSetRequest) (*CreateGameSetResponse, error)
	JoinGameSet(context.Context, *JoinGameSetRequest) (*JoinGameSetResponse, error)
	LeaveGameSet(context.Context, *LeaveGameSetRequest) (*LeaveGameSetResponse, error)
	StartGame(context.Context, *StartGameRequest) (*StartGameResponse, error)
	GetGameSetForPlayer(context.Context, *GetGameSetForPlayerRequest) (*GetGameSetForPlayerResponse, error)
	GetGameSetsForPlayer(context.Context, *GetGameSetsForPlayerRequest) (*GetGameSetsForPlayerResponse, error)
	mustEmbedUnimplementedBlotServiceServer()
}

// UnimplementedBlotServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBlotServiceServer struct{}

func (UnimplementedBlotServiceServer) CreateGameSet(context.Context, *CreateGameSetRequest) (*CreateGameSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGameSet not implemented")
}
func (UnimplementedBlotServiceServer) JoinGameSet(context.Context, *JoinGameSetRequest) (*JoinGameSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGameSet not implemented")
}
func (UnimplementedBlotServiceServer) LeaveGameSet(context.Context, *LeaveGameSetRequest) (*LeaveGameSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveGameSet not implemented")
}
func (UnimplementedBlotServiceServer) StartGame(context.Context, *StartGameRequest) (*StartGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartGame not implemented")
}
func (UnimplementedBlotServiceServer) GetGameSetForPlayer(context.Context, *GetGameSetForPlayerRequest) (*GetGameSetForPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameSetForPlayer not implemented")
}
func (UnimplementedBlotServiceServer) GetGameSetsForPlayer(context.Context, *GetGameSetsForPlayerRequest) (*GetGameSetsForPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameSetsForPlayer not implemented")
}
func (UnimplementedBlotServiceServer) mustEmbedUnimplementedBlotServiceServer() {}
func (UnimplementedBlotServiceServer) testEmbeddedByValue()                     {}

// UnsafeBlotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlotServiceServer will
// result in compilation errors.
type UnsafeBlotServiceServer interface {
	mustEmbedUnimplementedBlotServiceServer()
}

func RegisterBlotServiceServer(s grpc.ServiceRegistrar, srv BlotServiceServer) {
	// If the following call pancis, it indicates UnimplementedBlotServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BlotService_ServiceDesc, srv)
}

func _BlotService_CreateGameSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlotServiceServer).CreateGameSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlotService_CreateGameSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlotServiceServer).CreateGameSet(ctx, req.(*CreateGameSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlotService_JoinGameSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinGameSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlotServiceServer).JoinGameSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlotService_JoinGameSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlotServiceServer).JoinGameSet(ctx, req.(*JoinGameSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlotService_LeaveGameSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveGameSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlotServiceServer).LeaveGameSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlotService_LeaveGameSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlotServiceServer).LeaveGameSet(ctx, req.(*LeaveGameSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlotService_StartGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlotServiceServer).StartGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlotService_StartGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlotServiceServer).StartGame(ctx, req.(*StartGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlotService_GetGameSetForPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameSetForPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlotServiceServer).GetGameSetForPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlotService_GetGameSetForPlayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlotServiceServer).GetGameSetForPlayer(ctx, req.(*GetGameSetForPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlotService_GetGameSetsForPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameSetsForPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlotServiceServer).GetGameSetsForPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlotService_GetGameSetsForPlayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlotServiceServer).GetGameSetsForPlayer(ctx, req.(*GetGameSetsForPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlotService_ServiceDesc is the grpc.ServiceDesc for BlotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blotservice.v1beta1.BlotService",
	HandlerType: (*BlotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGameSet",
			Handler:    _BlotService_CreateGameSet_Handler,
		},
		{
			MethodName: "JoinGameSet",
			Handler:    _BlotService_JoinGameSet_Handler,
		},
		{
			MethodName: "LeaveGameSet",
			Handler:    _BlotService_LeaveGameSet_Handler,
		},
		{
			MethodName: "StartGame",
			Handler:    _BlotService_StartGame_Handler,
		},
		{
			MethodName: "GetGameSetForPlayer",
			Handler:    _BlotService_GetGameSetForPlayer_Handler,
		},
		{
			MethodName: "GetGameSetsForPlayer",
			Handler:    _BlotService_GetGameSetsForPlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blotservice/v1beta1/blotservice.proto",
}
