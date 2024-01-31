package v1

import "github.com/gogf/gf/v2/frame/g"

type UserAccountReq struct {
	g.Meta `path:"user/u{uid}" method:"get"`
}
type UserAccountRes struct {
	UserInfo map[string]any
}
