package account

import (
	"context"
	"fmt"

	dbutil "github.com/Andrew-M-C/go.util/mysql"
	"github.com/Andrew-M-C/trpc-go-demo/entity/errs"
	pb "github.com/Andrew-M-C/trpc-go-demo/proto/user"
	"github.com/go-playground/validator/v10"
	"trpc.group/trpc-go/trpc-database/mysql"
)

// Repo 用户账户仓库
type Repo struct {
	Dependency

	tzCorrector dbutil.TimezoneCorrector
}

// Dependency 表示用户账户仓库初始化依赖
type Dependency struct {
	DBGetter func(context.Context) (mysql.Client, error) `validate:"required"`
}

// New 新建帐户仓库
func New(d Dependency) (*Repo, error) {
	if err := validator.New().Struct(d); err != nil {
		return nil, err
	}
	tz, err := dbutil.NewTimezoneCorrectorBySelector(
		context.Background(), tzCorrector(d.DBGetter),
	)
	if err != nil {
		return nil, fmt.Errorf("new timezone corrector error (%w)", err)
	}
	r := &Repo{
		Dependency:  d,
		tzCorrector: tz,
	}
	return r, nil
}

// tzCorrector 用于实现 mysql.Selector 接口
type tzCorrector func(context.Context) (mysql.Client, error)

func (t tzCorrector) SelectContext(ctx context.Context, dest any, query string, args ...any) error {
	db, err := t(ctx)
	if err != nil {
		return fmt.Errorf("get DB error (%w)", err)
	}
	return db.Select(ctx, dest, query, args...)
}

func (r *Repo) getDB(ctx context.Context) (mysql.Client, error) {
	c, err := r.DBGetter(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errs.InternalError, err)
	}
	return c, nil
}

// QueryAccountByUsername 根据用户名查询用户账户
func (r *Repo) QueryAccountByUsername(
	ctx context.Context, username string,
) (*pb.UserInfo, error) {
	db, err := r.getDB(ctx)
	if err != nil {
		return nil, err
	}
	var res []userAccountItem
	const query = "SELECT id, username, password_hash, nickname FROM " +
		userAccountTableName +
		" WHERE username = ? AND delete_at_ms = 0 LIMIT 1"
	if err := db.Select(ctx, &res, query, username); err != nil {
		return nil, fmt.Errorf("查询DB失败 (%w)", err)
	}
	if len(res) == 0 {
		return nil, nil // 使用双 nil 表示数据不存在
	}
	return res[0].toPB(r.tzCorrector), nil
}
