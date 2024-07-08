package grpc

import (
	"context"
	"grpc/business/usecase"
	"grpc/handler/grpc/large"
	"grpc/handler/grpc/user"
	pb "grpc/pb"
)

type Service interface {
	user.Interface
	large.Interface
}

type serviceImpl struct {
	user  user.Interface
	large large.Interface
}

func InitService(u *usecase.Usecase) Service {
	uc := &serviceImpl{
		user:  user.Init(u.User),
		large: large.Init(),
	}

	return uc
}

func (s *serviceImpl) GetList(ctx context.Context, req *pb.GetListRequest) (*pb.GetListResponse, error) {
	return s.user.GetList(ctx, req)
}

func (s *serviceImpl) GetLarge(ctx context.Context, req *pb.GetLargeRequest) (*pb.GetLargeResponse, error) {
	return s.large.GetLarge(ctx, req)
}
