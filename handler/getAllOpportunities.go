package handler

import (
	"net/http"

	"github.com/arielribeiror/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List all opportunities
// @Description List all job opportunities
// @Tags Opportunities
// @Accept json
// @Produce json
// @Success 200 {object} GetAllOpportunities
// @Failures 500 {object} ErrorResponse
// @Router /opportunities [get]
func GetAllOpportunityHandler(ctx *gin.Context) {
	opportunities := []schemas.Opportunity{}

	if err := db.Find(&opportunities).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing opportunities")
		return
	}

	sendSuccess(ctx, "list-opportunities", opportunities)
}
