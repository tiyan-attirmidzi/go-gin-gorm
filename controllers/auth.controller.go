package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tiyan-attirmidzi/go-gin-gorm/dto"
	"github.com/tiyan-attirmidzi/go-gin-gorm/entities"
	"github.com/tiyan-attirmidzi/go-gin-gorm/helpers"
	"github.com/tiyan-attirmidzi/go-gin-gorm/services"
)

type AuthController interface {
	SignIn(ctx *gin.Context)
	SignUp(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtservice  services.JWTService
}

func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtservice:  jwtService,
	}
}

func (c *authController) SignIn(ctx *gin.Context) {

	var signInDTO dto.AuthSignInDTO
	errDTO := ctx.ShouldBind(&signInDTO)

	if errDTO != nil {
		response := helpers.ResponseError("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult := c.authService.VerifyCredential(signInDTO.Email, signInDTO.Password)

	if v, ok := authResult.(entities.User); ok {
		generatedToken := c.jwtservice.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helpers.ResponseSuccess(true, "User Sign In Successfully!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}

	if v, ok := authResult.(services.Err); ok {
		var errMsg string
		if v.Field == "email" {
			errMsg = "Email Not Found"
		} else {
			errMsg = "Password Invalid"
		}
		response := helpers.ResponseError("Please check again your credential", errMsg, helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.ResponseError("Please check again your credential", "Invalid Credential", helpers.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)

}

func (c *authController) SignUp(ctx *gin.Context) {

	var signUpDTO dto.AuthSignUpDTO
	errDTO := ctx.ShouldBind(&signUpDTO)

	if errDTO != nil {
		response := helpers.ResponseError("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(signUpDTO.Email) {
		response := helpers.ResponseError("Failed to process request", "Email Has Registered", helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	} else {
		createdUser := c.authService.CreateUser(signUpDTO)
		token := c.jwtservice.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helpers.ResponseSuccess(true, "User Sign UP Successfully!", createdUser)
		ctx.JSON(http.StatusCreated, response)
		return
	}

}
