package middleware

import (
	"chat/internal/consts"
	"chat/internal/utils"
	httpException "chat/internal/utils/http_exception"

	"github.com/gogf/gf/v2/net/ghttp"
)

func MiddlewareAuth(r *ghttp.Request) {
	token := r.Header.Get("token")
	customClaims, err := utils.ParseTokenHs256(token)
	if err != nil || customClaims == nil {
		r.SetError(httpException.UnauthorizedHttpException)
	} else {
		r.SetCtxVar(consts.CTX_VAR_KEY_USERNAME, customClaims.User)
		r.Middleware.Next()
	}
}
