package service

import (
	unisms "github.com/apistd/uni-go-sdk/sms"
)

type SmsHandler struct {
	Phone        string
	Signature    string
	TemplateId   string
	TemplateData map[string]string
}

func (s *SmsHandler) SendMessage() (nres map[string]string, err error) {

	//初始化
	client := unisms.NewClient("khT5bVeXRYnScmtum4Q26hpyeEdRKwDsoJHQbrxANUD4fSqP3")

	//构建信息
	message := unisms.BuildMessage()
	message.SetTo(s.Phone)
	message.SetSignature(s.Signature)
	message.SetTemplateId(s.TemplateId)
	message.SetTemplateData(s.TemplateData)

	//发送短信
	res, err := client.Send(message)
	if err != nil {
		return
	}
	nres = map[string]string{
		"msg": res.Message,
	}
	return
}
