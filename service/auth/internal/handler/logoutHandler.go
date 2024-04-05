package handler

import (
	"akita/panda-im/common/manage"
	"akita/panda-im/service/auth/internal/logic"
	"akita/panda-im/service/auth/internal/svc"
	"net/http"
)

func logoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout()
		manage.ResponseHandler(w, resp, err)
	}
}
