package logic

import (
	"akita/panda-im/common/constants"
	"akita/panda-im/common/util/encrypt"
	"akita/panda-im/common/util/rds_cache"
	"akita/panda-im/service/auth/code"
	"akita/panda-im/service/auth/internal/svc"
	"akita/panda-im/service/auth/internal/types"
	"akita/panda-im/service/user/rpc/user"
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"regexp"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// 去除请求中的空白字符
	req.Nickname = strings.TrimSpace(req.Nickname)
	req.Mobile = strings.TrimSpace(req.Mobile)
	req.Password = strings.TrimSpace(req.Password)
	req.VerificationCode = strings.TrimSpace(req.VerificationCode)

	// 检查用户名是否为空
	if len(req.Nickname) == 0 {
		return nil, code.ErrNameEmpty
	}

	// 检查手机号是否为空
	if len(req.Mobile) == 0 {
		return nil, code.ErrMobileEmpty
	}

	// 检查手机号格式
	if !CheckMobile(req.Mobile) {
		return nil, code.ErrMobileFormatError
	}

	// 检查验证码是否为空
	if len(req.VerificationCode) == 0 {
		return nil, code.ErrVerificationCodeEmpty
	}

	// 检查密码是否为空
	if len(req.Password) == 0 {
		return nil, code.ErrPasswordEmpty
	}

	// 检查验证码是否正确
	err = checkVerificationCode(l.svcCtx.BizRedis, req.Mobile, req.VerificationCode)
	if err != nil {
		return nil, err
	}

	// 对手机号进行加密处理
	mobile, err := encrypt.EncMobile(req.Mobile)
	if err != nil {
		logx.Errorf("手机号加密失败: %s error: %v", req.Mobile, err)
		return nil, code.ErrRegisterFailed
	}

	// 如果手机号没注册，就对密码进行加密处理
	password := encrypt.EncPassword(req.Password)

	registerResp, err := l.svcCtx.UserRPC.Register(l.ctx, &user.RegisterRequest{
		Nickname: req.Nickname,
		Mobile:   mobile,
		Password: password,
	})

	if err != nil {
		_ = rds_cache.DelActivationCacheByMobile(req.Mobile, constants.PrefixActivation, l.svcCtx.BizRedis)

		logx.Errorf("注册失败: %s error: %v", req.Mobile, err)
		return nil, err
	}

	//  删除验证码缓存
	_ = rds_cache.DelActivationCacheByMobile(req.Mobile, constants.PrefixActivation, l.svcCtx.BizRedis)

	return &types.RegisterResponse{
		UserID:  registerResp.UserId,
		Message: registerResp.Message,
	}, nil
}

// checkVerificationCode 函数用于检查验证码是否正确
func checkVerificationCode(rds *redis.Redis, mobile, c string) error {
	// 从缓存中获取验证码
	cacheCode, err := rds_cache.GetActivationCacheByMobile(mobile, constants.PrefixActivation, rds)
	if err != nil {
		return code.ErrVerificationCodeNotExist
	}
	// 如果缓存中的验证码为空，说明验证码已过期
	if cacheCode == "" {
		return code.ErrVerificationCodeExpired
	}
	// 如果缓存中的验证码与请求中的验证码不一致，返回验证码错误
	if cacheCode != c {
		return code.ErrVerificationCode
	}
	return nil
}

// CheckMobile 检查手机号格式是否正确
func CheckMobile(mobile string) bool {
	// 手机号正则表达式
	// 正则表达式中：^表示匹配字符串的开头，$表示匹配字符串的结尾
	reg := `^1\d{10}$`

	// 编译正则表达式
	regex := regexp.MustCompile(reg)

	// 使用正则表达式匹配手机号
	return regex.MatchString(mobile)
}
