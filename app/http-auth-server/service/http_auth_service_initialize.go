// Package service 实现 HTTP 认证服务
package service

import (
	"github.com/Andrew-M-C/trpc-go-demo/proto/httpauth"
	"trpc.group/trpc-go/trpc-go/server"
)

// RegisterAuthService 注册 HTTP 认证服务
func RegisterAuthService(s server.Service, dep Dependency) error {
	impl, err := newAuthServiceImpl(dep)
	if err != nil {
		return err
	}
	httpauth.RegisterAuthService(s, impl)
	return nil
}

type Dependency struct {
	Repo Repo
}

type Repo interface {
	GetEnv() string
}

type authServiceImpl struct {
	dep Dependency
}

func newAuthServiceImpl(dep Dependency) (*authServiceImpl, error) {
	impl := &authServiceImpl{
		dep: dep,
	}
	return impl, nil
}
