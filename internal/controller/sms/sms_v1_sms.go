package sms

import (
	"UniApi/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"

	"UniApi/api/sms/v1"
)

func (c *ControllerV1) Sms(ctx context.Context, req *v1.SmsReq) (res *v1.SmsRes, err error) {
	//构建验证码数据
	smscode := grand.Digits(6)
	smsdata := map[string]string{
		"phone": req.Phone,
		"code":  smscode,
	}
	//存储验证码
	md := g.Model("sms")
	result, err := md.Insert(smsdata)
	if err != nil {
		//错误
	} else {
		fmt.Println(result)
		g.Log("成功")
	}
	//发送短信
	sms := service.SmsHandler{
		Phone:        req.Phone,
		Signature:    req.Signature,
		TemplateId:   req.TemplateId,
		TemplateData: smsdata,
	}
	sres, serr := sms.SendMessage()
	if serr == nil {
		res = &v1.SmsRes{SmsStatus: sres["msg"]}
	}
	return
}
