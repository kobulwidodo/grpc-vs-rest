package usecase

import (
	"grpc/business/domain"
	"grpc/business/usecase/user"
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
