// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	Auth "akita/panda-im/service/auth/internal/handler/Auth"
	"akita/panda-im/service/auth/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/verification",
				Handler: VerificationHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/api/auth"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/authenticate",
					Handler: Auth.AuthenticateHandlerGetHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authenticate",
					Handler: Auth.AuthenticateHandlerPostHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/logout",
					Handler: Auth.LogoutHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/test",
					Handler: Auth.TestHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1/api/auth"),
	)
}
