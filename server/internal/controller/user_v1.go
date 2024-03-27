package controller

import "github.com/gogf/gf/v2/net/ghttp"

type IUserV1 interface {
	Login(r *ghttp.Request)
}
