// Package service 实现 HTTP 认证服务
package service

import (
	"fmt"
	"time"

	"github.com/Andrew-M-C/trpc-go-demo/proto/httpauth"
	"github.com/Andrew-M-C/trpc-go-demo/proto/user"
	"github.com/go-playground/validator/v10"
	"trpc.group/trpc-go/trpc-go/server"
)

// RegisterAuthService 注册 HTTP 认证服务
func RegisterAuthService(dep Dependency) error {
	impl, err := newAuthServiceImpl(dep)
	if err != nil {
		return err
	}
	httpauth.RegisterAuthService(dep.Service, impl)
	return nil
}

type Dependency struct {
	Service   server.Service       `validate:"required"`
	UserProxy user.UserClientProxy `validate:"required"`
}

type Repo interface {
	GetEnv() string
}

type authServiceImpl struct {
	Dependency

	Timezone *time.Location
}

func newAuthServiceImpl(dep Dependency) (*authServiceImpl, error) {
	if err := validator.New().Struct(dep); err != nil {
		return nil, err
	}
	tz, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return nil, fmt.Errorf("加载时区失败 (%w)", err)
	}
	impl := &authServiceImpl{
		Dependency: dep,
		Timezone:   tz,
	}
	return impl, nil
}
