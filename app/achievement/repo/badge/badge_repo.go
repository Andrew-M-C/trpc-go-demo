package badge

import (
	"context"
	"fmt"
	"time"

	"github.com/Andrew-M-C/trpc-go-demo/app/achievement/entity/badge"
	"github.com/Andrew-M-C/trpc-go-demo/entity/errs"
	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
	"trpc.group/trpc-go/trpc-database/mysql"
)

// Repo 用户徽章仓库
type Repo struct {
	Dependency
}

// Dependency 表示用户账户仓库初始化依赖
type Dependency struct {
	DBGetter func(context.Context) (mysql.Client, error) `validate:"required"`
}

func New(d Dependency) (*Repo, error) {
	if err := validator.New().Struct(d); err != nil {
		return nil, err
	}
	return &Repo{d}, nil
}

func (r *Repo) getDB(ctx context.Context) (mysql.Client, error) {
	c, err := r.DBGetter(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.InternalError, err)
	}
	return c, nil
}

func (r *Repo) GetUserBadges(ctx context.Context, userID string) ([]badge.Item, error) {
	db, err := r.getDB(ctx)
	if err != nil {
		return nil, err
	}

	const statement = "SELECT * FROM " + tableName + " WHERE user_id = ?"
	var items []badgeItem
	if err := db.Select(ctx, &items, statement, userID); err != nil {
		return nil, fmt.Errorf("%w: %v", errs.DBError, err)
	}

	res := lo.Map(items, func(item badgeItem, _ int) badge.Item {
		return badge.Item{
			Name:       item.BadgeName,
			CreateTime: time.UnixMilli(item.CreateMsec),
		}
	})
	return res, nil
}

// --------
// MARK: table 定义

const tableName = "t_trpc_demo_user_badge"

type badgeItem struct {
	ID         uint32 `db:"id"`
	UserID     string `db:"user_id"`
	BadgeName  string `db:"badge_name"`
	CreateMsec int64  `db:"create_msec"`
}
