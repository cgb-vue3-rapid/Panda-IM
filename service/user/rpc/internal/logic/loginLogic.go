package logic

import (
	"akita/panda-im/service/user/rpc/code"
	"akita/panda-im/service/user/rpc/internal/svc"
	"akita/panda-im/service/user/rpc/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	// todo: add your logic here and delete this line
	// 1. 手机号是否存在
	u, err := l.svcCtx.UserDao.FindByMobile(l.ctx, in.Mobile)
	if err != nil {
		return nil, code.ErrMobileExist
	}

	// 2. 密码是否正确
	if u.PassWord != in.Password {
		return nil, code.ErrPassword
	}
	return &pb.LoginResponse{
		UserId:   u.ID,
		Nickname: u.NickName,
		Role:     u.Role,
	}, nil
}
