package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SystemCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"health": "GOOD",
	})
}
