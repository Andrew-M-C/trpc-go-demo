package service

import (
	"context"
	"fmt"

	"github.com/Andrew-M-C/trpc-go-demo/app/user/repo"
	"github.com/Andrew-M-C/trpc-go-demo/entity/errs"
	pb "github.com/Andrew-M-C/trpc-go-demo/proto/user"
	"github.com/go-playground/validator/v10"
)

func New(d Dependency) (pb.UserService, error) {
	if err := validator.New().Struct(d); err != nil {
		return nil, err
	}
	impl := &impl{Dependency: d}
	return impl, nil
}

// Dependency 表示用户服务初始化依赖
type Dependency struct {
	AccountRepo repo.AccountRepo `validate:"required"`
}

type impl struct {
	Dependency
}

// GetAccountByUserName 根据用户名获取帐户信息
func (impl *impl) GetAccountByUserName(
	ctx context.Context, req *pb.GetAccountByUserNameRequest,
) (*pb.GetAccountByUserNameResponse, error) {
	u, err := impl.AccountRepo.QueryAccountByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.DBError, err)
	}
	if u == nil {
		return nil, errs.UserNotExist
	}

	rsp := &pb.GetAccountByUserNameResponse{
		Data: &pb.GetAccountByUserNameResponse_Data{
			UserInfo: u,
		},
	}
	return rsp, nil
}
