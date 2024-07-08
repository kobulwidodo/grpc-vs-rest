package rest

import (
	pb "apigateway/proto/grpc_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *rest) GetLargeGrpc(ctx *gin.Context) {
	res, err := r.grpcClient.GrpcService.GetLarge(ctx.Request.Context(), &pb.GetLargeRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	ctx.JSON(http.StatusOK, res)
}

func (r *rest) GetLargeRest(ctx *gin.Context) {
	res, err := r.restClient.User.GetLarge()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	ctx.JSON(http.StatusOK, res)
}
