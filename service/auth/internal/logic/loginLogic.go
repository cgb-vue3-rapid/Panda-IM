package logic

import (
	"akita/panda-im/common/encrypt"
	"akita/panda-im/common/token_manager"
	"akita/panda-im/service/auth/code"
	"akita/panda-im/service/auth/internal/svc"
	"akita/panda-im/service/auth/internal/types"
	"akita/panda-im/service/user/rpc/user"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	prefixTokenCache = "panda:user:login:id:%s"
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
		return nil, code.ErrRegisterFailed
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

	////  生成token
	//token, err := token_manager.GenerateToken(token_manager.TokenOptions{
	//	AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
	//	AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
	//	Payloads: token_manager.Payloads{
	//		UserID:   u.UserId,
	//		NickName: u.Nickname,
	//		Role:     u.Role,
	//	},
	//})

	// 生成访问令牌,将用户ID作为JWT的payload
	token, err := token_manager.BuildTokens(token_manager.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId":   u.UserId,
			"role":     u.Role,
			"nickname": u.Nickname,
		},
	})

	_ = CacheToken(u.UserId, token.AccessToken, l.svcCtx.BizRedis)

	return &types.LoginResponse{Token: types.Token{
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
	}, Message: "登录成功"}, nil
}

// CacheToken 缓存验证码
func CacheToken(id int64, token string, rds *redis.Redis) error {
	key := fmt.Sprintf(prefixTokenCache, strconv.Itoa(int(id)))
	return rds.Setex(key, token, expireActivation)
}
