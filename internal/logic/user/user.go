package user

import (
	"UniApi/internal/dao"
	"UniApi/internal/model/do"
	"UniApi/internal/model/entity"
	"UniApi/internal/service"
	"context"
)

func init() {
	service.RegisterUser(&sUser{})
}

type sUser struct{}

func (s sUser) GetList(ctx context.Context) (users []entity.NmcUser, err error) {
	err = dao.NmcUser.Ctx(ctx).Scan(&users)
	return
}

func (s sUser) Add(ctx context.Context, user do.NmcUser) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s sUser) Edit(ctx context.Context, user do.NmcUser) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s sUser) Del(ctx context.Context) (err error) {
	//TODO implement me
	panic("implement me")
}
