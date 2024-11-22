package main

import (
	"github.com/Andrew-M-C/trpc-go-demo/app/mq/service"
	"github.com/Andrew-M-C/trpc-go-demo/utils/filter/count"
	"github.com/Andrew-M-C/trpc-go-demo/utils/filter/elapse"
	"github.com/Andrew-M-C/trpc-go-utils/client/kafka"
	"github.com/Andrew-M-C/trpc-go-utils/codec"
	"github.com/Andrew-M-C/trpc-go-utils/errs"
	"github.com/Andrew-M-C/trpc-go-utils/log"
	metricslog "github.com/Andrew-M-C/trpc-go-utils/metrics/log"
	"github.com/Andrew-M-C/trpc-go-utils/recovery"
	"github.com/Andrew-M-C/trpc-go-utils/tracelog"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/server"
)

func main() {
	defer recovery.CatchPanic(recovery.WithErrorLog())

	s := initalize()
	d := service.Dependency{
		StringKafka: kafka.ClientGetter("kafka.amc.demo.string"),
		JSONKafka:   kafka.ClientGetter("kafka.amc.demo.json"),
	}
	if err := service.RegisterMQService(s, d); err != nil {
		log.New().Text("服务初始化失败").Err(err).Fatal()
	}
	if err := s.Serve(); err != nil {
		log.New().Text("服务启动失败").Err(err).Fatal()
	}
}

func initalize() *server.Server {
	errs.RegisterErrToCodeFilter()
	tracelog.RegisterTraceLogFilter()
	metricslog.RegisterMetricsMySQL()
	elapse.RegisterFilter()
	count.RegisterFilter()
	codec.UseOfficialJSON()

	return trpc.NewServer()
}
