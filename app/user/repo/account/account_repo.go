package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/Andrew-M-C/trpc-go-demo/app/user/entity"
	"gorm.io/gorm"
	"trpc.group/trpc-go/trpc-database/mysql"
)

// UserAccountRepository 用户账户仓库
type UserAccountRepository struct {
	dep Dependency
}

// Dependency 表示用户账户仓库初始化依赖
type Dependency struct {
	DBGetter   func(context.Context) (mysql.Client, error)
	GormGetter func(context.Context) (*gorm.DB, error)
}

// InitializeUserAccountRepository 初始化用户账户仓库
func (r *UserAccountRepository) InitializeUserAccountRepository(d Dependency) error {
	r.dep = d
	return nil
}

// QueryAccountByUsername 根据用户名查询用户账户
func (r *UserAccountRepository) QueryAccountByUsername(
	ctx context.Context, username string,
) (*entity.Account, error) {
	if r.dep.DBGetter != nil {
		return r.queryAccountByUsernameBySqlx(ctx, username)
	} else if r.dep.GormGetter != nil {
		return r.queryAccountByUsernameByGorm(ctx, username)
	}

	return nil, errors.New("实例未被正确初始化")
}

func (r *UserAccountRepository) queryAccountByUsernameBySqlx(
	ctx context.Context, username string,
) (*entity.Account, error) {
	db, err := r.dep.DBGetter(ctx)
	if err != nil {
		return nil, fmt.Errorf("获取 DB 失败 (%w)", err)
	}

	var res []userAccountItem
	query := fmt.Sprintf("SELECT * FROM %s WHERE username = ? AND delete_at_ms = 0 LIMIT 1", userAccountItem{}.TableName())
	if err := db.Select(ctx, &res, query, username); err != nil {
		return nil, fmt.Errorf("查询DB失败 (%w)", err)
	}
	if len(res) == 0 {
		return nil, nil // 使用双 nil 表示数据不存在
	}
	return res[0].toEntity(), nil
}

func (r *UserAccountRepository) queryAccountByUsernameByGorm(
	ctx context.Context, username string,
) (*entity.Account, error) {
	db, err := r.dep.GormGetter(ctx)
	if err != nil {
		return nil, fmt.Errorf("获取 DB 失败 (%w)", err)
	}
	db = db.Session(&gorm.Session{
		SkipDefaultTransaction: true,
	})

	var res userAccountItem
	err = db.Model(userAccountItem{}).
		Where("username = ?", username).
		Where("delete_at_ms = ?", 0).
		First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 使用双 nil 表示数据不存在
		}
		return nil, fmt.Errorf("查询DB失败 (%w)", err)
	}

	return res.toEntity(), nil
}
