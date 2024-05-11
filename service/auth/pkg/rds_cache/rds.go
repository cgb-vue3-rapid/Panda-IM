package rds_cache

import (
	"akita/panda-im/common/token_manager"
	"akita/panda-im/service/auth/internal/svc"
	"fmt"
)

// SetNxExTokenCache 设置token缓存
func SetNxExTokenCache(prefix string, parseToken *token_manager.Payloads, token string, svc *svc.ServiceContext) (err error) {
	key := fmt.Sprintf(prefix, parseToken.UserID)
	_, err = svc.BizRedis.SetnxEx(key, token, int(svc.Config.Auth.AccessExpire))
	if err != nil {
		return err
	}
	return nil
}

// DelTokenCache 删除token缓存
func DelTokenCache(prefix string, parseToken *token_manager.Payloads, svc *svc.ServiceContext) (err error) {
	key := fmt.Sprintf(prefix, parseToken.UserID)
	_, err = svc.BizRedis.Del(key)
	if err != nil {
		return err
	}
	return nil
}

// GetCache 获取token缓存
func GetCache(prefix string, parseToken *token_manager.Payloads, token string, svc *svc.ServiceContext) (cache string, err error) {
	key := fmt.Sprintf(prefix, parseToken.UserID)
	cache, err = svc.BizRedis.Get(key)
	if err != nil {
		return "", err
	}
	return cache, nil
}
