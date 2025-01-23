package main

import (
	"github.com/Andrew-M-C/trpc-go-utils/log"
	"github.com/Andrew-M-C/trpc-go-utils/recovery"
)

func main() {
	defer recovery.CatchPanic(recovery.WithErrorLog())

	a, err := newApplication()
	if err != nil {
		log.New().Text("初始化失败").Err(err).Error()
		return
	}

	if err := a.Serve(); err != nil {
		log.New().Text("启动服务失败").Err(err).Error()
		return
	}
}

type application interface {
	Serve() error
}
