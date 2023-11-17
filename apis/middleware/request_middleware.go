package middleware

import (
	"net/http"

	"github.com/aniket-skroman/skroman_sales_service.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ValidateRequest(ctx *gin.Context) {
	id := ctx.Param("lead_id")
	if id == "" {
		ctx.Next()
		return
	}
	lead_id, err := uuid.Parse(id)
	if err != nil {
		response := utils.BuildFailedResponse("new error from middleware")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	ctx.Set("lead_id", lead_id)
	ctx.Next()
}
