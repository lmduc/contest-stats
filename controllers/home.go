package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home provides functions for handling the requests of controllers
type Home struct {
	cfg *Config
}

// NewHome returns the user holding the configuration
func NewHome(cfg *Config) *Home {
	return &Home{
		cfg: cfg,
	}
}

// HandlePing handles ping
func (h *Home) HandlePing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"pong": true,
	})
}
