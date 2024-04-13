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
)

// Login 实现 http 登录
func (impl *authServiceImpl) Login(
	ctx context.Context, req *httpauth.LoginRequest,
) (rsp *httpauth.LoginResponse, err error) {
	uReq := &user.GetAccountByUserNameRequest{
		Metadata: req.GetMetadata(),
		Username: req.GetUsername(),
	}
	uRsp, err := user.NewUserClientProxy().GetAccountByUserName(ctx, uReq)
	if err != nil {
		log.ErrorContextf(ctx, "user 服务返回错误: %v", err)
		return nil, err
	}

	log.DebugContextf(ctx, "rsp: '%v'", tracelog.ToJSON(uRsp))
	if req.GetPasswordHash() != uRsp.GetPasswordHash() {
		return nil, errentity.PasswordError
	}
	return &httpauth.LoginResponse{}, nil
}

// Synchronize 同步服务器状态
func (impl *authServiceImpl) Synchronize(
	ctx context.Context, req *httpauth.SynchronizeRequest,
) (*httpauth.SynchronizeResponse, error) {
	rsp := &httpauth.SynchronizeResponse{
		Data: &httpauth.SynchronizeResponse_Data{},
	}
	rsp.Data.Env = impl.dep.Repo.GetEnv()
	rsp.Data.TsMsec = time.Now().UnixMilli()
	rsp.Data.Timezone = time.Local.String()

	return rsp, nil
}
