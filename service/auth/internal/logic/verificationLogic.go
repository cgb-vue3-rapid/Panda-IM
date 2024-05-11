package logic

import (
	"akita/panda-im/common/constants"
	"akita/panda-im/common/util/random_munber"
	"akita/panda-im/common/util/rds_cache"
	"akita/panda-im/service/auth/code"
	"akita/panda-im/service/auth/internal/svc"
	"akita/panda-im/service/auth/internal/types"
	"context"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
// prefixVerificationCount = "biz#verification#count#%s" // 验证码计数的 Redis 键前缀
// verificationLimitPerDay = 100                         // 每日发送验证码的上限
// expireActivation        = 60 * 100000                 // 验证码缓存失效时间，单位为秒
)

type VerificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerificationLogic {
	return &VerificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerificationLogic) Verification(req *types.VerificationRequest) (resp *types.VerificationResponse, err error) {
	req.Mobile = strings.TrimSpace(req.Mobile)

	// 检查手机号是否为空
	if len(req.Mobile) == 0 {
		return nil, code.ErrMobileEmpty
	}

	// 检查手机号格式
	if !CheckMobile(req.Mobile) {
		return nil, code.ErrMobileFormatError
	}

	// 获取今日发送验证码次数
	count, err := rds_cache.GetVerificationCountByMobile(req.Mobile, constants.PrefixVerificationCount, l.svcCtx.BizRedis)
	if err != nil {
		logx.Error(l.ctx, "GetVerificationCountByMobile mobile: %s error: %v", req.Mobile, err)
	}
	if count > constants.VerificationLimitPerDay {
		// 发送验证码次数超过上限
		return nil, code.ErrVerificationLimitExceeded
	}

	// 生成或获取已存在的验证码
	verifyCode, err := rds_cache.GetActivationCacheByMobile(req.Mobile, constants.PrefixActivation, l.svcCtx.BizRedis)
	if err != nil {
		logx.Errorf("etActivationCacheByMobile mobile: %s error: %v", req.Mobile, err)
	}
	if len(verifyCode) == 0 {
		verifyCode = random_munber.RandomNumeric(6)
	}

	//// 发送验证码
	//_, err = l.svcCtx.UserRPC.SendSms(l.ctx, &user.SendSmsRequest{
	//	Mobile: req.Mobile,
	//})

	if err != nil {
		logx.Error(l.ctx, "发送验证码失败: %s error: %v", req.Mobile, err)
		//  发送验证码失败，返回错误
		return nil, code.ErrSendSmsFailed
	}

	// 缓存验证码
	err = rds_cache.SaveActivationCacheByMobile(req.Mobile, verifyCode, constants.PrefixActivation, l.svcCtx.BizRedis, constants.VerificationExpire)
	if err != nil {
		logx.Error(l.ctx, "缓存验证码失败: %s error: %v", req.Mobile, err)
		return nil, code.ErrSendSmsFailed
	}

	// 增加验证码发送次数计数
	err = rds_cache.IncrVerificationCountByMobile(req.Mobile, constants.PrefixVerificationCount, l.svcCtx.BizRedis)
	if err != nil {
		logx.Error(l.ctx, "发送次数计数失败: %s error: %v", req.Mobile, err)
	}

	return &types.VerificationResponse{
		Message: "发送成功",
	}, nil

}
