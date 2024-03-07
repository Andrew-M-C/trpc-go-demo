// Package service 实现 HTTP 认证服务
package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Andrew-M-C/trpc-go-demo/proto/httpauth"
	"github.com/Andrew-M-C/trpc-go-demo/proto/user"
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
) (rsp *httpauth.LoginResponse, _ error) {
	rsp = &httpauth.LoginResponse{}

	defer func(start time.Time) {
		metrics.IncrCounter("demo.auth.Login.cnt", 1)
		if rsp.ErrCode == 0 {
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
	if err != nil {
		rsp.ErrCode, rsp.ErrMsg = -1, fmt.Sprintf("调用 user 服务失败: %v", err)
		return
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
