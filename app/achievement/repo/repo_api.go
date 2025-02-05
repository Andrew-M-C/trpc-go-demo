package repo

import (
	"context"

	"github.com/Andrew-M-C/trpc-go-demo/app/achievement/entity/badge"
	"github.com/Andrew-M-C/trpc-go-demo/app/achievement/entity/reputation"
)

type Badge interface {
	GetUserBadges(ctx context.Context, userID string) ([]badge.Item, error)
}

type Reputation interface {
	GetUserReputation(ctx context.Context, userID string) (reputation.Reputation, error)
}
