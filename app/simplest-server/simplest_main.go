package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Andrew-M-C/trpc-go-demo/proto/simplest"
	"trpc.group/trpc-go/trpc-go"
)

func main() {
	s := trpc.NewServer()
	simplest.RegisterHelloWorldService(s, helloWorldImpl{})
	_ = s.Serve()
}

type helloWorldImpl struct{}

func (helloWorldImpl) Hello(_ context.Context, req *simplest.HelloRequest) (*simplest.HelloResponse, error) {
	rsp := &simplest.HelloResponse{}
	rsp.Response = fmt.Sprintf("%s to you, too", req.Greeting)
	rsp.Timestamp = float64(time.Now().UnixMilli()) / 1000
	return rsp, nil
}
