package service

import (
	"context"

	"github.com/Andrew-M-C/trpc-go-demo/app/user/entity"
	"github.com/Andrew-M-C/trpc-go-demo/proto/user"
	"trpc.group/trpc-go/trpc-go/log"
	"trpc.group/trpc-go/trpc-go/server"
)

// RegisterUserService 注册用户服务
func RegisterUserService(s server.Service, d Dependency) error {
	impl := &userImpl{dep: d}
	user.RegisterUserService(s, impl)
	return nil
}

// Dependency 表示用户服务初始化依赖
type Dependency interface {
	// QueryAccountByUsername 通过用户名查询帐户信息, 如果帐户不存在则返回 (nil, nil)
	QueryAccountByUsername(ctx context.Context, username string) (*entity.Account, error)
}

type userImpl struct {
	dep Dependency
}

// GetAccountByUserName 根据用户名获取帐户信息
func (impl *userImpl) GetAccountByUserName(
	ctx context.Context, req *user.GetAccountByUserNameRequest,
) (rsp *user.GetAccountByUserNameResponse, _ error) {
	rsp = &user.GetAccountByUserNameResponse{}

	u, err := impl.dep.QueryAccountByUsername(ctx, req.Username)
	if err != nil {
		log.ErrorContextf(ctx, "查询 username '%s' 失败: %v", req.Username, err)
		rsp.ErrCode = -1 // TODO: 采用规范的错误码定义
		rsp.ErrMsg = err.Error()
		return
	}
	if u == nil {
		log.InfoContextf(ctx, "username '%s' 不存在", req.Username)
		rsp.ErrCode = 404
		rsp.ErrMsg = "用户不存在"
		return
	}

	rsp.UserId = u.ID
	rsp.Username = u.Username
	rsp.PasswordHash = u.PasswordHash
	rsp.ErrMsg = "success"
	return
}
