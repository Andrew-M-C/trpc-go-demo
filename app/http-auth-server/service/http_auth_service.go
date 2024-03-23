// Package service 实现 HTTP 认证服务
package service

import (
	"context"
	"time"

	errentity "github.com/Andrew-M-C/trpc-go-demo/entity/errs"
	"github.com/Andrew-M-C/trpc-go-demo/proto/httpauth"
	"github.com/Andrew-M-C/trpc-go-demo/proto/user"
	"github.com/Andrew-M-C/trpc-go-utils/tracelog"
	"trpc.group/trpc-go/trpc-go/log"
	"trpc.group/trpc-go/trpc-go/metrics"
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
	defer func(start time.Time) {
		metrics.IncrCounter("demo.auth.Login.cnt", 1)
		if err == nil {
			metrics.IncrCounter("demo.auth.Login.succ", 1)
		} else {
			metrics.IncrCounter("demo.auth.Login.fail", 1)
		}
		metrics.SetGauge("demo.auth.Login.ela", float64(time.Since(start))/float64(time.Second))
	}(time.Now())

	uReq := &user.GetAccountByUserNameRequest{
		Metadata: req.GetMetadata(),
		Username: req.GetUsername(),
	}
	uRsp, err := user.NewUserClientProxy().GetAccountByUserName(ctx, uReq)
	log.DebugContextf(ctx, "rsp: '%v'", tracelog.ToJSON(uRsp))
	if err != nil {
		log.ErrorContextf(ctx, "user 服务返回错误: %v", err)
		return nil, err
	}
	if req.GetPasswordHash() != uRsp.GetPasswordHash() {
		return nil, errentity.PasswordError
	}
	return &httpauth.LoginResponse{}, nil
}
