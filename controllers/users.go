package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lmduc/contest-stats/repositories"
)

// CreateUserCommand contains the structure to create users
type CreateUserCommand struct {
	LeetcodeHandle string `json:"leetcode_handle"`
}

// User provides functions for handling the requests of controllers
type User struct {
	cfg *Config
}

// NewUser returns the user holding the configuration
func NewUser(cfg *Config) *User {
	return &User{
		cfg: cfg,
	}
}

// CreateUser creates a user
func (u *User) CreateUser(ctx *gin.Context) {
	defer errorHandler(ctx)

	authenticate(ctx)

	command := CreateUserCommand{}
	ctx.BindJSON(&command)

	user := repositories.NewUser()
	user.LeetcodeHandle = &command.LeetcodeHandle
	u.cfg.DB.Connection().Create(user)

	ctx.JSON(http.StatusCreated, gin.H{
		"leetcode_handle": command.LeetcodeHandle,
	})
}
