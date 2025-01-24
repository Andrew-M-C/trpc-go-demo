package main

import (
	"context"

	"github.com/Andrew-M-C/trpc-go-demo/proto/user"
	"github.com/Andrew-M-C/trpc-go-utils/log"
	"github.com/Andrew-M-C/trpc-go-utils/recovery"
	"trpc.group/trpc-go/trpc-go/client"
)

func main() {
	defer recovery.CatchPanic(recovery.WithErrorLog())

	a, err := newApplication()
	if err != nil {
		log.New().Text("初始化失败").Err(err).Error()
		return
	}

	if err := a.Serve(); err != nil {
		log.New().Text("启动服务失败").Err(err).Error()
		return
	}
}

type application interface {
	Serve() error
}

type wrappedUserClient struct {
	svc user.UserService
}

func (c wrappedUserClient) GetAccountByUserName(
	ctx context.Context, req *user.GetAccountByUserNameRequest,
	opts ...client.Option,
) (*user.GetAccountByUserNameResponse, error) {
	_ = opts
	return c.svc.GetAccountByUserName(ctx, req)
}
