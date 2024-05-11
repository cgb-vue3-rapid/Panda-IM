package logic

import (
	"context"

	"akita/panda-im/service/auth/internal/svc"
	"akita/panda-im/service/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test(req *types.TestRequest) (resp *types.TestResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
