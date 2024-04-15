package token_manager

//// ParseToken 解析 JWT token
//func ParseToken(tokenString string, secretKey string) (*Payloads, error) {
//	// 解析 token
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		// 检查 token 签名方法是否正确
//		//if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//		//	return nil, xcode.New(102, "无效令牌")
//		//}
//		return []byte(secretKey), nil
//	})
//
//	// 检查解析是否出错
//	if err != nil {
//		return nil, xcode.New(101, "解析令牌失败")
//	}
//
//	// 检查 token 是否有效
//	if !token.Valid {
//		return nil, xcode.New(102, "无效令牌")
//	}
//
//	// 解析 token 中的声明数据
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if !ok {
//		return nil, xcode.New(102, "无效令牌")
//	}
//
//	//构造 Payloads 结构体并返回
//	payloads := &Payloads{
//		UserID:   int64(claims["user_id"].(float64)),
//		NickName: claims["nick_name"].(string),
//		Role:     int32(claims["role"].(float64)),
//		// 解析其他字段
//	}
//	return payloads, nil
//}
