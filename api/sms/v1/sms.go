package v1

import (
	"fmt"
	unisms "github.com/apistd/uni-go-sdk/sms"
)

func main() {
	// 初始化
	client := unisms.NewClient("your access key id", "your access key secret") // 若使用简易验签模式仅传入第一个参数即可

	// 构建信息
	message := unisms.BuildMessage()
	message.SetTo("your phone number")
	message.SetSignature("UniSMS")
	message.SetTemplateId("login_tmpl")
	message.SetTemplateData(map[string]string{"code": "6666"}) // 设置自定义参数 (变量短信)

	// 发送短信
	res, err := client.Send(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
