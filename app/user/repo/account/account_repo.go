package account

import (
	"context"
	"fmt"

	"github.com/Andrew-M-C/trpc-go-demo/app/user/entity"
	"trpc.group/trpc-go/trpc-database/mysql"
)

// UserAccountRepository 用户账户仓库
type UserAccountRepository struct {
	dep Dependency
}

// Dependency 表示用户账户仓库初始化依赖
type Dependency struct {
	DBGetter func(context.Context) (mysql.Client, error)
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
	db, err := r.dep.DBGetter(ctx)
	if err != nil {
		return nil, fmt.Errorf("获取 DB 失败 (%w)", err)
	}

	var res []userAccountItem
	query := fmt.Sprintf("SELECT * FROM %s WHERE username = ? LIMIT 1", userAccountItem{}.TableName())
	if err := db.Select(ctx, &res, query, username); err != nil {
		return nil, fmt.Errorf("查询DB失败 (%w)", err)
	}
	if len(res) == 0 {
		return nil, nil // 使用双 nil 表示数据不存在
	}
	return res[0].toEntity(), nil
}
