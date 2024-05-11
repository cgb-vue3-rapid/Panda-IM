package Auth

import (
	"net/http"

	"akita/panda-im/service/user/api/internal/logic/Auth"
	"akita/panda-im/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := Auth.NewUserConfigLogic(r.Context(), svcCtx)
		resp, err := l.UserConfig()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
