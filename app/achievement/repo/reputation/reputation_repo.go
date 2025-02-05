package reputation

import (
	"context"
	"fmt"

	"github.com/Andrew-M-C/trpc-go-demo/app/achievement/entity/reputation"
	"github.com/Andrew-M-C/trpc-go-demo/entity/errs"
	"github.com/go-playground/validator/v10"
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

func (r *Repo) GetUserReputation(ctx context.Context, userID string) (res reputation.Reputation, err error) {
	db, err := r.getDB(ctx)
	if err != nil {
		return res, err
	}

	const statement = "SELECT * FROM " + tableName + " WHERE user_id = ? LIMIT 1"
	var items []reputationItem
	if err := db.Select(ctx, &items, statement, userID); err != nil {
		return res, fmt.Errorf("%w: %v", errs.DBError, err)
	}

	if len(items) == 0 {
		return res, nil // 返回全空
	}
	res.ExpPoints = items[0].ExpPoints
	res.Rank = uint32(items[0].RankLevel)
	return res, nil
}

// --------
// MARK: table 定义

const tableName = "t_trpc_demo_user_reputation"

type reputationItem struct {
	ID        uint32 `db:"id"`
	UserID    string `db:"exp_points"`
	ExpPoints uint32 `db:"user_id"`
	RankLevel uint16 `db:"rank_level"`
}
