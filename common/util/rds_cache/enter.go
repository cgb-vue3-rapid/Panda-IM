package rds_cache

import (
	"akita/panda-im/common/util/random_munber"
	"akita/panda-im/service/user/rpc/user"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
	"time"
)

/*
用户信息相关
*/

// CheckCacheByUserID 检查缓存是否存在
func CheckCacheByUserID(userID int64, prefix string, rds *redis.Redis) (string, error) {
	key := fmt.Sprintf(prefix, strconv.Itoa(int(userID)))
	token, err := rds.Get(key)
	if err != nil {
		return "", err
	}
	return token, nil
}

// DeleteTokenByUserID 删除缓存
func DeleteTokenByUserID(id int64, prefix string, rds *redis.Redis) error {
	key := fmt.Sprintf(prefix, strconv.Itoa(int(id)))
	_, err := rds.Del(key)
	if err != nil {
		return err
	}
	return nil
}

// CacheTokenByUserID 缓存用户Token
func CacheTokenByUserID(id int64, token string, prefix string, rds *redis.Redis, ex int) error {
	key := fmt.Sprintf(prefix, strconv.Itoa(int(id)))
	return rds.Setex(key, token, ex)
}

// CacheUserInfoByUserID 缓存用户信息
func CacheUserInfoByUserID(id int64, userInfo *user.UserInfoResponse, prefix string, rds *redis.Redis) error {
	// 构建 Redis 中哈希表的键名
	key := fmt.Sprintf(prefix, strconv.Itoa(int(id)))

	// 构建用户信息的 map
	userMap := map[string]interface{}{
		"UserID":     strconv.FormatInt(userInfo.UserId, 10),
		"Nickname":   userInfo.Nickname,
		"Mobile":     userInfo.Mobile,
		"Avatar":     userInfo.Avatar,
		"Role":       strconv.Itoa(int(userInfo.Role)),
		"Gender":     userInfo.Gender,
		"Addr":       userInfo.Addr,
		"CreatedAt":  userInfo.CreateAt,
		"UpdatedAt":  userInfo.UpdateAt,
		"DeletedAt":  userInfo.DeleteAt,
		"Is_Deleted": strconv.FormatBool(userInfo.IsDelete),
	}

	// 使用循环遍历 map，并将键值对存入 Redis 中
	for field, value := range userMap {
		err := rds.Hset(key, field, value.(string))
		if err != nil {
			return err
		}
	}

	return nil
}

/*
验证码相关
*/

// GetVerificationCountByMobile 获取今日验证码发送次数
func GetVerificationCountByMobile(mobile, prefix string, rds *redis.Redis) (int, error) {
	key := fmt.Sprintf(prefix, mobile)
	val, err := rds.Get(key)
	if err != nil {
		return 0, err
	}
	if len(val) == 0 {
		return 0, nil
	}

	return strconv.Atoi(val)
}

// IncrVerificationCountByMobile 增加验证码发送次数计数
func IncrVerificationCountByMobile(mobile, prefix string, rds *redis.Redis) error {
	key := fmt.Sprintf(prefix, mobile)
	_, err := rds.Incr(key)
	if err != nil {
		return err
	}

	return rds.Expireat(key, random_munber.EndOfDay(time.Now()).Unix())
}

// GetActivationCacheByMobile 获取验证码缓存
func GetActivationCacheByMobile(mobile, prefix string, rds *redis.Redis) (string, error) {
	key := fmt.Sprintf(prefix, mobile)
	return rds.Get(key)
}

// SaveActivationCacheByMobile 缓存验证码
func SaveActivationCacheByMobile(mobile, code, prefix string, rds *redis.Redis, ex int) error {
	key := fmt.Sprintf(prefix, mobile)
	return rds.Setex(key, code, ex)
}

// DelActivationCacheByMobile 删除验证码缓存
func DelActivationCacheByMobile(mobile, prefix string, rds *redis.Redis) error {
	key := fmt.Sprintf(prefix, mobile)
	_, err := rds.Del(key)
	return err
}
