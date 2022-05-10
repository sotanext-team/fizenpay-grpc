// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: fizenpay-be/merchant.proto

package fizenpay_be

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

// MerchantServiceClient is the client API for MerchantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MerchantServiceClient interface {
	GetAllMerchant(ctx context.Context, in *GetAllMerchantRequest, opts ...grpc.CallOption) (*GetAllMerchantResponse, error)
	ActivateAccount(ctx context.Context, in *ActivateAccountRequest, opts ...grpc.CallOption) (*ActivateAccountResponse, error)
	DeactivateAccount(ctx context.Context, in *DeactivateAccountRequest, opts ...grpc.CallOption) (*DeactivateAccountResponse, error)
	GetOneBySession(ctx context.Context, in *GetOneBySessionRequest, opts ...grpc.CallOption) (*GetOneBySessionResponse, error)
	GetOneByApiKey(ctx context.Context, in *GetOneByApiKeyRequest, opts ...grpc.CallOption) (*GetOneByApiKeyResponse, error)
}

type merchantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantServiceClient(cc grpc.ClientConnInterface) MerchantServiceClient {
	return &merchantServiceClient{cc}
}

func (c *merchantServiceClient) GetAllMerchant(ctx context.Context, in *GetAllMerchantRequest, opts ...grpc.CallOption) (*GetAllMerchantResponse, error) {
	out := new(GetAllMerchantResponse)
	err := c.cc.Invoke(ctx, "/fizenpay_be.MerchantService/GetAllMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) ActivateAccount(ctx context.Context, in *ActivateAccountRequest, opts ...grpc.CallOption) (*ActivateAccountResponse, error) {
	out := new(ActivateAccountResponse)
	err := c.cc.Invoke(ctx, "/fizenpay_be.MerchantService/ActivateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) DeactivateAccount(ctx context.Context, in *DeactivateAccountRequest, opts ...grpc.CallOption) (*DeactivateAccountResponse, error) {
	out := new(DeactivateAccountResponse)
	err := c.cc.Invoke(ctx, "/fizenpay_be.MerchantService/DeactivateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) GetOneBySession(ctx context.Context, in *GetOneBySessionRequest, opts ...grpc.CallOption) (*GetOneBySessionResponse, error) {
	out := new(GetOneBySessionResponse)
	err := c.cc.Invoke(ctx, "/fizenpay_be.MerchantService/GetOneBySession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) GetOneByApiKey(ctx context.Context, in *GetOneByApiKeyRequest, opts ...grpc.CallOption) (*GetOneByApiKeyResponse, error) {
	out := new(GetOneByApiKeyResponse)
	err := c.cc.Invoke(ctx, "/fizenpay_be.MerchantService/GetOneByApiKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantServiceServer is the server API for MerchantService service.
// All implementations must embed UnimplementedMerchantServiceServer
// for forward compatibility
type MerchantServiceServer interface {
	GetAllMerchant(context.Context, *GetAllMerchantRequest) (*GetAllMerchantResponse, error)
	ActivateAccount(context.Context, *ActivateAccountRequest) (*ActivateAccountResponse, error)
	DeactivateAccount(context.Context, *DeactivateAccountRequest) (*DeactivateAccountResponse, error)
	GetOneBySession(context.Context, *GetOneBySessionRequest) (*GetOneBySessionResponse, error)
	GetOneByApiKey(context.Context, *GetOneByApiKeyRequest) (*GetOneByApiKeyResponse, error)
	mustEmbedUnimplementedMerchantServiceServer()
}

// UnimplementedMerchantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMerchantServiceServer struct {
}

func (UnimplementedMerchantServiceServer) GetAllMerchant(context.Context, *GetAllMerchantRequest) (*GetAllMerchantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMerchant not implemented")
}
func (UnimplementedMerchantServiceServer) ActivateAccount(context.Context, *ActivateAccountRequest) (*ActivateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateAccount not implemented")
}
func (UnimplementedMerchantServiceServer) DeactivateAccount(context.Context, *DeactivateAccountRequest) (*DeactivateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateAccount not implemented")
}
func (UnimplementedMerchantServiceServer) GetOneBySession(context.Context, *GetOneBySessionRequest) (*GetOneBySessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneBySession not implemented")
}
func (UnimplementedMerchantServiceServer) GetOneByApiKey(context.Context, *GetOneByApiKeyRequest) (*GetOneByApiKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneByApiKey not implemented")
}
func (UnimplementedMerchantServiceServer) mustEmbedUnimplementedMerchantServiceServer() {}

// UnsafeMerchantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MerchantServiceServer will
// result in compilation errors.
type UnsafeMerchantServiceServer interface {
	mustEmbedUnimplementedMerchantServiceServer()
}

func RegisterMerchantServiceServer(s grpc.ServiceRegistrar, srv MerchantServiceServer) {
	s.RegisterService(&MerchantService_ServiceDesc, srv)
}

func _MerchantService_GetAllMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMerchantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).GetAllMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fizenpay_be.MerchantService/GetAllMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).GetAllMerchant(ctx, req.(*GetAllMerchantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_ActivateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).ActivateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fizenpay_be.MerchantService/ActivateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).ActivateAccount(ctx, req.(*ActivateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_DeactivateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).DeactivateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fizenpay_be.MerchantService/DeactivateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).DeactivateAccount(ctx, req.(*DeactivateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_GetOneBySession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneBySessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).GetOneBySession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fizenpay_be.MerchantService/GetOneBySession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).GetOneBySession(ctx, req.(*GetOneBySessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_GetOneByApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneByApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).GetOneByApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fizenpay_be.MerchantService/GetOneByApiKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).GetOneByApiKey(ctx, req.(*GetOneByApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MerchantService_ServiceDesc is the grpc.ServiceDesc for MerchantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MerchantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fizenpay_be.MerchantService",
	HandlerType: (*MerchantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllMerchant",
			Handler:    _MerchantService_GetAllMerchant_Handler,
		},
		{
			MethodName: "ActivateAccount",
			Handler:    _MerchantService_ActivateAccount_Handler,
		},
		{
			MethodName: "DeactivateAccount",
			Handler:    _MerchantService_DeactivateAccount_Handler,
		},
		{
			MethodName: "GetOneBySession",
			Handler:    _MerchantService_GetOneBySession_Handler,
		},
		{
			MethodName: "GetOneByApiKey",
			Handler:    _MerchantService_GetOneByApiKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fizenpay-be/merchant.proto",
}
