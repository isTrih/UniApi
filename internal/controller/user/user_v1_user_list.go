package user

import (
	"UniApi/api/user/v1"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	r := g.RequestFromCtx(ctx)
	md := g.Model("user")
	user, err := md.One()
	if err == nil {
		g.Log().Debugf(ctx, `receive say: %+v`, user.Json())
	}
	r.Response.WriteJson(user)
	return
}
