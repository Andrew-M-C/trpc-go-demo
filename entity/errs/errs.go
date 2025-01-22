// Package errs 用于表示以 RPC 为主的错误信息, 同时也可以进行相关错误转换逻辑
package errs

import (
	"errors"

	"golang.org/x/exp/constraints"
	"trpc.group/trpc-go/trpc-go/errs"
)

//go:generate ./go_generate.sh

// New 返回带 code 和 message 的错误, 是 trpc 框架 errs.New 的封装
func New[T constraints.Integer](code T, message string) error {
	return errs.New(int(code), message)
}

// Raw 表示原生 errors.New 的封装
func Raw(message string) error {
	return errors.New(message)
}

// MARK: 其他标准 errors 包的封装

// 封装官方 errors.Join, 主要是方便调用方不用 import 两个包
func Join(errs ...error) error {
	return errors.Join(errs...)
}

// 封装官方 errors.Is, 主要是方便调用方不用 import 两个包
func Is(err, target error) bool {
	return errors.Is(err, target)
}
