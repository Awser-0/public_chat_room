package middleware

import (
	httpException "chat/internal/utils/http_exception"
	"chat/internal/utils/result"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func MiddlewareException(r *ghttp.Request) {
	r.Middleware.Next()
	var err = r.GetError()
	if err != nil {
		r.Response.ClearBuffer()
		switch e := err.(type) {
		case httpException.HttpException:
			{
				r.Response.Status = e.GetStatus()
				r.Response.WriteJson(result.New(e.GetStatus(), e.GetMsg(), nil).ToMap())
				break
			}
		default:
			{
				glog.Error(r.Context(), err)
				r.Response.Status = httpException.ServerHttpException.GetStatus()
				httpE := httpException.ServerHttpException
				r.Response.WriteJson(result.New(httpE.GetStatus(), e.Error(), nil).ToMap())
			}
		}
	}
}
