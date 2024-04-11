package logic

import (
	"akita/panda-im/common/encrypt"
	"akita/panda-im/service/user/rpc/code"
	"context"

	"akita/panda-im/service/user/rpc/internal/svc"
	"akita/panda-im/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByIDLogic {
	return &GetUserInfoByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByIDLogic) GetUserInfoByID(in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	u, err := l.svcCtx.UserModelDao.GetUserInfoByID(l.ctx, in.UserId)
	if err != nil {
		return nil, code.ErrUserNotExist
	}
	return &pb.UserInfoResponse{
		UserId:   u.ID,
		Nickname: u.NickName,
		Mobile:   u.Mobile,
		Avatar:   u.Avatar,
		Role:     u.Role,
		Gender:   string(u.Gender),
		Addr:     u.Addr,
		CreateAt: encrypt.FormatCreateTime(u.CreatedAt),
		UpdateAt: encrypt.FormatUpdateTime(u.UpdatedAt),
		DeleteAt: "",
		IsDelete: false,
	}, nil
}
