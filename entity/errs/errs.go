// Package errs 定义整个应用的通用错误信息
package errs

import "trpc.group/trpc-go/trpc-go/errs"

var (
	// 认证相关错误

	PasswordError = errs.New(10001, "密码错误")
)
