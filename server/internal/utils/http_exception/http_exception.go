package httpException

var (
	BadRequestHttpException   = New(400, "错误的客户端请求")
	UnauthorizedHttpException = New(401, "无效的认证信息")
	ForbiddenHttpException    = New(403, "无权限操作")
	ServerHttpException       = New(500, "服务器错误")
)

func New(status int, msg string) HttpException {
	return HttpException{
		status: status,
		msg:    msg,
	}
}

type HttpException struct {
	status int
	msg    string
}

func (e HttpException) Error() string {
	return e.msg
}

func (e *HttpException) GetStatus() int {
	return e.status
}

func (e *HttpException) GetMsg() string {
	return e.msg
}

func (e HttpException) SetMsg(msg string) HttpException {
	e.msg = msg
	return e
}
