// Code generated by trpc-go/trpc-cmdline v1.0.9. DO NOT EDIT.
// source: user.proto

package user

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

// UserService defines service.
type UserService interface {
	GetAccountByUserName(ctx context.Context, req *GetAccountByUserNameRequest) (*GetAccountByUserNameResponse, error)
}

func UserService_GetAccountByUserName_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &GetAccountByUserNameRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(UserService).GetAccountByUserName(ctx, reqbody.(*GetAccountByUserNameRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// UserServer_ServiceDesc descriptor for server.RegisterService.
var UserServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "demo.account.User",
	HandlerType: ((*UserService)(nil)),
	Methods: []server.Method{
		{
			Name: "/demo.account.User/GetAccountByUserName",
			Func: UserService_GetAccountByUserName_Handler,
		},
	},
}

// RegisterUserService registers service.
func RegisterUserService(s server.Service, svr UserService) {
	if err := s.Register(&UserServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("User register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedUser struct{}

func (s *UnimplementedUser) GetAccountByUserName(ctx context.Context, req *GetAccountByUserNameRequest) (*GetAccountByUserNameResponse, error) {
	return nil, errors.New("rpc GetAccountByUserName of service User is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// UserClientProxy defines service client proxy
type UserClientProxy interface {
	GetAccountByUserName(ctx context.Context, req *GetAccountByUserNameRequest, opts ...client.Option) (rsp *GetAccountByUserNameResponse, err error)
}

type UserClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewUserClientProxy = func(opts ...client.Option) UserClientProxy {
	return &UserClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *UserClientProxyImpl) GetAccountByUserName(ctx context.Context, req *GetAccountByUserNameRequest, opts ...client.Option) (*GetAccountByUserNameResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/demo.account.User/GetAccountByUserName")
	msg.WithCalleeServiceName(UserServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("User")
	msg.WithCalleeMethod("GetAccountByUserName")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &GetAccountByUserNameResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END
