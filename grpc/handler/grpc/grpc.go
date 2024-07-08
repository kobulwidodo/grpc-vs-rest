package grpc

import (
	"grpc/business/usecase"
	pb "grpc/pb"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

var once = &sync.Once{}

type GRPC interface {
	Run()
}

type grpcHandler struct {
	grpcServer  *grpc.Server
	uc          *usecase.Usecase
	grpcService *Service
}

func Init(uc *usecase.Usecase) GRPC {
	g := &grpcHandler{}
	once.Do(func() {

		grpcServer := grpc.NewServer()
		grpcService := InitService(uc)

		g = &grpcHandler{
			uc:          uc,
			grpcServer:  grpcServer,
			grpcService: &grpcService,
		}

		g.Register()
	})

	return g
}

func (g *grpcHandler) Run() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalln("Failed to listen tcp:", err)
	}
	if err := g.grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

func (g *grpcHandler) Register() {
	pb.RegisterGrpcServiceServer(g.grpcServer, *g.grpcService)
}
