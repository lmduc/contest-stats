package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func errorHandler(ctx *gin.Context) {
	if r := recover(); r != nil {
		switch v := r.(type) {
		case InvalidTokenErr:
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": v.Error(),
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
		}
	}
}
