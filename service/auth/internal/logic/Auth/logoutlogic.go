package Auth

import (
	"akita/panda-im/common/constants"
	"akita/panda-im/common/util/rds_cache"
	"akita/panda-im/service/auth/code"
	"akita/panda-im/service/auth/internal/svc"
	"akita/panda-im/service/auth/internal/types"
	"akita/panda-im/service/user/rpc/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
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
	userId := l.ctx.Value(constants.UserId).(int64)

	_, err = l.svcCtx.UserRPC.Logout(l.ctx, &user.LogoutRequest{UserId: userId})
	if err != nil {
		return nil, code.ErrUserNotExist
	}

	// 删除缓存
	_ = rds_cache.DeleteTokenByUserID(userId, constants.PrefixUserLoginCache, l.svcCtx.BizRedis)

	return &types.LogoutResponse{Message: "登出成功"}, nil
}
