package controllers

import (
	"crypto/subtle"

	"github.com/gin-gonic/gin"
	"github.com/lmduc/contest-stats/env"
)

var (
	// AuthToken is the authentication token
	AuthToken string = env.GetEnv("TOKEN")
)

type authHeader struct {
	Token string `json:"token"`
}

// InvalidTokenErr means the token is invalid
type InvalidTokenErr string

func (e InvalidTokenErr) Error() string {
	return string(e)
}

func authenticate(ctx *gin.Context) {
	auth := authHeader{}
	if err := ctx.ShouldBindHeader(&auth); err != nil {
		panic(err)
	}

	if subtle.ConstantTimeCompare([]byte(auth.Token), []byte(AuthToken)) == 0 {
		panic(InvalidTokenErr("Invalid token"))
	}
}
