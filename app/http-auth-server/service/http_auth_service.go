// Package service 实现 HTTP 认证服务
package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Andrew-M-C/trpc-go-demo/entity/errs"
	"github.com/Andrew-M-C/trpc-go-demo/proto/achieve"
	"github.com/Andrew-M-C/trpc-go-demo/proto/httpauth"
	"github.com/Andrew-M-C/trpc-go-demo/proto/user"
	"github.com/Andrew-M-C/trpc-go-utils/log"
	"trpc.group/trpc-go/trpc-go"
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

	if req.GetPasswordHash() != uRsp.GetData().GetUserInfo().GetPasswordHash() {
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

// UserSpace 返回用户空间信息
func (impl *authServiceImpl) UserSpace(
	ctx context.Context, req *httpauth.UserSpaceRequest,
) (*httpauth.UserSpaceResponse, error) {
	// 首先换取用户信息
	uReq := &user.GetAccountByUserNameRequest{
		Metadata: req.Metadata,
		Username: req.Username,
	}
	uRsp, err := impl.UserProxy.GetAccountByUserName(ctx, uReq)
	if err != nil {
		return nil, fmt.Errorf("调用用户服务失败 (%w)", err)
	}
	rsp := &httpauth.UserSpaceResponse{
		Data: &httpauth.UserSpaceResponse_Data{},
	}
	rsp.Data.UserInfo = &httpauth.UserSpaceResponse_UserInfo{
		UserId:      uRsp.GetData().GetUserInfo().GetUserId(),
		Username:    uRsp.GetData().GetUserInfo().GetUsername(),
		Nickname:    uRsp.GetData().GetUserInfo().GetNickname(),
		CreateTsSec: uRsp.GetData().GetUserInfo().GetCreateTsSec(),
	}

	// 然后用 user_id 换取用户空间信息

	// 首先是徽章信息
	badgeProc := func() error {
		ctx := log.CloneContextForConcurrency(ctx)
		bReq := &achieve.GetUserBadgesRequest{
			Metadata: req.GetMetadata(),
			UserId:   uRsp.GetData().GetUserInfo().GetUserId(),
		}
		bRsp, err := impl.AchievementProxy.GetUserBadges(ctx, bReq)
		if err != nil {
			return fmt.Errorf("请求成就服务获取徽章信息失败 (%w)", err)
		}
		rsp.Data.Badges = bRsp.GetData().GetBadges()
		return nil
	}
	// 然后是声望信息
	reputationProc := func() error {
		ctx := log.CloneContextForConcurrency(ctx)
		rReq := &achieve.GetUserReputationRequest{
			Metadata: req.Metadata,
			UserId:   uRsp.GetData().GetUserInfo().GetUserId(),
		}
		rRsp, err := impl.AchievementProxy.GetUserReputation(ctx, rReq)
		if err != nil {
			return fmt.Errorf("请求成就服务声望徽章信息失败 (%w)", err)
		}
		rsp.Data.Reputation = rRsp.GetData().GetReputation()
		return nil
	}
	// 并发调用
	if err := trpc.GoAndWait(badgeProc, reputationProc); err != nil {
		return nil, err
	}
	// 返回
	return rsp, nil
}
