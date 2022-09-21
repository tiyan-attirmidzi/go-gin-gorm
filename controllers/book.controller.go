package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tiyan-attirmidzi/go-gin-gorm/dto"
	"github.com/tiyan-attirmidzi/go-gin-gorm/entities"
	"github.com/tiyan-attirmidzi/go-gin-gorm/helpers"
	"github.com/tiyan-attirmidzi/go-gin-gorm/services"
)

type BookController interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bookController struct {
	bookService services.BookService
	jwtService  services.JWTService
}

func NewBookController(bookService services.BookService, jwtService services.JWTService) BookController {
	return &bookController{
		bookService: bookService,
		jwtService:  jwtService,
	}
}

func (c *bookController) Index(ctx *gin.Context) {
	var books []entities.Book = c.bookService.Index()
	res := helpers.ResponseSuccess(true, "Books Retrieved Successfully!", books)
	ctx.JSON(http.StatusOK, res)
}

func (c *bookController) Show(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helpers.ResponseError("No params `id` was found.", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	var book entities.Book = c.bookService.Show(id)
	if (book == entities.Book{}) {
		res := helpers.ResponseError("Book not found.", "No data with given `id`", helpers.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.ResponseSuccess(true, "Book Retrieved Successfully!", book)
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *bookController) Store(ctx *gin.Context) {
	var bookCreateDTO dto.BookCreate
	errDTO := ctx.ShouldBind(&bookCreateDTO)
	if errDTO != nil {
		res := helpers.ResponseError("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := ctx.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err != nil {
			bookCreateDTO.UserID = convertedUserID
		}
		result := c.bookService.Store(bookCreateDTO)
		res := helpers.ResponseSuccess(true, "Book Stored Successfully!", result)
		ctx.JSON(http.StatusCreated, res)
	}
}

func (c *bookController) Update(ctx *gin.Context) {
	var bookUpdateDTO dto.BookUpdate
	errDTO := ctx.ShouldBind(&bookUpdateDTO)
	if errDTO != nil {
		res := helpers.ResponseError("Failed to process request.", errDTO.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, bookUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			bookUpdateDTO.UserID = id
		}
		result := c.bookService.Update(bookUpdateDTO)
		res := helpers.ResponseSuccess(true, "Book Updated Successfully!", result)
		ctx.JSON(http.StatusOK, res)
	} else {
		res := helpers.ResponseError("You dont have permission", "You're not the owner", helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, res)
	}
}

func (c *bookController) Delete(ctx *gin.Context) {
	var book entities.Book
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := helpers.ResponseError("Failed to get `id`", "No param `id` were found", helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}
	book.ID = id
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, book.ID) {
		c.bookService.Delete(book)
		res := helpers.ResponseSuccess(true, "Book Deleted Successfully!", helpers.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	} else {
		res := helpers.ResponseError("You don't have permission", "You're not the owner", helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, res)
	}
}

func (c *bookController) getUserIDByToken(token string) string {
	tokenValid, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := tokenValid.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
