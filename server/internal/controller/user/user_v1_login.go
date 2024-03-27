package user

import (
	"chat/internal/utils"
	httpException "chat/internal/utils/http_exception"
	"chat/internal/utils/result"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type loginDTO struct {
	Username string `json:"username" v:"required|length:2,10#请输入username|username长度为2-10"`
	Password string `json:"password" v:"required#请输入password"`
}

func (c *ControllerV1) Login(r *ghttp.Request) {
	var form *loginDTO
	if err := r.Parse(&form); err != nil {
		r.SetError(httpException.BadRequestHttpException.SetMsg(err.Error()))
		return
	}
	token, err := utils.GenerateTokenUsingHs256(form.Username)
	if err != nil {
		panic(gerror.Wrap(err, "token 生成失败"))
	}
	if form.Password == "123456" {
		r.Response.WriteJson(result.OK.SetMsg("登录成功").SetData(g.Map{
			"username": form.Username,
			"token":    token,
		}).ToMap())
	} else {
		r.Response.WriteJson(result.UserFailPass.ToMap())
	}
}
