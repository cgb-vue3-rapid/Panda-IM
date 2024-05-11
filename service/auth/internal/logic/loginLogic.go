package logic

import (
	"akita/panda-im/common/constants"
	"akita/panda-im/common/util/encrypt"
	"akita/panda-im/common/util/rds_cache"
	"akita/panda-im/common/util/token_manager"
	"akita/panda-im/common/xcode"
	"akita/panda-im/service/auth/code"
	"akita/panda-im/service/auth/internal/svc"
	"akita/panda-im/service/auth/internal/types"
	"akita/panda-im/service/user/rpc/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	// 判断请求参数是否合法
	if req.Mobile == "" {
		return nil, code.ErrMobileEmpty
	}
	if req.Password == "" {
		return nil, code.ErrPasswordEmpty
	}

	// 对手机号进行加密处理
	mobile, err := encrypt.EncMobile(req.Mobile)
	if err != nil {
		logx.Errorf("手机号加密失败: %s error: %v", req.Mobile, err)
		return nil, code.ErrLogin
	}

	// 如果手机号没注册，就对密码进行加密处理
	password := encrypt.EncPassword(req.Password)

	// 登录
	u, err := l.svcCtx.UserRPC.Login(l.ctx, &user.LoginRequest{
		Mobile:   mobile,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	//生成访问令牌
	AccessToken, RefreshToken, err := token_manager.GenToken(u.UserId, u.Nickname, l.svcCtx.Config.Token.AccessSecret, l.svcCtx.Config.Token.RefreshSecret, u.Role)

	if err != nil {
		return nil, xcode.TokenGenerateErr
	}

	_ = rds_cache.CacheTokenByUserID(u.UserId, RefreshToken, constants.PrefixUserLoginCache, l.svcCtx.BizRedis, int(time.Hour*constants.JwtExpire*3/time.Second))

	return &types.LoginResponse{
		JWT: types.JWT{
			AccessToken:  AccessToken,
			RefreshToken: AccessToken + " " + RefreshToken,
		},
		Message: "登录成功",
	}, nil
}
