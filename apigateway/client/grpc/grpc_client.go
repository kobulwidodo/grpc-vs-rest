package grpc_client

import (
	"apigateway/client/grpc/grpc_service"
	pb "apigateway/proto/grpc_service"
)

type GrpcClient struct {
	GrpcService pb.GrpcServiceClient
}

func Init() *GrpcClient {
	gc := &GrpcClient{
		GrpcService: grpc_service.InitServiceClient(),
	}
	return gc
}
