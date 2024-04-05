package handler

import (
	"akita/panda-im/common/manage"
	"akita/panda-im/service/auth/internal/logic"
	"akita/panda-im/service/auth/internal/svc"
	"net/http"
)

func authenticateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAuthenticateLogic(r.Context(), svcCtx)
		resp, err := l.Authenticate()
		manage.ResponseHandler(w, resp, err)
	}
}
