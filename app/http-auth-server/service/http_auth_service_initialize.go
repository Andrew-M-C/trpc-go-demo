// Package service 实现 HTTP 认证服务
package service

import (
	"fmt"
	"time"

	"github.com/Andrew-M-C/trpc-go-demo/proto/httpauth"
	"github.com/Andrew-M-C/trpc-go-demo/proto/user"
	"github.com/go-playground/validator/v10"
)

func New(d Dependency) (httpauth.AuthService, error) {
	impl, err := newAuthServiceImpl(d)
	if err != nil {
		return nil, err
	}
	return impl, nil
}

type Dependency struct {
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
