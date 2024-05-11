package middleware

import (
	"akita/panda-im/common/constants"
	"akita/panda-im/common/util/token_manager"
	"akita/panda-im/common/xcode"
	"akita/panda-im/service/user/api/internal/config"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
)

type JwtAuthMiddleware struct {
	Config config.Config
}

func NewJwtAuthMiddleware(c config.Config) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		Config: c,
	}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("token")
		if token == "" {
			httpx.ErrorCtx(r.Context(), w, xcode.TokenIsEmpty)
			return
		}

		parts := strings.Split(token, " ")

		if len(parts) != 2 {
			httpx.ErrorCtx(r.Context(), w, xcode.TokenFormatErr)
			return
		}

		parseToken, isExpire, err := token_manager.ParseToken(parts[0], parts[1], m.Config.Token.AccessSecret, m.Config.Token.RefreshSecret)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, xcode.TokenParseErr)
			return
		}

		if isExpire {
			parts[0], parts[1], err = token_manager.GenToken(parseToken.UserID, parseToken.Nickname, m.Config.Token.AccessSecret, m.Config.Token.RefreshSecret, parseToken.Role)
			if err != nil {
				logx.Errorf("token generate err: %v", err)
				httpx.ErrorCtx(r.Context(), w, xcode.TokenGenerateErr)
				return
			}
		}
		r = r.WithContext(context.WithValue(r.Context(), constants.UserId, parseToken.UserID))
		r = r.WithContext(context.WithValue(r.Context(), constants.AccessToken, parts[0]))
		r = r.WithContext(context.WithValue(r.Context(), constants.RefreshToken, parts[1]))

		next(w, r)
		//logx.Errorf("重新登陆: %v", err)
		//httpx.ErrorCtx(r.Context(), w, xcode.TokenInvalid)
	}
}
