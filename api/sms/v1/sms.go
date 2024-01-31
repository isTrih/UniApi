package v1

import "github.com/gogf/gf/v2/frame/g"

type SmsReq struct {
	g.Meta     `path:"sms" type:"get"`
	Phone      string `dc:"接收手机号" v:"required|phone#请输入手机号|请输入正确的手机号"`
	Signature  string `dc:"发送短信的签名" d:"三氢不是氕氘氚"`
	TemplateId string `dc:"发送短信的模版" d:"pub_verif_basic2"`
}

type SmsRes struct {
	SmsStatus string `dc:"返回，短信发送情况"`
}
