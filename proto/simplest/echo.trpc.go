// Code generated by trpc-go/trpc-cmdline v1.0.9. DO NOT EDIT.
// source: echo.proto

package simplest

import (
	"context"
	"errors"
	"fmt"

	_ "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/codec"
	_ "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/server"
)

// START ======================================= Server Service Definition ======================================= START

// HelloWorldService defines service.
type HelloWorldService interface {
	// Hello Hello says hello.
	Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) // @alias=/demo/Hello
}

func HelloWorldService_Hello_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &HelloRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(HelloWorldService).Hello(ctx, reqbody.(*HelloRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// HelloWorldServer_ServiceDesc descriptor for server.RegisterService.
var HelloWorldServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.demo.simplest.HelloWorld",
	HandlerType: ((*HelloWorldService)(nil)),
	Methods: []server.Method{
		{
			Name: "/demo/Hello",
			Func: HelloWorldService_Hello_Handler,
		},
		{
			Name: "/trpc.demo.simplest.HelloWorld/Hello",
			Func: HelloWorldService_Hello_Handler,
		},
	},
}

// RegisterHelloWorldService registers service.
func RegisterHelloWorldService(s server.Service, svr HelloWorldService) {
	if err := s.Register(&HelloWorldServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("HelloWorld register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedHelloWorld struct{}

// Hello Hello says hello.
func (s *UnimplementedHelloWorld) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, errors.New("rpc Hello of service HelloWorld is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// HelloWorldClientProxy defines service client proxy
type HelloWorldClientProxy interface {
	// Hello Hello says hello.
	Hello(ctx context.Context, req *HelloRequest, opts ...client.Option) (rsp *HelloResponse, err error) // @alias=/demo/Hello
}

type HelloWorldClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewHelloWorldClientProxy = func(opts ...client.Option) HelloWorldClientProxy {
	return &HelloWorldClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *HelloWorldClientProxyImpl) Hello(ctx context.Context, req *HelloRequest, opts ...client.Option) (*HelloResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/demo/Hello")
	msg.WithCalleeServiceName(HelloWorldServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("demo")
	msg.WithCalleeServer("simplest")
	msg.WithCalleeService("HelloWorld")
	msg.WithCalleeMethod("Hello")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &HelloResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END
