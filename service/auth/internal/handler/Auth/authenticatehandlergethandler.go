package Auth

import (
	"net/http"

	"akita/panda-im/service/auth/internal/logic/Auth"
	"akita/panda-im/service/auth/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuthenticateHandlerGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := Auth.NewAuthenticateHandlerGetLogic(r.Context(), svcCtx)
		resp, err := l.AuthenticateHandlerGet()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
