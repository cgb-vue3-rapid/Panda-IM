package code

import "akita/panda-im/common/xcode"

var (
	// 用户不存在
	ErrUserNotExist = xcode.New(30001, "用户不存在")
	// 请传入用户ID
	ErrUserIdIsNil  = xcode.New(30002, "请传入用户ID")
	ErrUpdateFailed = xcode.New(10014, "更新失败")
)
