package config

import (
	"context"
	"fmt"

	"github.com/Andrew-M-C/trpc-go-utils/config"
	"github.com/Andrew-M-C/trpc-go-utils/config/etcd"
)

// GeneralConfigGetter 获取 http-auth-server 的通用配置
type GeneralConfigGetter struct {
	data *configData
}

// Dependency GeneralConfigGetter 初始化依赖
type Dependency struct {
	GeneralConfigKeyName string
}

// InitializeGeneralConfigGetter 初始化 GeneralConfigGetter
func (c *GeneralConfigGetter) InitializeGeneralConfigGetter(dep Dependency) error {
	err := config.Bind(
		context.Background(), etcd.API{}, "json",
		dep.GeneralConfigKeyName, &c.data,
	)
	if err != nil {
		return fmt.Errorf("绑定配置失败 (%w)", err)
	}

	var yml *string
	err = config.Bind(
		context.Background(), etcd.API{}, "text",
		"/http-auth-server/client.yaml", &yml,
	)
	if err != nil {
		return fmt.Errorf("绑定测试配置失败 (%w)", err)
	}

	return nil
}

// GetEnv 获取运行环境
func (c *GeneralConfigGetter) GetEnv() string {
	if ptr := c.data; ptr != nil {
		return ptr.Env
	}
	return "unknown"
}

type configData struct {
	Env string `json:"env"`
}
