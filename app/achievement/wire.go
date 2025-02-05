//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"fmt"

	"github.com/Andrew-M-C/go.util/runtime/caller"
	"github.com/Andrew-M-C/trpc-go-demo/app/achievement/repo"
	"github.com/Andrew-M-C/trpc-go-demo/app/achievement/repo/badge"
	"github.com/Andrew-M-C/trpc-go-demo/app/achievement/repo/reputation"
	"github.com/Andrew-M-C/trpc-go-demo/app/achievement/service"
	"github.com/Andrew-M-C/trpc-go-demo/proto/achieve"
	"github.com/Andrew-M-C/trpc-go-demo/utils/filter/count"
	"github.com/Andrew-M-C/trpc-go-demo/utils/filter/elapse"
	"github.com/Andrew-M-C/trpc-go-utils/client/sqlx"
	"github.com/Andrew-M-C/trpc-go-utils/codec"
	"github.com/Andrew-M-C/trpc-go-utils/config/etcd"
	"github.com/Andrew-M-C/trpc-go-utils/errs"
	"github.com/Andrew-M-C/trpc-go-utils/log"
	metricslog "github.com/Andrew-M-C/trpc-go-utils/metrics/log"
	"github.com/Andrew-M-C/trpc-go-utils/recovery"
	"github.com/google/wire"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/server"

	_ "trpc.group/trpc-go/trpc-naming-polarismesh"
	_ "trpc.group/trpc-go/trpc-naming-polarismesh/selector"
)

func newApplication() (application, error) {
	wire.Build(
		trpcSet,
		repoSet,
		serviceSet,
	)
	return nil, nil
}

var trpcSet = wire.NewSet(
	provideTRPCService,
)

// ----------------
// MARK: TRPC 初始化

func provideTRPCService() (s *server.Server, err error) {
	defer recovery.CatchPanic(
		recovery.WithCallback(func(_ context.Context, info any, stack []caller.Caller) {
			err = fmt.Errorf("panic info '%v', stack '%v'", info, log.ToJSON(stack))
		}),
	)

	// trpc 初始化前置
	errs.SetDigestDescription("错误码")
	errs.RegisterErrToCodeFilter()
	log.RegisterTraceLogFilter()
	metricslog.RegisterMetricsMySQL()
	elapse.RegisterFilter()
	count.RegisterFilter()
	codec.UseOfficialJSON()

	// 初始化 trpc
	s = trpc.NewServer()

	// trpc 初始化后置
	const clientYAMLKey = "/demo/general/client.yaml"
	if err := etcd.RegisterClientProvider(context.Background(), s, clientYAMLKey); err != nil {
		return nil, err
	}

	return s, nil
}

// ----------------
// MARK: repo 初始化

var repoSet = wire.NewSet(
	provideBadgeRepo,
	provideReputationRepo,
)

func provideBadgeRepo() (repo.Badge, error) {
	d := badge.Dependency{
		DBGetter: sqlx.ClientGetter("mysql.demo.achievement.badge"),
	}
	return badge.New(d)
}

func provideReputationRepo() (repo.Reputation, error) {
	d := reputation.Dependency{
		DBGetter: sqlx.ClientGetter("mysql.demo.achievement.reputation"),
	}
	return reputation.New(d)
}

// ----------------
// MARK: service 初始化

var serviceSet = wire.NewSet(
	provideAchievementService,
)

func provideAchievementService(
	svr *server.Server,
	badgeRepo repo.Badge,
	reputationRepo repo.Reputation,
) (application, error) {
	d := service.Dependency{
		BadgeRepo:      badgeRepo,
		ReputationRepo: reputationRepo,
	}
	svcImpl, err := service.New(d)
	if err != nil {
		return nil, err
	}
	achieve.RegisterAchievementService(svr, svcImpl)
	return svr, nil
}
