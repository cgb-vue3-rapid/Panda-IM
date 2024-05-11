package Auth

import (
	"akita/panda-im/common/constants"
	"akita/panda-im/common/util/rds_cache"
	"akita/panda-im/service/auth/internal/svc"
	"akita/panda-im/service/auth/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticateHandlerPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticateHandlerPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticateHandlerPostLogic {
	return &AuthenticateHandlerPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticateHandlerPostLogic) AuthenticateHandlerPost(req *types.AuthenticateRequest) (resp *types.AuthenticateResponse, err error) {
	// todo: add your logic here and delete this line
	//获取用户ID
	userId := l.ctx.Value(constants.UserId).(int64)
	refreshToken := l.ctx.Value(constants.RefreshToken).(string)

	_ = rds_cache.CacheSetNxEx(userId, refreshToken, constants.PrefixUserLoginCache, constants.JwtExpire, l.svcCtx.BizRedis)

	return &types.AuthenticateResponse{Message: "认证成功"}, nil
}
