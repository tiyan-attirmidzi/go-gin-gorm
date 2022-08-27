package main

import (
	"go-rest-api/configs"
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB                   = configs.DatabaseConnection()
	auth controllers.Authentication = controllers.AuthenticationController()
)

func main() {
	defer configs.DatabaseRejection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/signin", auth.SignIn)
		authRoutes.POST("/signup", auth.SignUp)
	}

	r.Run()
}
