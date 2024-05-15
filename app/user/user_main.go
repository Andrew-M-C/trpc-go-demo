package main

import (
	"context"

	"github.com/Andrew-M-C/trpc-go-demo/app/user/repo"
	"github.com/Andrew-M-C/trpc-go-demo/app/user/service"
	"github.com/Andrew-M-C/trpc-go-demo/utils/filter/count"
	"github.com/Andrew-M-C/trpc-go-demo/utils/filter/elapse"
	"github.com/Andrew-M-C/trpc-go-utils/codec"
	etcdutil "github.com/Andrew-M-C/trpc-go-utils/config/etcd"
	"github.com/Andrew-M-C/trpc-go-utils/errs"
	"github.com/Andrew-M-C/trpc-go-utils/plugin"
	"github.com/Andrew-M-C/trpc-go-utils/tracelog"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
	"trpc.group/trpc-go/trpc-go/server"

	_ "trpc.group/trpc-go/trpc-naming-polarismesh"
)

func main() {
	s, err := initServer()
	if err != nil {
		log.Fatal(err)
	}

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

func initServer() (*server.Server, error) {
	// 获取 trpc_go.yaml 中的配置
	clientYamlConf := struct {
		Key string `yaml:"key"`
	}{}
	plugin.Bind("config", "client_yaml", &clientYamlConf)

	// trpc 基础注册
	errs.RegisterErrToCodeFilter()
	tracelog.RegisterTraceLogFilter()
	elapse.RegisterFilter()
	count.RegisterFilter()
	codec.UseOfficialJSON()

	// 初始化 trpc
	s := trpc.NewServer()

	// 后置配置
	if k := clientYamlConf.Key; k != "" {
		if err := etcdutil.RegisterClientProvider(context.Background(), s, k); err != nil {
			return nil, err
		}
	}

	return s, nil
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
