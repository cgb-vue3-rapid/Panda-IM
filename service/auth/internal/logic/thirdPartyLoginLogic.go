package logic

import (
	"context"

	"akita/panda-im/service/auth/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ThirdPartyLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThirdPartyLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThirdPartyLoginLogic {
	return &ThirdPartyLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThirdPartyLoginLogic) ThirdPartyLogin() (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
