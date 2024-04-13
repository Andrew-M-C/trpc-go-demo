package main

import (
	"github.com/Andrew-M-C/trpc-go-demo/app/http-auth-server/repo"
	"github.com/Andrew-M-C/trpc-go-demo/app/http-auth-server/service"
	"github.com/Andrew-M-C/trpc-go-demo/utils/filter/count"
	"github.com/Andrew-M-C/trpc-go-demo/utils/filter/elapse"
	"github.com/Andrew-M-C/trpc-go-utils/codec"
	"github.com/Andrew-M-C/trpc-go-utils/errs"
	metricslog "github.com/Andrew-M-C/trpc-go-utils/metrics/log"
	"github.com/Andrew-M-C/trpc-go-utils/tracelog"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	initialize()
	s := trpc.NewServer()

	r, err := repo.NewRepo(repo.Dependency{
		GeneralConfigKeyName: "/http-auth-server/general.json",
	})
	if err != nil {
		log.Fatalf("初始化 repo 失败: %v")
	}

	dep := service.Dependency{Repo: r}
	if err := service.RegisterAuthService(s, dep); err != nil {
		log.Fatalf("初始化服务失败: %v", err)
	}

	if err := s.Serve(); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}

func initialize() {
	errs.RegisterErrToCodeFilter()
	tracelog.RegisterTraceLogFilter()
	metricslog.RegisterMetricsMySQL()
	elapse.RegisterFilter()
	count.RegisterFilter()
	codec.UseOfficialJSON()
}
