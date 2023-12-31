package handler

import (
	"net/http"

	"github.com/arielribeiror/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create opportunity
// @Description Create a new job opportunity
// @Tags Opportunities
// @Accept json
// @Produce json
// @Param request body CreateOpportunityRequest true "Request Body"
// @Success 200 {object} CreateOpportunityResponse
// @Failures 400 {object} ErrorResponse
// @Failures 500 {object} ErrorResponse
// @Router /opportunity [post]
func CreateOpportunityHandler(ctx *gin.Context) {
	request := CreateOpportunityRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opportunity := schemas.Opportunity{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
		Range:    request.Range,
	}

	if err := db.Create(&opportunity).Error; err != nil {
		logger.Errorf("error creating opportunity: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opportunity on database")
		return
	}

	sendSuccess(ctx, "create-opportunity", opportunity)
}
