package myError

import "errors"

// 自定义错误常量集
var (
	ERROR_USER_NOT_EXISTS = errors.New("用户不存在")
	ERROR_USER_EXISTS     = errors.New("用户已存在")
	ERROR_USER_PWD        = errors.New("密码不正确")
)
