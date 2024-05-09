package main

import (
	"context"
	"fmt"

	"github.com/Andrew-M-C/trpc-go-demo/app/http-auth-server/repo"
	"github.com/Andrew-M-C/trpc-go-demo/app/http-auth-server/service"
	"github.com/Andrew-M-C/trpc-go-demo/utils/filter/count"
	"github.com/Andrew-M-C/trpc-go-demo/utils/filter/elapse"
	"github.com/Andrew-M-C/trpc-go-utils/codec"
	etcdutil "github.com/Andrew-M-C/trpc-go-utils/config/etcd"
	"github.com/Andrew-M-C/trpc-go-utils/errs"
	metricslog "github.com/Andrew-M-C/trpc-go-utils/metrics/log"
	"github.com/Andrew-M-C/trpc-go-utils/recovery"
	"github.com/Andrew-M-C/trpc-go-utils/tracelog"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
	"trpc.group/trpc-go/trpc-go/server"

	_ "trpc.group/trpc-go/trpc-naming-polarismesh"
	_ "trpc.group/trpc-go/trpc-naming-polarismesh/selector"
)

func main() {
	defer recovery.CatchPanic(recovery.WithErrorLog())

	s, err := initServer()
	if err != nil {
		log.Fatal(err)
	}
	dep, err := initDependency()
	if err != nil {
		log.Fatal(err)
	}
	if err := service.RegisterAuthService(s, dep); err != nil {
		log.Fatalf("初始化服务失败: %v", err)
	}

	if err := s.Serve(); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}

func initDependency() (service.Dependency, error) {
	dep := service.Dependency{}

	r, err := repo.NewRepo(repo.Dependency{
		GeneralConfigKeyName: "/http-auth-server/general.json",
	})
	if err != nil {
		return dep, fmt.Errorf("初始化 repo 失败: %w", err)
	}

	dep.Repo = r
	return dep, nil
}

func initServer() (*server.Server, error) {
	// trpc 基础注册
	errs.RegisterErrToCodeFilter()
	tracelog.RegisterTraceLogFilter()
	metricslog.RegisterMetricsMySQL()
	elapse.RegisterFilter()
	count.RegisterFilter()
	codec.UseOfficialJSON()

	// 初始化 trpc
	s := trpc.NewServer()

	// 后置配置
	if err := etcdutil.RegisterClientProvider(
		context.Background(), s, "/http-auth-server/client.yaml",
	); err != nil {
		return nil, err
	}

	return s, nil
}
