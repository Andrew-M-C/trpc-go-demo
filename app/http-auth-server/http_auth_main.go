package main

import (
	"github.com/Andrew-M-C/trpc-go-demo/app/http-auth-server/service"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()

	if err := service.RegisterAuthService(s); err != nil {
		log.Fatalf("初始化服务失败: %v", err)
	}
	if err := s.Serve(); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}
