package service

import (
	"context"
	"fmt"

	"github.com/Andrew-M-C/trpc-go-demo/app/user/entity"
	"github.com/Andrew-M-C/trpc-go-demo/entity/errs"
	pb "github.com/Andrew-M-C/trpc-go-demo/proto/user"
	"trpc.group/trpc-go/trpc-go/server"
)

// RegisterUserService 注册用户服务
func RegisterUserService(s server.Service, d Dependency) error {
	impl := &userImpl{dep: d}
	pb.RegisterUserService(s, impl)
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
	ctx context.Context, req *pb.GetAccountByUserNameRequest,
) (rsp *pb.GetAccountByUserNameResponse, _ error) {
	u, err := impl.dep.QueryAccountByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.DBError, err)
	}
	if u == nil {
		return nil, errs.UserNotExist
	}

	rsp = &pb.GetAccountByUserNameResponse{Data: &pb.GetAccountByUserNameResponse_Data{}}
	rsp.Data.UserId = u.ID
	rsp.Data.Username = u.Username
	rsp.Data.PasswordHash = u.PasswordHash
	rsp.ErrMsg = "success"
	return
}
