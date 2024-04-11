package code

import "akita/panda-im/common/xcode"

//var (
//	ErrEmptyByUsernameOrPassword = manage.NewCodeError(10001, "用户名或密码不能为空")
//	ErrWrongByUsernameOrPassword = manage.NewCodeError(10002, "用户名或密码错误")
//	ErrLogin                     = manage.NewCodeError(10004, "登录失败")
//	ErrLogout                    = manage.NewCodeError(10005, "登出失败")
//	ErrAuthenticate              = manage.NewCodeError(10006, "认证失败") // 认证失败
//	ErrMobileEmpty               = manage.NewCodeError(10003, "手机号不能为空")
//	ErrNameEmpty                 = manage.NewCodeError(10004, "用户名不能为空")
//	ErrVerificationCodeEmpty     = manage.NewCodeError(10005, "验证码不能为空")
//	ErrPasswordEmpty             = manage.NewCodeError(10008, "密码不能为空")
//	ErrMobileFormatError         = manage.NewCodeError(10007, "手机号格式错误")
//	ErrVerificationLimitExceeded = manage.NewCodeError(10009, "验证码发送过于频繁，请稍后再试")
//	ErrVerificationCode          = manage.NewCodeError(10009, "验证码错误")
//	ErrVerificationCodeExpired   = manage.NewCodeError(10006, "验证码已过期")
//	ErrSendSmsFailed             = manage.NewCodeError(10007, "发送短信失败")
//	ErrMobileExist               = manage.NewCodeError(10008, "手机号已经注册")
//	ErrRegisterFailed            = manage.NewCodeError(10010, "注册失败")
//)

var (
	ErrEmptyByUsernameOrPassword = xcode.New(10001, "用户名或密码不能为空")
	ErrWrongByUsernameOrPassword = xcode.New(10002, "用户名或密码错误")
	ErrLogin                     = xcode.New(10004, "登录失败")
	ErrLogout                    = xcode.New(10005, "登出失败")
	ErrAuthenticate              = xcode.New(10006, "认证失败") // 认证失败
	ErrMobileEmpty               = xcode.New(10003, "手机号不能为空")
	ErrNameEmpty                 = xcode.New(10004, "用户名不能为空")
	ErrVerificationCodeEmpty     = xcode.New(10005, "验证码不能为空")
	ErrPasswordEmpty             = xcode.New(10008, "密码不能为空")
	ErrMobileFormatError         = xcode.New(10007, "手机号格式错误")
	ErrVerificationLimitExceeded = xcode.New(10009, "验证码发送过于频繁，请稍后再试")
	ErrVerificationCode          = xcode.New(10009, "验证码错误")
	ErrVerificationCodeNotExist  = xcode.New(10010, "请先获取验证码")
	ErrVerificationCodeExpired   = xcode.New(10006, "验证码已过期")
	ErrSendSmsFailed             = xcode.New(10007, "发送短信失败")
	ErrMobileExist               = xcode.New(10008, "手机号已注册")
	ErrRegisterFailed            = xcode.New(10010, "注册失败")
	ErrTokenInvalid              = xcode.New(10011, "无效令牌")
	ErrUserNotExist              = xcode.New(10012, "用户不存在")
)
