// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"UniApi/api/user/v1"
)

type IUserV1 interface {
	UserAccount(ctx context.Context, req *v1.UserAccountReq) (res *v1.UserAccountRes, err error)
	UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error)
	UserRegister(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error)
}
