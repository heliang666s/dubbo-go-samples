// Code generated by protoc-gen-triple. DO NOT EDIT.
//
// Source: cartservice.proto
package hipstershop

import (
	"context"
)

import (
	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/client"
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/protocol/triple/triple_protocol"
	"dubbo.apache.org/dubbo-go/v3/server"
)

// This is a compile-time assertion to ensure that this generated file and the Triple package
// are compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of Triple newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of Triple or updating the Triple
// version compiled into your binary.
const _ = triple_protocol.IsAtLeastVersion0_1_0

const (
	// CartServiceName is the fully-qualified name of the CartService service.
	CartServiceName = "hipstershop.CartService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CartServiceAddItemProcedure is the fully-qualified name of the CartService's AddItem RPC.
	CartServiceAddItemProcedure = "/hipstershop.CartService/AddItem"
	// CartServiceGetCartProcedure is the fully-qualified name of the CartService's GetCart RPC.
	CartServiceGetCartProcedure = "/hipstershop.CartService/GetCart"
	// CartServiceEmptyCartProcedure is the fully-qualified name of the CartService's EmptyCart RPC.
	CartServiceEmptyCartProcedure = "/hipstershop.CartService/EmptyCart"
)

var (
	_ CartService = (*CartServiceImpl)(nil)
)

// CartService is a client for the hipstershop.CartService service.
type CartService interface {
	AddItem(ctx context.Context, req *AddItemRequest, opts ...client.CallOption) (*Empty, error)
	GetCart(ctx context.Context, req *GetCartRequest, opts ...client.CallOption) (*Cart, error)
	EmptyCart(ctx context.Context, req *EmptyCartRequest, opts ...client.CallOption) (*Empty, error)
}

// NewCartService constructs a client for the hipstershop.CartService service.
func NewCartService(cli *client.Client, opts ...client.ReferenceOption) (CartService, error) {
	conn, err := cli.DialWithInfo("hipstershop.CartService", &CartService_ClientInfo, opts...)
	if err != nil {
		return nil, err
	}
	return &CartServiceImpl{
		conn: conn,
	}, nil
}

func SetConsumerService(srv common.RPCService) {
	dubbo.SetConsumerServiceWithInfo(srv, &CartService_ClientInfo)
}

// CartServiceImpl implements CartService.
type CartServiceImpl struct {
	conn *client.Connection
}

func (c *CartServiceImpl) AddItem(ctx context.Context, req *AddItemRequest, opts ...client.CallOption) (*Empty, error) {
	resp := new(Empty)
	if err := c.conn.CallUnary(ctx, []interface{}{req}, resp, "AddItem", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CartServiceImpl) GetCart(ctx context.Context, req *GetCartRequest, opts ...client.CallOption) (*Cart, error) {
	resp := new(Cart)
	if err := c.conn.CallUnary(ctx, []interface{}{req}, resp, "GetCart", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CartServiceImpl) EmptyCart(ctx context.Context, req *EmptyCartRequest, opts ...client.CallOption) (*Empty, error) {
	resp := new(Empty)
	if err := c.conn.CallUnary(ctx, []interface{}{req}, resp, "EmptyCart", opts...); err != nil {
		return nil, err
	}
	return resp, nil
}

var CartService_ClientInfo = client.ClientInfo{
	InterfaceName: "hipstershop.CartService",
	MethodNames:   []string{"AddItem", "GetCart", "EmptyCart"},
	ConnectionInjectFunc: func(dubboCliRaw interface{}, conn *client.Connection) {
		dubboCli := dubboCliRaw.(*CartServiceImpl)
		dubboCli.conn = conn
	},
}

// CartServiceHandler is an implementation of the hipstershop.CartService service.
type CartServiceHandler interface {
	AddItem(context.Context, *AddItemRequest) (*Empty, error)
	GetCart(context.Context, *GetCartRequest) (*Cart, error)
	EmptyCart(context.Context, *EmptyCartRequest) (*Empty, error)
}

func RegisterCartServiceHandler(srv *server.Server, hdlr CartServiceHandler, opts ...server.ServiceOption) error {
	return srv.Register(hdlr, &CartService_ServiceInfo, opts...)
}

func SetProviderService(srv common.RPCService) {
	dubbo.SetProviderServiceWithInfo(srv, &CartService_ServiceInfo)
}

var CartService_ServiceInfo = server.ServiceInfo{
	InterfaceName: "hipstershop.CartService",
	ServiceType:   (*CartServiceHandler)(nil),
	Methods: []server.MethodInfo{
		{
			Name: "AddItem",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(AddItemRequest)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*AddItemRequest)
				res, err := handler.(CartServiceHandler).AddItem(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
		{
			Name: "GetCart",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(GetCartRequest)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*GetCartRequest)
				res, err := handler.(CartServiceHandler).GetCart(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
		{
			Name: "EmptyCart",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(EmptyCartRequest)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*EmptyCartRequest)
				res, err := handler.(CartServiceHandler).EmptyCart(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
	},
}
