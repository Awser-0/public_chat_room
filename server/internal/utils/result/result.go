package result

type Result struct {
	code int
	msg  string
	data any
}

func (r Result) SetData(data any) Result {
	r.data = data
	return r
}

func (r Result) SetMsg(msg string) Result {
	r.msg = msg
	return r
}

func (r Result) ToMap() map[string]any {
	return map[string]any{
		"code": r.code,
		"msg":  r.msg,
		"data": r.data,
	}
}

func New(code int, msg string, data any) Result {
	return Result{
		code: code,
		msg:  msg,
		data: data,
	}
}
