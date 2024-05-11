package repo

import (
	"fmt"

	"github.com/Andrew-M-C/trpc-go-demo/app/user/repo/account"
	"github.com/Andrew-M-C/trpc-go-utils/client/gorm"
	"github.com/Andrew-M-C/trpc-go-utils/client/sqlx"
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
	accDep := account.Dependency{}
	const useGorm = true
	if useGorm {
		accDep.GormGetter = gorm.ClientGetter(d.UserAccountDBClientName)
	} else {
		accDep.DBGetter = sqlx.ClientGetter(d.UserAccountDBClientName)
	}
	if err := r.InitializeUserAccountRepository(accDep); err != nil {
		return nil, fmt.Errorf("初始化用户仓库失败 (%w)", err)
	}

	// 成功返回
	return r, nil
}
