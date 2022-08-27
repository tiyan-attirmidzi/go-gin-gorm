package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type authentication struct {
}

type Authentication interface {
	SignIn(ctx *gin.Context)
	SignUp(ctx *gin.Context)
}

func AuthenticationController() Authentication {
	return &authentication{}
}

func (c *authentication) SignIn(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello SignIn",
	})
}

func (c *authentication) SignUp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello SignUp",
	})
}
