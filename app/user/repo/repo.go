package repo

import (
	"context"
	"fmt"

	"github.com/Andrew-M-C/trpc-go-demo/app/user/repo/account"
	"trpc.group/trpc-go/trpc-database/mysql"
)

// Repo 表示 user 服务所需的 repo 依赖全集
type Repo struct {
	account.UserAccountRepository
}

// Dependency 表示 repo 初始化参数
type Dependency struct {
	UserAccountDBClientName string
}

// NewRepo 新建 user 服务所需的 repo 依赖全集
func NewRepo(d Dependency) (*Repo, error) {
	r := &Repo{}

	// 初始化用户仓库
	accDep := account.Dependency{
		DBGetter: func(ctx context.Context) (mysql.Client, error) {
			return mysql.NewUnsafeClient(d.UserAccountDBClientName), nil
		},
	}
	if err := r.InitializeUserAccountRepository(accDep); err != nil {
		return nil, fmt.Errorf("初始化用户仓库失败 (%w)", err)
	}

	// 成功返回
	return r, nil
}
