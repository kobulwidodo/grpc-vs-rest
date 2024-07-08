package main

import (
	grpc_client "apigateway/client/grpc"
	rest_client "apigateway/client/rest"
	rest "apigateway/handler"
)

func main() {
	grpcClient := grpc_client.Init()

	restClient := rest_client.Init()

	r := rest.Init(grpcClient, restClient)

	r.Run()
}
