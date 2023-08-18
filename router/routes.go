package router

import (
	"github.com/arielribeiror/gopportunities/handler"
	"github.com/gin-gonic/gin"

	docs "github.com/arielribeiror/gopportunities/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/opportunity", handler.GetOpportunityHandler)
		v1.POST("/opportunity", handler.CreateOpportunityHandler)
		v1.DELETE("/opportunity", handler.DeleteOpportunityHandler)
		v1.PUT("/opportunity", handler.UpdateOpportunityHandler)
		v1.GET("/opportunities", handler.GetAllOpportunityHandler)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
