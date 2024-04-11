package Auth

import (
	"net/http"

	"akita/panda-im/service/auth/internal/logic/Auth"
	"akita/panda-im/service/auth/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuthenticateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := Auth.NewAuthenticateLogic(r.Context(), svcCtx)
		resp, err := l.Authenticate()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
