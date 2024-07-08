package grpc_service

import (
	pb "apigateway/proto/grpc_service"
	"fmt"

	"google.golang.org/grpc"
)

func InitServiceClient() pb.GrpcServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewGrpcServiceClient(cc)
}
