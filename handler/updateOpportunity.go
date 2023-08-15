package handler

import (
	"fmt"
	"net/http"

	"github.com/arielribeiror/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpportunityHandler(ctx *gin.Context) {
	request := UpdateOpportunityRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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

	// Update Opportunity
	if request.Role != "" {
		opportunity.Role = request.Role
	}
	if request.Company != "" {
		opportunity.Company = request.Company
	}
	if request.Location != "" {
		opportunity.Location = request.Location
	}
	if request.Remote != nil {
		opportunity.Remote = *request.Remote
	}
	if request.Link != "" {
		opportunity.Link = request.Link
	}
	if request.Salary != 0 {
		opportunity.Salary = request.Salary
	}
	if request.Range != 0 {
		opportunity.Range = request.Range
	}

	// Save Opportunity
	if err := db.Save(&opportunity).Error; err != nil {
		logger.Errorf("error updating opportunity: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opportunity on database")
		return
	}

	sendSuccess(ctx, "update-opportunity", opportunity)
}
