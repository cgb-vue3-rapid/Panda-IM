package logic

import (
	"akita/panda-im/service/user/rpc/code"
	"akita/panda-im/service/user/rpc/models/entity"
	"context"

	"akita/panda-im/service/user/rpc/internal/svc"
	"akita/panda-im/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Register rpc Login(LoginRequest) returns (LoginResponse);
func (l *RegisterLogic) Register(in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// todo: add your logic here and delete this line

	// 查询手机号是否已注册
	u, err := l.svcCtx.UserModelDao.FindByMobile(l.ctx, in.Mobile)

	// 手机号已注册
	if err == nil && u.ID > 0 {
		return nil, code.ErrMobileExist
	}

	// 创建用户
	userModel := entity.UserModel{
		Mobile:   in.Mobile,
		PassWord: in.Password,
	}

	err = l.svcCtx.UserModelDao.Insert(l.ctx, &userModel)
	if err != nil {
		return nil, code.ErrRegisterFail
	}

	return &pb.RegisterResponse{
		UserId:  userModel.ID,
		Message: "注册成功",
	}, nil
}
