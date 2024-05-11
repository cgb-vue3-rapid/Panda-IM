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
	u, err := l.svcCtx.UserDao.FindByMobile(l.ctx, in.Mobile)

	// 手机号已注册
	if err == nil && u.ID > 0 {
		logx.Errorf("[Register] mobile: %s, already exist", in.Mobile)
		return nil, code.ErrMobileExist
	}

	// 创建用户
	userModel := entity.UserModel{
		Mobile:   in.Mobile,
		PassWord: in.Password,
	}

	// 插入数据库
	err = l.svcCtx.UserDao.Insert(l.ctx, &userModel)
	if err != nil {
		logx.Errorf("[Register] insert user fail: %v", err)
		return nil, code.ErrRegisterFail
	}

	// 初始化用户配置表
	userConfig := &entity.UserConfModel{
		UserId:           userModel.ID,
		Oline:            false,
		FriendsOnline:    false,
		AllSounds:        true,
		SecureLink:       true,
		SavePwd:          false,
		SearchUser:       1,
		VerificationType: 2,
	}

	// 创建用户配置
	if err := l.svcCtx.Orm.WithContext(l.ctx).Create(userConfig).Error; err != nil {
		logx.Errorf("[Register] create user config fail: %v", err)
		return nil, code.ErrRegisterFail
	}

	return &pb.RegisterResponse{
		UserId:  userModel.ID,
		Message: "注册成功",
	}, nil
}
