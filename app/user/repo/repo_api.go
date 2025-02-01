package repo

import (
	"context"

	"github.com/Andrew-M-C/trpc-go-demo/proto/user"
)

type AccountRepo interface {
	QueryAccountByUsername(ctx context.Context, username string) (*user.UserInfo, error)
}
