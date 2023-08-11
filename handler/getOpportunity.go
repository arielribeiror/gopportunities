package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpportunityHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET opportunity",
	})
}
