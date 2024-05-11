package token_manager

import (
	"akita/panda-im/common/constants"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// CustomClaims /**
// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	UserID               int64  `json:"user_id"`
	Nickname             string `json:"nickname"`
	Role                 int32  `json:"role"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// GenToken 生成JWT
func GenToken(userId int64, nickname, access, refresh string, role int32) (string, string, error) {
	accessSecret := []byte(access)
	refreshSecret := []byte(refresh)
	// 创建一个我们自己声明的数据
	// accessToken 的数据
	accessClaims := CustomClaims{
		userId,
		nickname, // 自定义字段
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * constants.JwtExpire)), // 定义过期时间
			Issuer:    constants.OrganizationName,                                            // 签发人
			ID:        constants.OrganizationName,                                            // 编号
		},
	}
	logx.Infof(accessClaims.ExpiresAt.String())
	// refreshToken 的数据
	refreshClaims := CustomClaims{
		userId,
		nickname, // 自定义字段
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * constants.JwtExpire * 3)), // 定义过期时间
			Issuer:    constants.OrganizationName,                                                // 签发人
			ID:        constants.OrganizationName,                                                // 编号
		},
	}
	logx.Infof(refreshClaims.ExpiresAt.String())
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	// 生成签名字符串
	accessTokenSigned, err := accessToken.SignedString(accessSecret)
	if err != nil {
		fmt.Println("获取Token失败，Secret错误")
		return "", "", err
	}
	refreshTokenSigned, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		fmt.Println("获取Token失败，Secret错误")
		return "", "", err
	}
	return accessTokenSigned, refreshTokenSigned, nil
}

func ParseToken(accessTokenString, refreshTokenString, access, refresh string) (*CustomClaims, bool, error) {
	accessSecret := []byte(access)
	refreshSecret := []byte(refresh)

	accessToken, err := jwt.ParseWithClaims(accessTokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		logx.Infof("accessTokenString: %v", accessTokenString)
		return accessSecret, nil
	})

	if claims, ok := accessToken.Claims.(*CustomClaims); ok {
		if accessToken.Valid {
			return claims, false, nil
		} else if !claims.Expired() {
			logx.Infof("accessToken过期续期token")
			return claims, true, nil
		}
	}

	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		logx.Infof("refreshTokenString: %v", refreshTokenString)
		return refreshSecret, nil
	})

	if err != nil {
		logx.Errorf("refreshToken解析失败：%v", err)
		return nil, false, err
	}

	if claims, ok := refreshToken.Claims.(*CustomClaims); ok && refreshToken.Valid {
		if !claims.Expired() {
			// refreshToken没过期刷新token
			logx.Info("refreshToken没过期续期token")
			return claims, true, nil
		}
	}
	logx.Errorf("两个token全部过期，需要重新登陆")
	return nil, false, errors.New("invalid token")
}

func (c *CustomClaims) Expired() bool {
	logx.Infof("c.ExpiresAt: %v", c.ExpiresAt)
	return c.ExpiresAt.Unix() < time.Now().Unix()
}
