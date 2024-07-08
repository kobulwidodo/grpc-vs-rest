package large

import (
	"context"
	pb "grpc/pb"
)

type Interface interface {
	GetLarge(ctx context.Context, req *pb.GetLargeRequest) (*pb.GetLargeResponse, error)
}

type large struct {
}

func Init() Interface {
	l := &large{}

	return l
}

func (l *large) GetLarge(ctx context.Context, req *pb.GetLargeRequest) (*pb.GetLargeResponse, error) {
	res := &pb.GetLargeResponse{}
	var i int64 = 1
	for i <= 10000 {
		res.Data = append(res.Data, &pb.GetLargeData{
			A: i,
		})
		i++
	}

	return res, nil
}
