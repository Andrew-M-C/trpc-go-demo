// Package service 实现 HTTP 认证服务
package service

import (
	"context"

	"github.com/Andrew-M-C/trpc-go-demo/proto/httpauth"
	"github.com/Andrew-M-C/trpc-go-demo/proto/user"
	"trpc.group/trpc-go/trpc-go/log"
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
	uReq := &user.GetAccountByUserNameRequest{
		Metadata: req.GetMetadata(),
		Username: req.GetUsername(),
	}
	uRsp, err := user.NewUserClientProxy().GetAccountByUserName(ctx, uReq)
	if err != nil {
		log.ErrorContextf(ctx, "调用 user 服务失败: %v", err)
		return nil, err
	}
	// 用户存在与否
	if uRsp.GetErrCode() != 0 {
		rsp.ErrCode, rsp.ErrMsg = uRsp.GetErrCode(), uRsp.GetErrMsg()
		return
	}

	// 密码检查
	if p := uRsp.GetPasswordHash(); p != req.PasswordHash {
		log.DebugContextf(ctx, "请求的密码为: %s, 实际密码为 %s", req.PasswordHash, p)
		rsp.ErrCode, rsp.ErrMsg = 404, "密码错误"
		return
	}

	rsp.ErrCode, rsp.ErrMsg = 0, "success"
	return
}
