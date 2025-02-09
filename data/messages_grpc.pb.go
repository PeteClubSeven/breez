// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: messages.proto

package data

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

// BreezAPIClient is the client API for BreezAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BreezAPIClient interface {
	GetLSPList(ctx context.Context, in *LSPListRequest, opts ...grpc.CallOption) (*LSPList, error)
	ConnectToLSP(ctx context.Context, in *ConnectLSPRequest, opts ...grpc.CallOption) (*ConnectLSPReply, error)
	AddFundInit(ctx context.Context, in *AddFundInitRequest, opts ...grpc.CallOption) (*AddFundInitReply, error)
	GetFundStatus(ctx context.Context, in *FundStatusRequest, opts ...grpc.CallOption) (*FundStatusReply, error)
	AddInvoice(ctx context.Context, in *AddInvoiceRequest, opts ...grpc.CallOption) (*AddInvoiceReply, error)
	PayInvoice(ctx context.Context, in *PayInvoiceRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
	RestartDaemon(ctx context.Context, in *RestartDaemonRequest, opts ...grpc.CallOption) (*RestartDaemonReply, error)
	ListPayments(ctx context.Context, in *ListPaymentsRequest, opts ...grpc.CallOption) (*PaymentsList, error)
}

type breezAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewBreezAPIClient(cc grpc.ClientConnInterface) BreezAPIClient {
	return &breezAPIClient{cc}
}

func (c *breezAPIClient) GetLSPList(ctx context.Context, in *LSPListRequest, opts ...grpc.CallOption) (*LSPList, error) {
	out := new(LSPList)
	err := c.cc.Invoke(ctx, "/data.BreezAPI/GetLSPList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *breezAPIClient) ConnectToLSP(ctx context.Context, in *ConnectLSPRequest, opts ...grpc.CallOption) (*ConnectLSPReply, error) {
	out := new(ConnectLSPReply)
	err := c.cc.Invoke(ctx, "/data.BreezAPI/ConnectToLSP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *breezAPIClient) AddFundInit(ctx context.Context, in *AddFundInitRequest, opts ...grpc.CallOption) (*AddFundInitReply, error) {
	out := new(AddFundInitReply)
	err := c.cc.Invoke(ctx, "/data.BreezAPI/AddFundInit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *breezAPIClient) GetFundStatus(ctx context.Context, in *FundStatusRequest, opts ...grpc.CallOption) (*FundStatusReply, error) {
	out := new(FundStatusReply)
	err := c.cc.Invoke(ctx, "/data.BreezAPI/GetFundStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *breezAPIClient) AddInvoice(ctx context.Context, in *AddInvoiceRequest, opts ...grpc.CallOption) (*AddInvoiceReply, error) {
	out := new(AddInvoiceReply)
	err := c.cc.Invoke(ctx, "/data.BreezAPI/AddInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *breezAPIClient) PayInvoice(ctx context.Context, in *PayInvoiceRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, "/data.BreezAPI/PayInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *breezAPIClient) RestartDaemon(ctx context.Context, in *RestartDaemonRequest, opts ...grpc.CallOption) (*RestartDaemonReply, error) {
	out := new(RestartDaemonReply)
	err := c.cc.Invoke(ctx, "/data.BreezAPI/RestartDaemon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *breezAPIClient) ListPayments(ctx context.Context, in *ListPaymentsRequest, opts ...grpc.CallOption) (*PaymentsList, error) {
	out := new(PaymentsList)
	err := c.cc.Invoke(ctx, "/data.BreezAPI/ListPayments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BreezAPIServer is the server API for BreezAPI service.
// All implementations must embed UnimplementedBreezAPIServer
// for forward compatibility
type BreezAPIServer interface {
	GetLSPList(context.Context, *LSPListRequest) (*LSPList, error)
	ConnectToLSP(context.Context, *ConnectLSPRequest) (*ConnectLSPReply, error)
	AddFundInit(context.Context, *AddFundInitRequest) (*AddFundInitReply, error)
	GetFundStatus(context.Context, *FundStatusRequest) (*FundStatusReply, error)
	AddInvoice(context.Context, *AddInvoiceRequest) (*AddInvoiceReply, error)
	PayInvoice(context.Context, *PayInvoiceRequest) (*PaymentResponse, error)
	RestartDaemon(context.Context, *RestartDaemonRequest) (*RestartDaemonReply, error)
	ListPayments(context.Context, *ListPaymentsRequest) (*PaymentsList, error)
	mustEmbedUnimplementedBreezAPIServer()
}

// UnimplementedBreezAPIServer must be embedded to have forward compatible implementations.
type UnimplementedBreezAPIServer struct {
}

func (UnimplementedBreezAPIServer) GetLSPList(context.Context, *LSPListRequest) (*LSPList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLSPList not implemented")
}
func (UnimplementedBreezAPIServer) ConnectToLSP(context.Context, *ConnectLSPRequest) (*ConnectLSPReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectToLSP not implemented")
}
func (UnimplementedBreezAPIServer) AddFundInit(context.Context, *AddFundInitRequest) (*AddFundInitReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFundInit not implemented")
}
func (UnimplementedBreezAPIServer) GetFundStatus(context.Context, *FundStatusRequest) (*FundStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFundStatus not implemented")
}
func (UnimplementedBreezAPIServer) AddInvoice(context.Context, *AddInvoiceRequest) (*AddInvoiceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInvoice not implemented")
}
func (UnimplementedBreezAPIServer) PayInvoice(context.Context, *PayInvoiceRequest) (*PaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayInvoice not implemented")
}
func (UnimplementedBreezAPIServer) RestartDaemon(context.Context, *RestartDaemonRequest) (*RestartDaemonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartDaemon not implemented")
}
func (UnimplementedBreezAPIServer) ListPayments(context.Context, *ListPaymentsRequest) (*PaymentsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPayments not implemented")
}
func (UnimplementedBreezAPIServer) mustEmbedUnimplementedBreezAPIServer() {}

// UnsafeBreezAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BreezAPIServer will
// result in compilation errors.
type UnsafeBreezAPIServer interface {
	mustEmbedUnimplementedBreezAPIServer()
}

func RegisterBreezAPIServer(s grpc.ServiceRegistrar, srv BreezAPIServer) {
	s.RegisterService(&BreezAPI_ServiceDesc, srv)
}

func _BreezAPI_GetLSPList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LSPListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BreezAPIServer).GetLSPList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.BreezAPI/GetLSPList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BreezAPIServer).GetLSPList(ctx, req.(*LSPListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BreezAPI_ConnectToLSP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectLSPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BreezAPIServer).ConnectToLSP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.BreezAPI/ConnectToLSP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BreezAPIServer).ConnectToLSP(ctx, req.(*ConnectLSPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BreezAPI_AddFundInit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFundInitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BreezAPIServer).AddFundInit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.BreezAPI/AddFundInit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BreezAPIServer).AddFundInit(ctx, req.(*AddFundInitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BreezAPI_GetFundStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FundStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BreezAPIServer).GetFundStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.BreezAPI/GetFundStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BreezAPIServer).GetFundStatus(ctx, req.(*FundStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BreezAPI_AddInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BreezAPIServer).AddInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.BreezAPI/AddInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BreezAPIServer).AddInvoice(ctx, req.(*AddInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BreezAPI_PayInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayInvoiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BreezAPIServer).PayInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.BreezAPI/PayInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BreezAPIServer).PayInvoice(ctx, req.(*PayInvoiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BreezAPI_RestartDaemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartDaemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BreezAPIServer).RestartDaemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.BreezAPI/RestartDaemon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BreezAPIServer).RestartDaemon(ctx, req.(*RestartDaemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BreezAPI_ListPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BreezAPIServer).ListPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.BreezAPI/ListPayments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BreezAPIServer).ListPayments(ctx, req.(*ListPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BreezAPI_ServiceDesc is the grpc.ServiceDesc for BreezAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BreezAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "data.BreezAPI",
	HandlerType: (*BreezAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLSPList",
			Handler:    _BreezAPI_GetLSPList_Handler,
		},
		{
			MethodName: "ConnectToLSP",
			Handler:    _BreezAPI_ConnectToLSP_Handler,
		},
		{
			MethodName: "AddFundInit",
			Handler:    _BreezAPI_AddFundInit_Handler,
		},
		{
			MethodName: "GetFundStatus",
			Handler:    _BreezAPI_GetFundStatus_Handler,
		},
		{
			MethodName: "AddInvoice",
			Handler:    _BreezAPI_AddInvoice_Handler,
		},
		{
			MethodName: "PayInvoice",
			Handler:    _BreezAPI_PayInvoice_Handler,
		},
		{
			MethodName: "RestartDaemon",
			Handler:    _BreezAPI_RestartDaemon_Handler,
		},
		{
			MethodName: "ListPayments",
			Handler:    _BreezAPI_ListPayments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messages.proto",
}
