package handler

import (
	"akita/panda-im/common/manage"
	"akita/panda-im/service/auth/internal/logic"
	"akita/panda-im/service/auth/internal/svc"
	"akita/panda-im/service/auth/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			manage.ResponseHandler(w, nil, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		manage.ResponseHandler(w, resp, err)
	}
}
