// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"UniApi/internal/model/do"
	"UniApi/internal/model/entity"
	"context"
)

type (
	IUser interface {
		GetList(ctx context.Context) (users []entity.NmcUser, err error)
		Add(ctx context.Context, user do.NmcUser) (err error)
		Edit(ctx context.Context, user do.NmcUser) (err error)
		Del(ctx context.Context) (err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
