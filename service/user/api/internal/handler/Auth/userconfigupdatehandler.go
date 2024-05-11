package Auth

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"akita/panda-im/service/user/api/internal/logic/Auth"
	"akita/panda-im/service/user/api/internal/svc"
	"akita/panda-im/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserConfigUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserConfigUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			logx.Errorf("[user] UserConfigUpdateHandler httpx.Parse error: %v", err)
			return
		}

		l := Auth.NewUserConfigUpdateLogic(r.Context(), svcCtx)
		resp, err := l.UserConfigUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
