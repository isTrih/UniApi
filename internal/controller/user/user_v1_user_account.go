package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"UniApi/api/user/v1"
)

func (c *ControllerV1) UserAccount(ctx context.Context, req *v1.UserAccountReq) (res *v1.UserAccountRes, err error) {
	r := g.RequestFromCtx(ctx)
	uid := r.GetRouter("uid")
	user, err := g.Model("user").Where("uid", uid).One()
	if err == nil {
		if user == nil {
			//r.Response.Writeln("查询不到该用户")
			err = gerror.New("查询不到该用户")
		} else {
			//r.Response.WriteJson(user)

			res = &v1.UserAccountRes{UserInfo: user.Map()}
		}
		g.Log().Debugf(ctx, `receive say: %+v`, user)
	}
	return
}
