// Package service 实现 HTTP 认证服务
package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Andrew-M-C/trpc-go-demo/entity/errs"
	"github.com/Andrew-M-C/trpc-go-demo/proto/httpauth"
	"github.com/Andrew-M-C/trpc-go-demo/proto/user"
)

// Login 实现 http 登录
func (impl *authServiceImpl) Login(
	ctx context.Context, req *httpauth.LoginRequest,
) (rsp *httpauth.LoginResponse, err error) {
	if err := validateLoginReq(req); err != nil {
		return nil, err
	}

	uReq := &user.GetAccountByUserNameRequest{
		Metadata: req.GetMetadata(),
		Username: req.GetUsername(),
	}
	uRsp, err := impl.UserProxy.GetAccountByUserName(ctx, uReq)
	if err != nil {
		return nil, fmt.Errorf("调用 user 服务失败 (%w)", err)
	}

	if req.GetPasswordHash() != uRsp.GetData().GetPasswordHash() {
		return nil, errs.PasswordError
	}
	return &httpauth.LoginResponse{}, nil
}

func validateLoginReq(req *httpauth.LoginRequest) error {
	if req.GetUsername() == "" {
		return errs.New(errs.ParameterWrongCode, "缺少用户名参数")
	}
	if req.GetPasswordHash() == "" {
		return errs.New(errs.ParameterWrongCode, "缺少密码参数")
	}
	return nil
}

// Synchronize 同步服务器状态
func (impl *authServiceImpl) Synchronize(
	context.Context, *httpauth.SynchronizeRequest,
) (*httpauth.SynchronizeResponse, error) {
	rsp := &httpauth.SynchronizeResponse{
		Data: &httpauth.SynchronizeResponse_Data{},
	}

	now := time.Now().In(impl.Timezone)
	rsp.Data.TsMsec = now.UnixMilli()
	rsp.Data.Desc = now.Format("2006-01-02 15:04:05.000")
	rsp.Data.Timezone = impl.Timezone.String()

	return rsp, nil
}
