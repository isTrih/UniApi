package v1

import "github.com/gogf/gf/v2/frame/g"

type UserRegisterReq struct {
	g.Meta   `path:"user/register" method:"get"`
	NickName string `dc:"用户名"`
	Phone    string `dc:"手机号"`
	Password string `dc:"密码"`
	SmsCode  string `dc:"验证码"`

	//注册的同时使用ImHandler注册云信账户
}

type UserRegisterRes struct {
	Reply string `dc:"消息"`
}
