package converter

import (
	"grpc/business/entity"
	pb "grpc/pb"
)

func UsersToGetListResponse(users []entity.User) []*pb.GetListData {
	res := []*pb.GetListData{}
	for _, u := range users {
		res = append(res, u.ToGetListData())
	}

	return res
}
