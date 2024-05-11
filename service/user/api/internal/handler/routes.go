// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	Auth "akita/panda-im/service/user/api/internal/handler/Auth"
	"akita/panda-im/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/friendInfo",
					Handler: Auth.FriendInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/userConfig",
					Handler: Auth.UserConfigHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/userConfigUpdate",
					Handler: Auth.UserConfigUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/userInfo",
					Handler: Auth.UserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1/api/user"),
	)
}
