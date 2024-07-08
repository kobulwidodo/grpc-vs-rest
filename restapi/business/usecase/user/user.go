package user

import (
	"context"
	userDom "restapi/business/domain/user"
	"restapi/business/entity"
)

type Interface interface {
	GetList(ctx context.Context) ([]entity.User, error)
}

type user struct {
	user userDom.Interface
}

func Init(ad userDom.Interface) Interface {
	u := &user{
		user: ad,
	}

	return u
}

func (u *user) GetList(ctx context.Context) ([]entity.User, error) {
	users, err := u.user.GetList(ctx)
	if err != nil {
		return users, err
	}

	return users, nil
}
