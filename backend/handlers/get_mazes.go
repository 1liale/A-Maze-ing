package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMazes(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"health": "GOOD",
	})
}