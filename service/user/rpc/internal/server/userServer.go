// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"akita/panda-im/service/user/rpc/internal/logic"
	"akita/panda-im/service/user/rpc/internal/svc"
	"akita/panda-im/service/user/rpc/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	l := logic.NewLogoutLogic(ctx, s.svcCtx)
	return l.Logout(in)
}

func (s *UserServer) GetUserInfoByID(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	l := logic.NewGetUserInfoByIDLogic(ctx, s.svcCtx)
	return l.GetUserInfoByID(in)
}

func (s *UserServer) GetUserConfigByID(ctx context.Context, in *pb.UserConfigRequest) (*pb.UserConfigResponse, error) {
	l := logic.NewGetUserConfigByIDLogic(ctx, s.svcCtx)
	return l.GetUserConfigByID(in)
}

func (s *UserServer) GetFriendsInfoByID(ctx context.Context, in *pb.FriendsInfoRequest) (*pb.FriendsInfoResponse, error) {
	l := logic.NewGetFriendsInfoByIDLogic(ctx, s.svcCtx)
	return l.GetFriendsInfoByID(in)
}
