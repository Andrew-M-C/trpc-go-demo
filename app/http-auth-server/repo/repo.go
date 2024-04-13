package repo

import (
	"fmt"

	"github.com/Andrew-M-C/trpc-go-demo/app/http-auth-server/repo/config"
)

// Repo 表示 http-auth-server 的 repo 总集合
type Repo struct {
	config.GeneralConfigGetter

	dep Dependency
}

// Dependency 表示总体初始化依赖
type Dependency struct {
	GeneralConfigKeyName string
}

// NewRepo 创建一个 http-auth-server 的 repo 总集合
func NewRepo(dep Dependency) (*Repo, error) {
	r := &Repo{
		dep: dep,
	}
	initializers := []func() error{
		r.initializeGeneralConfigGetter,
	}
	for _, ini := range initializers {
		if err := ini(); err != nil {
			return nil, err
		}
	}

	return r, nil
}

func (r *Repo) initializeGeneralConfigGetter() error {
	dep := config.Dependency{
		GeneralConfigKeyName: r.dep.GeneralConfigKeyName,
	}
	if err := r.InitializeGeneralConfigGetter(dep); err != nil {
		return fmt.Errorf("初始化通用配置失败 (%w)", err)
	}
	return nil
}
