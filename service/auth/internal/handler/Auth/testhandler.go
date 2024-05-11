package Auth

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"akita/panda-im/service/auth/internal/logic/Auth"
	"akita/panda-im/service/auth/internal/svc"
	"akita/panda-im/service/auth/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TestRequest
		if err := httpx.Parse(r, &req); err != nil {
			logx.Errorf("http parse error: %v", err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := Auth.NewTestLogic(r.Context(), svcCtx)
		resp, err := l.Test(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
