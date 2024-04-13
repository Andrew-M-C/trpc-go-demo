package count

import (
	"context"

	"github.com/Andrew-M-C/trpc-go-demo/utils/filter/internal"
	"trpc.group/trpc-go/trpc-go/filter"
	"trpc.group/trpc-go/trpc-go/metrics"
)

const (
	// FilterName 定义 filter 名称
	FilterName = "count"

	mPrefix = "utils.filter.count."
)

// RegisterFilter 注册调用次数 filter
func RegisterFilter() {
	filter.Register(FilterName, serverFilter, clientFilter)
}

func serverFilter(ctx context.Context, req any, next filter.ServerHandleFunc) (any, error) {
	rsp, err := next(ctx, req)

	metrics.Counter(mPrefix + "server." + internal.GetCallee(ctx) + ".cnt").Incr()
	if err != nil {
		metrics.Counter(mPrefix + "server." + internal.GetCallee(ctx) + ".fail").Incr()
	} else {
		metrics.Counter(mPrefix + "server." + internal.GetCallee(ctx) + ".succ").Incr()
	}
	return rsp, err
}

func clientFilter(ctx context.Context, req, rsp any, next filter.ClientHandleFunc) error {
	err := next(ctx, req, rsp)

	metrics.Counter(mPrefix + "client." + internal.GetCallee(ctx) + ".cnt").Incr()
	if err != nil {
		metrics.Counter(mPrefix + "client." + internal.GetCallee(ctx) + ".fail").Incr()
	} else {
		metrics.Counter(mPrefix + "client." + internal.GetCallee(ctx) + ".succ").Incr()
	}
	return err
}
