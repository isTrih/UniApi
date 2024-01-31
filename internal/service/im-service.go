package service

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gsha1"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"os"
)

type ImService struct {
	AppKey    string
	AppSecret string
}

type ImAccount struct {
	AccountId       string           `json:"account_id"`
	Token           string           `json:"token"`
	UserInformation *UserInformation `json:"user_information"`
}

type UserInformation struct {
	Name string `json:"name"`
	//Avatar string `json:"avatar"`
	Sign   string `json:"sign"`
	Mobile string `json:"mobile"`
}

func (h *ImService) httpDo(i ImAccount) {
	AppKey := h.AppKey

	//UTC时间戳获取
	CurTime := gconv.String(gtime.Timestamp())

	//随机数获取[0-100000000)
	Nonce := grand.Digits(8)

	//CheckSum计算方式AppSecret + Nonce + CurTime后去sh1哈希计算
	CheckSum := gsha1.Encrypt(h.AppSecret + Nonce + CurTime)

	println("AppKey =", AppKey)
	println("Nonce =", Nonce)
	println("CurTime =", CurTime)
	println("CheckSum =", CheckSum)

	//g.client实现
	c := g.Client()

	//header设置
	header := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"AppKey":       AppKey,
		"Nonce":        Nonce,
		"CurTime":      CurTime,
		"CheckSum":     CheckSum,
	}

	j, err := json.Marshal(i)
	if err != nil {
		fmt.Println("转换为json错误")
		os.Exit(-1)
	}

	c.SetHeaderMap(header)
	//if r, e := c.Post(gctx.New(), "https://open.yunxinapi.com/im/v2/accounts", j); e != nil {
	//	panic(e)
	//} else {
	//	defer r.Close()
	//	fmt.Println(r.ReadAllString())
	//}

	// 返回content为string类型
	content := c.PostContent(
		gctx.New(),
		"https://open.yunxinapi.com/im/v2/accounts",
		j,
	)
	fmt.Println(content)
}

func ImRegister(accid, token, name, sign, mobile string) {

	//pram string:用来参数 &连接 accid=123456&name=zhangsan
	//accid	String	32	必选	"123456"	云信 IM 账号，必须保证唯一性。若涉及字母，传参时请一律小写处理。只允许字母、数字、半角下划线_、@、半角点以及半角-。请注意以此接口返回结果中的accid为准。
	//token	String	128	选填	"abcdef"	用户账号对应的登录密钥token。如果未指定，云信会自动生成token，并在创建成功后返回。
	//name	String	64	选填	"zhangsan"	用户昵称
	//icon	String	1024	选填	"https://netease/xxx.png"	用户头像 URL
	//sign	String	256	选填	"Hello World"	用户签名
	//email	String	64	选填	"xxx@163.com"	用户邮箱地址
	//birth	String	16	选填	"xxxx-xx-xx"	用户生日
	//mobile	String	32	选填	"+852-xxxxxxxx"	用户手机号码，非中国大陆手机号码需要填写国家代码(如美国：+1-xxxxxxxxxx)或地区代码(如香港：+852-xxxxxxxx)
	//gender	int	/	选填	2	用户性别，0-未知，1-男，2-女。其它会报参数错误。
	//ex	String	1024	选填	{"level":1}	用户资料扩展字段，建议封装成JSON。

	iaccount := ImAccount{
		AccountId: accid,
		Token:     token,
	}
	tempUserInfo := new(UserInformation)
	tempUserInfo.Name = name
	if sign == "" {
		tempUserInfo.Sign = "因为ta太懒了，还没有设置个性签名哦！"
	}
	tempUserInfo.Sign = sign
	tempUserInfo.Mobile = mobile

	iaccount.UserInformation = tempUserInfo
	im := ImService{AppKey: "05b4c8c3883a6426f3590e5f8fba3762", AppSecret: "e3961aba4f93"}
	im.httpDo(iaccount)
}
