package Auth

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"akita/panda-im/service/user/api/internal/logic/Auth"
	"akita/panda-im/service/user/api/internal/svc"
	"akita/panda-im/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FriendInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			logx.Errorf("url:%s,err:%v", r.URL.Path, err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := Auth.NewFriendInfoLogic(r.Context(), svcCtx)
		resp, err := l.FriendInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
