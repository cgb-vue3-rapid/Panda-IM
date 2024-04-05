package logic

import (
	"context"

	"akita/panda-im/service/auth/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticateLogic {
	return &AuthenticateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticateLogic) Authenticate() (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
