package user

import (
	"context"
	userUc "grpc/business/usecase/user"
	pb "grpc/pb"
	"grpc/util/converter"
)

type Interface interface {
	GetList(ctx context.Context, req *pb.GetListRequest) (*pb.GetListResponse, error)
}

type user struct {
	user userUc.Interface
}

func Init(uu userUc.Interface) Interface {
	u := &user{
		user: uu,
	}

	return u
}

func (u *user) GetList(ctx context.Context, req *pb.GetListRequest) (*pb.GetListResponse, error) {
	users, err := u.user.GetList(ctx)
	if err != nil {
		return &pb.GetListResponse{
			Error: err.Error(),
			Data:  []*pb.GetListData{},
		}, err
	}

	return &pb.GetListResponse{
		Data: converter.UsersToGetListResponse(users),
	}, nil
}
