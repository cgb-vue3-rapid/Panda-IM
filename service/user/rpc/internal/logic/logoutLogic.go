package logic

import (
	"akita/panda-im/service/user/rpc/code"
	"context"

	"akita/panda-im/service/user/rpc/internal/svc"
	"akita/panda-im/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	// todo: add your logic here and delete this line

	// err == nil 表示用户存在
	_, err := l.svcCtx.UserModelDao.FindByID(l.ctx, in.UserId)
	if err != nil {
		return nil, code.ErrUserNotExist
	}

	return &pb.LogoutResponse{}, nil
}
