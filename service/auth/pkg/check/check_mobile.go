package check

//// CheckMobile 检查手机号格式是否正确
//func CheckMobile(mobile string) bool {
//	// 手机号正则表达式
//	// 中国大陆手机号格式：11位数字，以1开头
//	// 正则表达式中：^表示匹配字符串的开头，$表示匹配字符串的结尾
//	reg := `^1\d{10}$`
//
//	// 编译正则表达式
//	regex := regexp.MustCompile(reg)
//
//	// 使用正则表达式匹配手机号
//	return regex.MatchString(mobile)
//}
//
//// checkVerificationCode 函数用于检查验证码是否正确
//func checkVerificationCode(rds_cache *redis.Redis, mobile, c string) error {
//	// 从缓存中获取验证码
//	cacheCode, err := etActivationCacheByMobile(mobile, rds_cache)
//	if err != nil {
//		return applet_code.ErrVerificationCode
//	}
//	// 如果缓存中的验证码与请求中的验证码不一致，返回验证码错误
//	if cacheCode != c {
//		return applet_code.ErrVerificationCode
//	}
//	// 如果缓存中的验证码为空，说明验证码已过期
//	if cacheCode == "" {
//		return applet_code.ErrVerificationCodeExpired
//	}
//	return nil
//}
