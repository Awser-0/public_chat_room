package result

func newByCM(code int, msg string) Result {
	return New(code, msg, nil)
}

var (
	OK   = newByCM(10200, "请求成功")
	Fail = newByCM(10400, "请求失败")

	UserNotFound = newByCM(11001, "用户不存在")
	UserExists   = newByCM(11002, "用户已存在")
	UserFailPass = newByCM(11005, "账号或密码错误")
)
