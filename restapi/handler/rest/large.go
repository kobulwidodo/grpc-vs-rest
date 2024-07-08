package rest

import (
	"net/http"
	"restapi/business/entity"

	"github.com/gin-gonic/gin"
)

func (r *rest) GetLargeData(ctx *gin.Context) {
	res := []entity.S{}
	i := 1
	for i <= 10000 {
		res = append(res, entity.S{A: i})
		i++
	}

	ctx.JSON(http.StatusOK, gin.H{"data": res})
}
