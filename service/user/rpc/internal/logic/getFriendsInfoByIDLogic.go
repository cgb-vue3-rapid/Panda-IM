package logic

import (
	"context"

	"akita/panda-im/service/user/rpc/internal/svc"
	"akita/panda-im/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendsInfoByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendsInfoByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendsInfoByIDLogic {
	return &GetFriendsInfoByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendsInfoByIDLogic) GetFriendsInfoByID(in *pb.FriendsInfoRequest) (*pb.FriendsInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.FriendsInfoResponse{}, nil
}
