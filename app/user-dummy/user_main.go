package main

import (
	"github.com/Andrew-M-C/trpc-go-demo/app/user-dummy/service"
	"github.com/Andrew-M-C/trpc-go-utils/errs"
	"github.com/Andrew-M-C/trpc-go-utils/tracelog"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	initialize()
	s := trpc.NewServer()

	if err := service.RegisterUserService(s); err != nil {
		log.Fatalf("注册用户服务失败: %v", err)
	}

	if err := s.Serve(); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}

func initialize() {
	errs.RegisterErrToCodeFilter()
	tracelog.RegisterTraceLogFilter()
}
