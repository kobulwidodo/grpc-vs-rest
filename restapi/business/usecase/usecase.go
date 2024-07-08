package usecase

import (
	"restapi/business/domain"
	"restapi/business/usecase/user"
)

type Usecase struct {
	User user.Interface
}

func Init(d *domain.Domains) *Usecase {
	uc := &Usecase{
		User: user.Init(d.User),
	}

	return uc
}
