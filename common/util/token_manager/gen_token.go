package token_manager

//func GenerateToken(opt TokenOptions) (*Token, error) {
//	// 检查参数合法性
//	if opt.AccessSecret == "" {
//		return nil, errors.New("AccessSecret cannot be empty")
//	}
//	if opt.AccessExpire <= 0 {
//		return nil, errors.New("AccessExpire must be greater than 0")
//	}
//
//	// 调用 createAccessToken 函数生成 AccessToken
//	accessToken, err := createAccessToken(opt)
//	if err != nil {
//		return nil, err
//	}
//
//	// 构造 Token 结构体
//	token := &Token{
//		AccessToken:  accessToken,
//		AccessExpire: time.Now().Unix() + opt.AccessExpire,
//	}
//
//	return token, nil
//}
//
//func createAccessToken(opt TokenOptions) (string, error) {
//	claims := jwt.MapClaims{}
//	claims["exp"] = time.Now().Unix() + opt.AccessExpire
//	claims["iat"] = time.Now().Unix()
//	claims["user_id"] = opt.Payloads.UserID
//	claims["nick_name"] = opt.Payloads.NickName
//	claims["role"] = opt.Payloads.Role
//	// 添加其他字段
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	return token.SignedString([]byte(opt.AccessSecret))
//}
