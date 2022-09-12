package main

import (
	"github.com/tiyan-attirmidzi/go-rest-api/configs"
	"github.com/tiyan-attirmidzi/go-rest-api/controllers"
	"github.com/tiyan-attirmidzi/go-rest-api/middlewares"
	"github.com/tiyan-attirmidzi/go-rest-api/repositories"
	"github.com/tiyan-attirmidzi/go-rest-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                    = configs.DatabaseConnection()
	userRepository repositories.UserRepository = repositories.NewUserRepository(db)
	bookRepository repositories.BookRepository = repositories.NewBookRepository(db)
	jwtService     services.JWTService         = services.NewJWTService()
	authService    services.AuthService        = services.NewAuthService(userRepository)
	userService    services.UserService        = services.NewUserService(userRepository)
	bookService    services.BookService        = services.NewBookService(bookRepository)
	authController controllers.AuthController  = controllers.NewAuthController(authService, jwtService)
	userController controllers.UserController  = controllers.NewUserController(userService, jwtService)
	bookController controllers.BookController  = controllers.NewBookController(bookService, jwtService)
)

func main() {
	defer configs.DatabaseDisconnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/signin", authController.SignIn)
		authRoutes.POST("/signup", authController.SignUp)
	}

	userRoutes := r.Group("api/users", middlewares.Authorize(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PATCH("/profile", userController.Update)
	}

	bookRoutes := r.Group("api/books", middlewares.Authorize(jwtService))
	{
		bookRoutes.GET("/", bookController.Index)
		bookRoutes.GET("/:id", bookController.Show)
		bookRoutes.POST("/", bookController.Store)
		bookRoutes.PATCH("/:id", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}

	r.Run()
}
