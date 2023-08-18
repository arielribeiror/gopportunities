package handler

import (
	"fmt"
	"net/http"

	"github.com/arielribeiror/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func GetOpportunityHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opportunity := schemas.Opportunity{}

	// Find Opportunity
	if err := db.First(&opportunity, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opportunity with id: %s not found", id))
		return
	}

	sendSuccess(ctx, "get-opportunity", opportunity)
}
