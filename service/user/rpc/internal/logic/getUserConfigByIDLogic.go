package logic

import (
	"akita/panda-im/service/user/rpc/internal/svc"
	"akita/panda-im/service/user/rpc/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserConfigByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserConfigByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserConfigByIDLogic {
	return &GetUserConfigByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserConfigByIDLogic) GetUserConfigByID(in *pb.UserConfigRequest) (*pb.UserConfigResponse, error) {
	// todo: add your logic here and delete this line
	c, err := l.svcCtx.UserDao.GetConfigByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	config := &pb.UserConfigResponse{
		UserId:           c.UserId,
		RecallMessage:    c.RecallMessage,
		FriendsOnline:    c.FriendsOnline,
		AllSounds:        c.AllSounds,
		SecureLink:       c.SecureLink,
		SavePwd:          c.SavePwd,
		SearchUser:       c.SearchUser,
		VerificationType: c.VerificationType,
	}

	if c.VerificationType == 3 || c.VerificationType == 4 {
		if c.VerificationQuestion != nil {
			config.VerificationQuestion = &pb.VerificationQuestion{
				Question1: *c.VerificationQuestion.Question1,
				Question2: *c.VerificationQuestion.Question2,
				Question3: *c.VerificationQuestion.Question3,
				Answer1:   *c.VerificationQuestion.Answer1,
				Answer2:   *c.VerificationQuestion.Answer2,
				Answer3:   *c.VerificationQuestion.Answer3,
			}
		}
	}

	return config, nil
}
