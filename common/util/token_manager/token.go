package token_manager

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"
)

//type (
//	TokenOptions struct {
//		AccessSecret string
//		AccessExpire int64
//		Payloads
//	}
//
//	Token struct {
//		AccessToken  string `json:"access_token"`
//		AccessExpire int64  `json:"access_expire"`
//	}
//
//	Payloads struct {
//		UserID   int64  `json:"user_id"`
//		NickName string `json:"nick_name"`
//		Role     int32  `json:"role"`
//		// 添加其他字段
//	}
//)

type (
	TokenOptions struct {
		AccessSecret string
		AccessExpire int64
		Fields       map[string]interface{}
	}

	Token struct {
		AccessToken  string `json:"access_token"`
		AccessExpire int64  `json:"access_expire"`
	}
)

func BuildTokens(opt TokenOptions) (Token, error) {
	var token Token
	now := time.Now().Add(-time.Minute).Unix()
	accessToken, err := genToken(now, opt.AccessSecret, opt.Fields, opt.AccessExpire)
	if err != nil {
		return token, err
	}
	token.AccessToken = accessToken
	token.AccessExpire = now + opt.AccessExpire

	return token, nil
}

func genToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}

func ParseToken(tokenString, secretKey string) (map[string]interface{}, error) {
	// 解析 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 检查 token 签名方法是否正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("无效令牌")
		}
		return []byte(secretKey), nil
	})

	// 检查解析是否出错
	if err != nil {
		return nil, err
	}

	// 检查 token 是否有效
	if !token.Valid {
		return nil, err
	}

	// 解析 token 中的声明数据
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	// 返回声明数据
	return claims, nil
}
