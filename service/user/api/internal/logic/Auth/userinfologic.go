package Auth

import (
	"akita/panda-im/common/constants"
	"akita/panda-im/common/util/encrypt"
	"akita/panda-im/common/util/rds_cache"
	"akita/panda-im/service/user/api/internal/svc"
	"akita/panda-im/service/user/api/internal/types"
	"akita/panda-im/service/user/rpc/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	//获取用户ID
	userId := l.ctx.Value(constants.UserId).(int64)

	u, err := l.svcCtx.UserRPC.GetUserInfoByID(l.ctx, &user.UserInfoRequest{UserId: userId})
	if err != nil {
		return nil, err
	}

	mobile, err := encrypt.DecMobile(u.Mobile)
	if err != nil {
		logx.Errorf("解密手机号失败: %v", err)
	}

	u.Mobile = mobile
	// 缓存用户信息
	cacheErr := rds_cache.CacheUserInfoByUserID(u.UserId, u, constants.PrefixUserInfoCache, l.svcCtx.BizRedis)
	if cacheErr != nil {
		logx.Errorf("缓存用户信息失败: %v", cacheErr)
	}

	return &types.UserInfoResponse{
		UserID:   u.UserId,
		Nickname: u.Nickname,
		Mobile:   u.Mobile,
		Avatar:   u.Avatar,
		Abstract: u.Abstract,
		Gender:   u.Gender,
		Addr:     u.Addr,
		Message:  "获取用户信息成功",
	}, nil
}
