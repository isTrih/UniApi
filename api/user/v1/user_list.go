package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserListReq struct {
	g.Meta `path:"user/list" method:"get"`
}

type UserListRes struct {
}
