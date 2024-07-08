package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *rest) GetListUsers(ctx *gin.Context) {
	user, err := r.uc.User.GetList(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
