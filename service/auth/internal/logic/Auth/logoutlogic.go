package Auth

import (
	"akita/panda-im/service/auth/code"
	"akita/panda-im/service/user/rpc/user"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"

	"akita/panda-im/service/auth/internal/svc"
	"akita/panda-im/service/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	prefixTokenCache = "panda:user:login:id:%s"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() (resp *types.LogoutResponse, err error) {
	//获取用户ID
	userID, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, code.ErrTokenInvalid
	}

	_, err = l.svcCtx.UserRPC.Logout(l.ctx, &user.LogoutRequest{UserId: userID})
	if err != nil {
		return nil, code.ErrUserNotExist
	}

	// 删除缓存
	_ = deleteToken(userID, l.svcCtx.BizRedis)

	return &types.LogoutResponse{Message: "登出成功"}, nil
}

// deleteToken 删除缓存
func deleteToken(id int64, rds *redis.Redis) error {
	key := fmt.Sprintf(prefixTokenCache, strconv.Itoa(int(id)))
	_, err := rds.Del(key)
	if err != nil {
		return err
	}
	return nil
}
