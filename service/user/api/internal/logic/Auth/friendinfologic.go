package Auth

import (
	"context"

	"akita/panda-im/service/user/api/internal/svc"
	"akita/panda-im/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendInfoLogic {
	return &FriendInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendInfoLogic) FriendInfo(req *types.FriendInfoRequest) (resp *types.FriendInfoResponse, err error) {
	//// todo: add your logic here and delete this line
	//friendInfo, err := l.svcCtx.UserRPC.GetFriendsInfoByID(l.ctx, &pb.FriendsInfoRequest{
	//	UserId:   req.UserID,
	//	FriendId: req.FriendID,
	//	Role:     int32(req.Role),
	//})
	//
	//if err != nil {
	//	return nil, err
	//}
	//return &types.FriendInfoResponse{
	//	UserID:   friendInfo.UserId,
	//	Nickname: friendInfo.Nickname,
	//	Abstract: &friendInfo.Abstract,
	//	Avatar:   &friendInfo.Avatar,
	//	Gender:   &friendInfo.Gender,
	//	Addr:     &friendInfo.Addr,
	//	Notice:   friendInfo.Notice,
	//}, nil
	return &types.FriendInfoResponse{Notice: "aa"}, nil
}
