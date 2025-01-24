package repo

import (
	"context"

	"github.com/Andrew-M-C/trpc-go-demo/app/user/entity"
)

type AccountRepo interface {
	QueryAccountByUsername(ctx context.Context, username string) (*entity.Account, error)
}
