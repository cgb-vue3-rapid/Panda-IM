package logic

import (
	"context"

	"akita/panda-im/service/auth/internal/svc"
	"akita/panda-im/service/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThirdPartyLoignInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThirdPartyLoignInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThirdPartyLoignInfoLogic {
	return &ThirdPartyLoignInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThirdPartyLoignInfoLogic) ThirdPartyLoignInfo() (resp []types.ThirdPartyLoginInfoByQQReponse, err error) {
	// todo: add your logic here and delete this line

	return
}
