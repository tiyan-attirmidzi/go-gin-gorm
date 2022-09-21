package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tiyan-attirmidzi/go-gin-gorm/dto"
	"github.com/tiyan-attirmidzi/go-gin-gorm/helpers"
	"github.com/tiyan-attirmidzi/go-gin-gorm/services"
)

type UserController interface {
	Update(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
	jwtService  services.JWTService
}

func NewUserController(userService services.UserService, jwtService services.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(ctx *gin.Context) {

	var userUpdateDTO dto.UserUpdateDTO
	errDTO := ctx.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helpers.ResponseError("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// TODO: Create Helper For Validation Authorization
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)

	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)

	if err != nil {
		panic(err.Error())
	}

	userUpdateDTO.ID = id

	u := c.userService.Update(userUpdateDTO)
	res := helpers.ResponseSuccess(true, "User Updated Successfully!", u)
	ctx.JSON(http.StatusOK, res)

}

func (c *userController) Profile(ctx *gin.Context) {

	// TODO: Create Helper For Validation Authorization
	authHeader := ctx.GetHeader("Authorization")

	token, err := c.jwtService.ValidateToken(authHeader)

	if err != nil {
		panic(err.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helpers.ResponseSuccess(true, "User Retrieved Successfully!", user)
	ctx.JSON(http.StatusOK, res)

}
