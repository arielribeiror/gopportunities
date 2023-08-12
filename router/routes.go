package router

import (
	"github.com/arielribeiror/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opportunity", handler.GetOpportunityHandler)
		v1.POST("/opportunity", handler.CreateOpportunityHandler)
		v1.DELETE("/opportunity", handler.DeleteOpportunityHandler)
		v1.PUT("/opportunity", handler.UpdateOpportunityHandler)
		v1.GET("/opportunities", handler.GetAllOpportunityHandler)
	}
}
