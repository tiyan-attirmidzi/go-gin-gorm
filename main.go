package main

import (
	"github.com/tiyan-attirmidzi/go-rest-api/configs"
	"github.com/tiyan-attirmidzi/go-rest-api/controllers"
	"github.com/tiyan-attirmidzi/go-rest-api/repositories"
	"github.com/tiyan-attirmidzi/go-rest-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                    = configs.DatabaseConnection()
	userRepository repositories.UserRepository = repositories.NewUserRespository(db)
	jwtService     services.JWTService         = services.NewJWTService()
	authService    services.AuthService        = services.NewAuthService(userRepository)
	authController controllers.AuthController  = controllers.NewAuthController(authService, jwtService)
)

func main() {
	defer configs.DatabaseDisconnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/signin", authController.SignIn)
		authRoutes.POST("/signup", authController.SignUp)
	}

	r.Run()
}
