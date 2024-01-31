package user

import (
	"UniApi/api/user/v1"
	"UniApi/internal/service"
	"context"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) UserRegister(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	r := g.RequestFromCtx(ctx)
	md := g.Model("user")

	//云信业务代码
	accid := "u" + req.Phone
	token := gbase64.EncodeString(req.Phone)
	data := g.Map{
		"name":        req.NickName,
		"phone":       req.Phone,
		"password":    req.Password1,
		"sign":        "Ta太懒了，还没有留下个性签名",
		"accid":       accid,
		"acc_token":   token,
		"school_role": req.Role,
	}

	isuser, err := md.Where("phone", req.Phone).One()
	if err == nil {
		if isuser == nil {
			g.Log().Debugf(ctx, `无用户，插入数据`)
			result, err := md.Insert(data)
			if err == nil {
				r.Response.WriteJson(result)
				g.Log().Debugf(ctx, `成功: %+v`, result)
				service.ImRegister(
					accid,
					token,
					req.NickName,
					"Ta太懒了，还没有留下个性签名",
					req.Phone,
				)
			} else {
				r.Response.WriteJson(result)
				g.Log().Debugf(ctx, `发生错误: %+v`, result)
			}
		} else {
			err = gerror.New("用户已经存在")
		}
	}
	res = &v1.UserRegisterRes{
		Token:     accid,
		TPsession: token,
	}
	return
}
