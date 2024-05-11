package Auth

import (
	"akita/panda-im/common/constants"
	"akita/panda-im/service/user/api/internal/svc"
	"akita/panda-im/service/user/api/internal/types"
	"akita/panda-im/service/user/rpc/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserConfigLogic {
	return &UserConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserConfigLogic) UserConfig() (resp *types.UserConfigResponse, err error) {
	//获取用户ID
	userId := l.ctx.Value(constants.UserId).(int64)

	c, err := l.svcCtx.UserRPC.GetUserConfigByID(l.ctx, &user.UserConfigRequest{UserId: userId})
	if err != nil {
		return nil, err
	}

	config := &types.UserConfigResponse{
		UserID:           c.UserId,
		RecallMessage:    c.RecallMessage,
		FriendsOnline:    c.FriendsOnline,
		AllSounds:        c.AllSounds,
		SecureLink:       c.SecureLink,
		SavePwd:          c.SavePwd,
		SearchUser:       int8(c.SearchUser),
		VerificationType: int8(c.VerificationType),
		Message:          "获取用户配置信息成功",
	}

	if c.VerificationQuestion != nil {
		config.VerificationQuestion = &types.VerificationQuestion{
			Question1: c.VerificationQuestion.Question1,
			Question2: c.VerificationQuestion.Question2,
			Question3: c.VerificationQuestion.Question3,
			Answer1:   c.VerificationQuestion.Answer1,
			Answer2:   c.VerificationQuestion.Answer2,
			Answer3:   c.VerificationQuestion.Answer3,
		}
	}

	return config, nil
}
