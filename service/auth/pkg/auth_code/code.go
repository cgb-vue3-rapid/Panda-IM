package auth_code

import (
	"akita/panda-im/common/manage"
)

var (
	ErrMobileEmpty             = manage.NewCodeError(10001, "手机号不能为空")
	ErrNameEmpty               = manage.NewCodeError(10002, "用户名不能为空")
	ErrVerificationCodeEmpty   = manage.NewCodeError(10003, "验证码不能为空")
	ErrPasswordEmpty           = manage.NewCodeError(10004, "密码不能为空")
	ErrMobileFormatError       = manage.NewCodeError(10005, "手机号格式错误")
	ErrVerificationCodeExpired = manage.NewCodeError(10006, "验证码已过期")
	ErrVerificationCode        = manage.NewCodeError(10007, "验证码错误")
	ErrMobileExist             = manage.NewCodeError(10008, "手机号已经注册")
	ErrEncMobile               = manage.NewCodeError(10009, "加密手机号失败")
)
