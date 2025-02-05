// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"fmt"
	"github.com/Andrew-M-C/go.util/runtime/caller"
	repo2 "github.com/Andrew-M-C/trpc-go-demo/app/achievement/repo"
	"github.com/Andrew-M-C/trpc-go-demo/app/achievement/repo/badge"
	"github.com/Andrew-M-C/trpc-go-demo/app/achievement/repo/reputation"
	"github.com/Andrew-M-C/trpc-go-demo/app/achievement/service"
	service3 "github.com/Andrew-M-C/trpc-go-demo/app/http-auth-server/service"
	"github.com/Andrew-M-C/trpc-go-demo/app/user/repo"
	"github.com/Andrew-M-C/trpc-go-demo/app/user/repo/account"
	service2 "github.com/Andrew-M-C/trpc-go-demo/app/user/service"
	"github.com/Andrew-M-C/trpc-go-demo/proto/achieve"
	"github.com/Andrew-M-C/trpc-go-demo/proto/httpauth"
	"github.com/Andrew-M-C/trpc-go-demo/proto/user"
	"github.com/Andrew-M-C/trpc-go-demo/utils/filter/count"
	"github.com/Andrew-M-C/trpc-go-demo/utils/filter/elapse"
	"github.com/Andrew-M-C/trpc-go-utils/client/sqlx"
	"github.com/Andrew-M-C/trpc-go-utils/codec"
	"github.com/Andrew-M-C/trpc-go-utils/config/etcd"
	"github.com/Andrew-M-C/trpc-go-utils/errs"
	"github.com/Andrew-M-C/trpc-go-utils/log"
	log2 "github.com/Andrew-M-C/trpc-go-utils/metrics/log"
	"github.com/Andrew-M-C/trpc-go-utils/recovery"
	"github.com/google/wire"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/server"
)

import (
	_ "trpc.group/trpc-go/trpc-naming-polarismesh"
	_ "trpc.group/trpc-go/trpc-naming-polarismesh/selector"
)

// Injectors from wire.go:

func newApplication() (application, error) {
	server, err := provideTRPCService()
	if err != nil {
		return nil, err
	}
	accountRepo, err := provideAccountRepo()
	if err != nil {
		return nil, err
	}
	userService, err := provideUserService(accountRepo)
	if err != nil {
		return nil, err
	}
	badge, err := provideBadgeRepo()
	if err != nil {
		return nil, err
	}
	reputation, err := provideReputationRepo()
	if err != nil {
		return nil, err
	}
	achievementService, err := provideAchieveService(badge, reputation)
	if err != nil {
		return nil, err
	}
	authService, err := provideAuthService(userService, achievementService)
	if err != nil {
		return nil, err
	}
	mainApplication := provideApplication(server, authService)
	return mainApplication, nil
}

// wire.go:

var trpcSet = wire.NewSet(
	provideTRPCService,
)

func provideTRPCService() (s *server.Server, err error) {
	defer recovery.CatchPanic(recovery.WithCallback(func(_ context.Context, info any, stack []caller.Caller) {
		err = fmt.Errorf("panic info '%v', stack '%v'", info, log.ToJSON(stack))
	}),
	)
	errs.SetDigestDescription("错误码")
	errs.RegisterErrToCodeFilter()
	log.RegisterTraceLogFilter()
	log2.RegisterMetricsMySQL()
	elapse.RegisterFilter()
	count.RegisterFilter()
	codec.UseOfficialJSON()

	s = trpc.NewServer()

	// trpc 初始化后置
	const clientYAMLKey = "/demo/general/client.yaml"
	if err := etcd.RegisterClientProvider(context.Background(), s, clientYAMLKey); err != nil {
		return nil, err
	}

	return s, nil
}

var repoSet = wire.NewSet(
	provideAccountRepo,
	provideBadgeRepo,
	provideReputationRepo,
)

func provideAccountRepo() (repo.AccountRepo, error) {
	d := account.Dependency{
		DBGetter: sqlx.ClientGetter("mysql.demo.user.account"),
	}
	return account.New(d)
}

func provideBadgeRepo() (repo2.Badge, error) {
	d := badge.Dependency{
		DBGetter: sqlx.ClientGetter("mysql.demo.achievement.badge"),
	}
	return badge.New(d)
}

func provideReputationRepo() (repo2.Reputation, error) {
	d := reputation.Dependency{
		DBGetter: sqlx.ClientGetter("mysql.demo.achievement.reputation"),
	}
	return reputation.New(d)
}

var serviceSet = wire.NewSet(
	provideAchieveService,
	provideUserService,
	provideAuthService,
	provideApplication,
)

func provideAchieveService(
	badgeRepo repo2.Badge,
	reputationRepo repo2.Reputation,
) (achieve.AchievementService, error) {
	d := service.Dependency{
		BadgeRepo:      badgeRepo,
		ReputationRepo: reputationRepo,
	}
	return service.New(d)
}

func provideUserService(
	accountRepo repo.AccountRepo,
) (user.UserService, error) {
	d := service2.Dependency{
		AccountRepo: accountRepo,
	}
	return service2.New(d)
}

func provideAuthService(
	userService user.UserService,
	achieveService achieve.AchievementService,
) (httpauth.AuthService, error) {
	d := service3.Dependency{
		UserProxy:        wrappedUserClient{userService},
		AchievementProxy: wrappedAchieveClient{achieveService},
	}
	return service3.New(d)
}

func provideApplication(
	svc *server.Server,
	authSvc httpauth.AuthService,
) application {
	httpauth.RegisterAuthService(svc, authSvc)
	return svc
}
