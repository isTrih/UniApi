package v1

import "github.com/gogf/gf/v2/frame/g"

type UserRegisterReq struct {
	g.Meta    `path:"user/register" method:"get"`
	NickName  string `dc:"用户名" v:"required|max-length:16#请输入用户名|用户名最大长度为{max}位"`
	Phone     string `dc:"手机号" v:"required|phone#请输入手机号|请输入正确到手机号"`
	Password1 string `dc:"密码" v:"required|length:6,30#请输入密码|密码长度为{min}到{max}位"`
	Password2 string `dc:"重复密码" v:"required|same:Password1#请确认密码|两次密码不相符"`
	SmsCode   string `dc:"验证码" v:"required|size:6#请输入验证码|验证码长度不相符"`
	Role      string `dc:"用户角色" d:"113050000000" v:"size:12#请输入12位数字码"`
}

type UserRegisterRes struct {
	Token     string `dc:"acctoken与云信账号关联"`
	TPsession string `dc:"鉴权"`
}
