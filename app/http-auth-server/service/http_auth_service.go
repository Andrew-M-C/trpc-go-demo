// Package service 实现 HTTP 认证服务
package service

import (
	"context"
	"errors"

	"github.com/Andrew-M-C/trpc-go-demo/proto/httpauth"
	"trpc.group/trpc-go/trpc-go/server"
)

// RegisterAuthService 注册 HTTP 认证服务
func RegisterAuthService(s server.Service) error {
	impl := newAuthServiceImpl()
	httpauth.RegisterAuthService(s, impl)
	return nil
}

type Dependency interface{}

type authServiceImpl struct{}

func newAuthServiceImpl() authServiceImpl {
	return authServiceImpl{}
}

// Login 实现 http 登录
func (authServiceImpl) Login(
	ctx context.Context, req *httpauth.LoginRequest,
) (rsp *httpauth.LoginResponse, err error) {
	rsp = &httpauth.LoginResponse{}
	// TODO:
	err = errors.New("未实现")
	return
}
