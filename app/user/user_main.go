package main

import (
	"github.com/Andrew-M-C/trpc-go-demo/app/user/repo"
	"github.com/Andrew-M-C/trpc-go-demo/app/user/service"
	"github.com/Andrew-M-C/trpc-go-utils/tracelog"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"

	_ "trpc.group/trpc-go/trpc-naming-polarismesh"
)

func main() {
	tracelog.RegisterTraceLogFilter()
	s := trpc.NewServer()

	r, err := initializeRepo()
	if err != nil {
		log.Fatalf("初始化 repo 失败: %v", err)
	}

	if err := service.RegisterUserService(s, r); err != nil {
		log.Fatalf("注册用户服务失败: %v", err)
	}

	if err := s.Serve(); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}

func initializeRepo() (*repo.Repo, error) {
	dep := repo.Dependency{
		UserAccountDBClientName: "db.mysql.userAccount",
	}
	r, err := repo.NewRepo(dep)
	if err != nil {
		return nil, err
	}

	return r, nil
}
