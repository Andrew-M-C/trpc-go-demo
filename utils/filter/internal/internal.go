package internal

import (
	"context"
	"strings"

	"trpc.group/trpc-go/trpc-go/codec"
)

func GetCallee(ctx context.Context) string {
	if method := codec.Message(ctx).CalleeMethod(); method != "" {
		method = strings.TrimPrefix(method, "/")
		method = strings.ReplaceAll(method, "/", "-")
		return method
	}
	return "unknown"
}
