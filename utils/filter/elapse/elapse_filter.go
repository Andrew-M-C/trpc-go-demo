package elapse

import (
	"context"
	"fmt"
	"time"

	"github.com/Andrew-M-C/trpc-go-demo/utils/filter/internal"
	"trpc.group/trpc-go/trpc-go/filter"
	"trpc.group/trpc-go/trpc-go/log"
	"trpc.group/trpc-go/trpc-go/metrics"
)

const (
	// FilterName 定义 filter 名称
	FilterName = "elapse"

	mPrefix = "utils.filter.elapse."
)

// RegisterFilter 注册耗时统计 filter
func RegisterFilter() {
	filter.Register(FilterName, serverFilter, clientFilter)
}

func serverFilter(ctx context.Context, req any, next filter.ServerHandleFunc) (any, error) {
	start := time.Now()
	rsp, err := next(ctx, req)
	ela := time.Since(start)

	method := internal.GetCallee(ctx)
	key := fmt.Sprintf("%s.server.%s.usec", mPrefix, method)
	log.DebugContextf(ctx, "done server %s, ela %v", method, ela)
	metrics.SetGauge(key, float64(ela.Microseconds()))

	return rsp, err
}

func clientFilter(ctx context.Context, req, rsp any, next filter.ClientHandleFunc) error {
	start := time.Now()
	err := next(ctx, req, rsp)
	ela := time.Since(start)

	method := internal.GetCallee(ctx)
	key := fmt.Sprintf("%sclient.%s.usec", mPrefix, method)
	log.DebugContextf(ctx, "done client %s, ela %v", ela)
	metrics.SetGauge(key, float64(ela.Microseconds()))

	return err
}
