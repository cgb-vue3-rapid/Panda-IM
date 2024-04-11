package Auth

import (
	"akita/panda-im/service/auth/code"
	"akita/panda-im/service/user/rpc/user"
	"context"
	"encoding/json"
	"github.com/spf13/cast"

	"akita/panda-im/service/auth/internal/svc"
	"akita/panda-im/service/auth/internal/types"

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
	userID, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, code.ErrUserNotExist
	}

	u, err := l.svcCtx.UserRPC.GetUserInfoByID(l.ctx, &user.UserInfoRequest{UserId: userID})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResponse{
		UserID:     u.UserId,
		Nickname:   u.Nickname,
		Mobile:     u.Mobile,
		Avatar:     u.Avatar,
		Role:       u.Role,
		Gender:     cast.ToInt(u.Gender),
		Addr:       u.Addr,
		CreatedAt:  u.CreateAt,
		UpdatedAt:  u.UpdateAt,
		DeletedAt:  u.DeleteAt,
		Is_Deleted: u.IsDelete,
		Message:    "获取用户信息成功",
	}, nil
}
