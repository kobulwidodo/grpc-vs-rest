package user

import (
	"context"
	"restapi/business/entity"

	"gorm.io/gorm"
)

type Interface interface {
	GetList(ctx context.Context) ([]entity.User, error)
}

type user struct {
	db *gorm.DB
}

func Init(db *gorm.DB) Interface {
	u := &user{
		db: db,
	}

	return u
}

func (u *user) GetList(ctx context.Context) ([]entity.User, error) {
	res := []entity.User{}

	if err := u.db.Unscoped().Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}
