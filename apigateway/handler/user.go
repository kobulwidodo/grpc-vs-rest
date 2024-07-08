package rest

import (
	pb "apigateway/proto/grpc_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *rest) GetListUsersGrpc(ctx *gin.Context) {
	user, err := r.grpcClient.GrpcService.GetList(ctx.Request.Context(), &pb.GetListRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	ctx.JSON(http.StatusOK, user)
}

func (r *rest) GetListUsersRest(ctx *gin.Context) {
	user, err := r.restClient.User.GetList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	ctx.JSON(http.StatusOK, user)
}
