// Package service 实现假 user 服务, 固定返回错误
package service

import (
	"context"

	"github.com/Andrew-M-C/trpc-go-demo/proto/user"
	"github.com/Andrew-M-C/trpc-go-utils/errs"
	"trpc.group/trpc-go/trpc-go/server"
)

var (
	errServerIsNotReal = errs.New(10101, "当前服务是一个用于测试的假服务")
)

// RegisterUserService 注册用户服务
func RegisterUserService(s server.Service) error {
	impl := &userImpl{}
	user.RegisterUserService(s, impl)
	return nil
}

type userImpl struct{}

// GetAccountByUserName 根据用户名获取帐户信息
func (impl *userImpl) GetAccountByUserName(
	ctx context.Context, req *user.GetAccountByUserNameRequest,
) (*user.GetAccountByUserNameResponse, error) {
	return nil, errServerIsNotReal
}
