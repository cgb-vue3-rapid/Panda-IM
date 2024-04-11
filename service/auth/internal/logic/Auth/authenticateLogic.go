package Auth

import (
	"akita/panda-im/service/auth/code"
	"akita/panda-im/service/auth/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"

	"akita/panda-im/service/auth/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticateLogic {
	return &AuthenticateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticateLogic) Authenticate() (resp *types.AuthenticateResponse, err error) {
	// todo: add your logic here and delete this line
	//获取用户ID
	userID, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, code.ErrTokenInvalid
	}

	// 检查缓存是否存在
	token, err := CheckCache(userID, l.svcCtx.BizRedis)
	if err != nil {
		return nil, code.ErrAuthenticate
	}

	// 判断缓存是否存在
	if token != "" {
		logx.Error(context.Background(), "token已存在认证失败")
		return nil, code.ErrAuthenticate
	}

	return &types.AuthenticateResponse{Message: "认证成功"}, nil
}

// CheckCache 检查缓存是否存在
func CheckCache(userID int64, rds *redis.Redis) (string, error) {
	key := fmt.Sprintf(prefixTokenCache, strconv.Itoa(int(userID)))
	token, err := rds.Get(key)
	if err != nil {
		return "", err
	}
	return token, nil
}
