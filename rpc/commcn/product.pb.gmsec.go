// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package commcn

import (
	context "context"
	micro "github.com/gmsec/micro"
	client "github.com/gmsec/micro/client"
	server "github.com/gmsec/micro/server"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	common "rpc/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface
var _ server.Server
var _ client.Client
var _ micro.Service

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductClient interface {
	// Oauth 添加一个产品(可能更新)
	ReplaceProduct(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*common.ResultResp, error)
}

type productClient struct {
	cc client.Client
}

// GetProductName get client name(package.class)
func GetProductName() string {
	return "commcn.Product"
}

// GetProductClient get client by clientname
func GetProductClient() ProductClient {
	cc := micro.GetClient(GetProductName())
	return &productClient{cc}
}

// GetProductClientByName get client by custom name
func GetProductClientByName(name string) ProductClient {
	cc := micro.GetClient(name)
	return &productClient{cc}
}

func NewProductClient(cc client.Client) ProductClient {
	return &productClient{cc}
}

func (c *productClient) ReplaceProduct(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*common.ResultResp, error) {
	conn, err := c.cc.Next()
	defer conn.Close()
	if err != nil {
		return nil, err
	}
	out := new(common.ResultResp)
	err = conn.Invoke(ctx, "/commcn.Product/ReplaceProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
type ProductServer interface {
	// Oauth 添加一个产品(可能更新)
	ReplaceProduct(context.Context, *AddReq) (*common.ResultResp, error)
}

// UnimplementedProductServer can be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (*UnimplementedProductServer) ReplaceProduct(context.Context, *AddReq) (*common.ResultResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplaceProduct not implemented")
}

func RegisterProductServer(s server.Server, srv ProductServer) {
	s.GetServer().RegisterService(&_Product_serviceDesc, srv)
}

func _Product_ReplaceProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ReplaceProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commcn.Product/ReplaceProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ReplaceProduct(ctx, req.(*AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Product_serviceDesc = grpc.ServiceDesc{
	ServiceName: "commcn.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReplaceProduct",
			Handler:    _Product_ReplaceProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commcn/product.proto",
}